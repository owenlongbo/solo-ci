package test

import (
	"testing"
	"os/exec"
	"os"
	"path"
	"path/filepath"
)

func Test_NormalTest(t *testing.T) {
	execFileRelativePath, _ := exec.LookPath(os.Args[0])
	execDirRelativePath, _ := path.Split(execFileRelativePath)
	execDirAbsPath, _ := filepath.Abs(execDirRelativePath)
	t.Log(execDirAbsPath)
}
