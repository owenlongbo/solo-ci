package utils

import (
	"os/exec"
	"os"
	"solo-ci/models"
	"path"
	"path/filepath"
)

func GetWorkSpacePath(project *models.Project) (string){
	execFileRelativePath, _ := exec.LookPath(os.Args[0])
	execDirRelativePath, _ := path.Split(execFileRelativePath)
	execDirAbsPath, _ := filepath.Abs(execDirRelativePath)
	if _, err := os.Stat(execDirAbsPath + "/workspace/" + project.Name); os.IsNotExist(err) {
		os.Mkdir(execDirAbsPath + "/workspace/" + project.Name, os.ModeDir)
	}
	return execDirAbsPath + "/workspace/" + project.Name
}

func GetBuildPath(project *models.Project, build *models.Build)(string) {
	workSpace := GetWorkSpacePath(project)
	os.Mkdir(workSpace + "/" + build.Name, os.ModeDir)
	return workSpace + "/" + build.Name
}
