package models

import (
	"time"
	"os/exec"
	"bytes"
	"os"
	"io/ioutil"
	"solo-ci/conf"
	"github.com/astaxie/beego/orm"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego"
)

type Build struct {
	Id        int   `orm:"pk;auto;unique" json:"id"` //主键
	Name      string `json:"name" json:"name"`
	Result    string `json:"result" json:"result"`
	Project   *Project `orm:"rel(fk)" json:"-"`
	IsSuccess bool `json:"is_success"`
}

type BuildConfig struct {
	GetList      []string `json:"get_list"`    //Go get list
	ZipList      []string `json:"zip_list"`    //需要打包的文件
	AfterScript  string `json:"after_script"`  //build 之后
	BeforeScript string `json:"before_script"` //build 之前
}

func NewBuild(project *Project) {
	//创建build 记录
	build := new(Build)
	build.Project = project
	build.Name = time.Now().Format("2006-01-02T15:04:05.000Z")
	buildPath := getBuildPath(project, build)
	var result bytes.Buffer
	//git clone
	errGit := RunScript(result, []*exec.Cmd{
		exec.Command("git", "clone", "-b", project.Branch, project.Url, buildPath),
	})
	if errGit != nil {
		saveBuild(build, false, errGit, result)
		return
	}
	//read config
	fileData, errFile := getFileData(project, build)
	if errFile != nil {
		saveBuild(build, false, errFile, result)
		return
	}
	buildConfig := new(BuildConfig)
	errConfig := json.Unmarshal(fileData, buildConfig)
	if errConfig != nil {
		fmt.Println(errConfig)
		saveBuild(build, false, errConfig, result)
		return
	}
	fmt.Println(buildConfig)
	//BeforeScript
	if buildConfig.BeforeScript != "" {
		in := bytes.NewBuffer(nil)
		cmd := exec.Command("sh")
		cmd.Stdin = in
		in.WriteString("cd " + buildPath + "\n")
		in.WriteString(buildConfig.BeforeScript)
		beforeResp, errBefore := cmd.CombinedOutput()
		if errBefore != nil {
			saveBuild(build, false, errBefore, result)
			return
		}
		result.Write(beforeResp)
	}
	//exec config
	getList := make([]*exec.Cmd, len(buildConfig.GetList))
	for index, pack := range buildConfig.GetList {
		getList[index] = exec.Command(conf.GOROOT + "/bin/go", "get", pack)
	}
	errGet := RunScript(result, getList)
	if errGet != nil {
		beego.Error(errGet.Error())
		saveBuild(build, false, errGet, result)
		return
	}
	//exec build and clean
	errBuild := RunScript(result, []*exec.Cmd{
		exec.Command("ln", "-s", buildPath, conf.GOPATH + "/src"),
		exec.Command(conf.GOROOT + "/bin/go", "build", "-o", build.Name, project.MainPath),
		exec.Command(conf.GOROOT + "/bin/go", "clean", build.Name),
		exec.Command(conf.GOROOT + "/bin/go", "test", build.Name),
		exec.Command("mv", build.Name, "workspace/" + project.Name + "/" + build.Name + "/" + project.Name),
		exec.Command("rm", conf.GOPATH + "/src/" + build.Name),
	})
	if errBuild != nil {
		saveBuild(build, false, errBuild, result)
		return
	}
	//pack
	os.Mkdir(buildPath + "/pack-" + project.Name, 0766)
	if len(buildConfig.ZipList) != 0 {
		zipLength := len(buildConfig.GetList)
		zipList := make([]*exec.Cmd, zipLength)
		for index, pack := range buildConfig.ZipList {
			zipList[index] = exec.Command("cp", "-R", buildPath + "/" + pack, buildPath + "/pack-" + project.Name + "/")
		}
		errZip := RunScript(result, zipList)
		if errZip != nil {
			saveBuild(build, false, errZip, result)
			return
		}
	}
	exec.Command("cp", "-R", buildPath + "/" + project.Name, buildPath + "/pack-" + project.Name + "/").Run()
	exec.Command("tar", "-zcvf", buildPath + "/" + project.Name + ".tar.gz", "-C", buildPath, "pack-" + project.Name).Run()
	exec.Command("rm", "-rf", buildPath + "/pack-" + project.Name).Run()
	//AfterScript
	if buildConfig.AfterScript != "" {
		in := bytes.NewBuffer(nil)
		cmd := exec.Command("sh")
		cmd.Stdin = in
		in.WriteString("cd " + buildPath + "\n")
		in.WriteString(buildConfig.AfterScript)
		afterResp, errAfter := cmd.CombinedOutput()
		if errAfter != nil {
			saveBuild(build, false, errAfter, result)
			return
		}
		result.Write(afterResp)
	}
	//finish
	saveBuild(build, true, nil, result)
}

func getBuildPath(project *Project, build *Build) (string) {
	workSpace := GetWorkSpacePath(project)
	os.Mkdir(workSpace + "/" + build.Name, 0766)
	return workSpace + "/" + build.Name
}

func RunScript(result bytes.Buffer, cmdList []*exec.Cmd) (error) {
	for _, cmd := range cmdList {
		out, err := cmd.CombinedOutput()
		result.Write(out)
		if err != nil {
			result.WriteString(err.Error())
			return err
		}
	}
	return nil
}

func getFileData(project *Project, build *Build) ([]byte, error) {
	file, errFile := os.Open(getBuildPath(project, build) + "/" + project.Path)
	if errFile != nil {
		return nil, errFile
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

func saveBuild(build *Build, status bool, err error, result bytes.Buffer) {
	if err != nil {
		result.WriteString(err.Error())
	}
	build.IsSuccess = status
	build.Result = result.String()
	orm.NewOrm().Insert(build)
}