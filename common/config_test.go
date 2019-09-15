package common

import "testing"

func TestGetNBSecretPassword(t *testing.T) {
	InitConf()
	result := GetNB()
	if result.NBSecretPassword == "" {
		t.Fatal("failed test", "result", result)
	}
	if result.NBRandomPassword == "" {
		t.Fatal("failed test", "result", result)
	}
}
func TestGetWebPort(t *testing.T) {
	InitConf()
	result := GetWebPort()
	t.Log("result", result)
	if result != ":8080" {
		t.Fatal("failed test", "result", result)
	}
}
