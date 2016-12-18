package test

import (
	"testing"
	"github.com/satori/go.uuid"
)

func Test_NormalTest(t *testing.T) {
	u1 := uuid.NewV1()
	t.Log(u1)
}
