package common

import (
	"testing"

	"github.com/Lim79Plus/go_gin_sample/logger"
)

func initTestConf() {
	ConfName = "../config.ini"
	InitConf()
	logger.LogInit()
}

func TestGetNB(t *testing.T) {
	initTestConf()
	result := GetNB()
	if result.NBSecretPassword == "" {
		t.Fatal("failed test", "NBSecretPassword", result)
	}
	if result.NBRandomPassword == "" {
		t.Fatal("failed test", "NBRandomPassword", result)
	}
}
func TestGetWebPort(t *testing.T) {
	initTestConf()
	result := GetWebPort()
	t.Log("result", result)
	if result != ":8080" {
		t.Fatal("failed test", "result", result)
	}
}

func TestGetConnectInfo(t *testing.T) {
	initTestConf()
	dbms, cinfo := GetConnectInfo()
	if len(dbms) == 0 {
		t.Fatal("failed test", "dbms", dbms)
	}

	if len(cinfo) == 0 {
		t.Fatal("failed test", "cinfo", cinfo)
	}
}
