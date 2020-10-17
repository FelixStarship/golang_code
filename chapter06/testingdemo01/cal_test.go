package testingdemo01

import (
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := AddUpper(100)
	if res == 0 {
		t.Fatalf("AddUpper()执行错误")
	}
	t.Log("AddUpper()执行正确")
}
