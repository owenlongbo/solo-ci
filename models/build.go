package models

import (
	"time"
	"os/exec"
	"solo-ci/utils"
	"bytes"
	"solo-ci/conf"
)

type Build struct {
	Id      int   `orm:"pk;auto;unique" json:"id"` //主键
	Name    string `json:"name" json:"name"`
	Result  string `json:"result" json:"result"`
	Project *Project `orm:"rel(fk)" json:"-"`
}

func NewBuild(project *Project) (*Build) {
	build := new(Build)
	build.Project = project
	build.Name = time.Now().Format("2006-01-02T15:04:05.000Z")
	//git clone
	execList := []exec.Cmd{
		exec.Command("git clone", project.Url, utils.GetBuildPath(project, build)),
		exec.Command("ln -s", utils.GetBuildPath(project, build), conf.GOPATH),
		exec.Command("go build", "-o", project.Name),
		exec.Command("go clean"),
	}
	build.Result = GetResult(execList)
	return build
}

func GetResult(cmdList []*exec.Cmd) (string){
	var buffer bytes.Buffer
	for _, cmd := range cmdList {
		out,err := cmd.CombinedOutput()
		buffer.Write(out)
		if err != nil {
			buffer.WriteString(err.Error())
			break
		}
	}
	return buffer.String()
}
