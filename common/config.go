package common

import (
	"github.com/Lim79Plus/go_gin_sample/logger"
	"github.com/go-ini/ini"
)

// ConfName config file name
var ConfName = "config.ini"

// ConfigListDB include db
type ConfigListDB struct {
	DBName    string
	SQLDriver string
	DBPass    string
	DBAddress string
	DBPort    string
	DBUser    string
}

// ConfiglistWeb include web
type ConfiglistWeb struct {
	Port string
}

// ConfigConst include constract
type ConfigConst struct {
	NBSecretPassword string
	NBRandomPassword string
}

// Conf object
var Conf *ini.File

// InitConf generater
func InitConf() {
	cfg, err := ini.Load(ConfName)
	if err != nil {
		panic(err)
	}
	Conf = cfg
}

func getDBConf() *ConfigListDB {
	return &ConfigListDB{
		DBName:    Conf.Section("db").Key("name").String(),
		SQLDriver: Conf.Section("db").Key("driver").String(),
		DBPass:    Conf.Section("db").Key("pass").String(),
		DBAddress: Conf.Section("db").Key("address").String(),
		DBPort:    Conf.Section("db").Key("port").String(),
		DBUser:    Conf.Section("db").Key("user").String(),
	}
}

func getWebConf() *ConfiglistWeb {
	return &ConfiglistWeb{
		Port: Conf.Section("web").Key("port").String(),
	}
}

func getNB() *ConfigConst {
	return &ConfigConst{
		NBSecretPassword: Conf.Section("const").Key("NBSecretPassword").String(),
		NBRandomPassword: Conf.Section("const").Key("NBRandomPassword").String(),
	}
}

// GetNB return password and random
func GetNB() *ConfigConst {
	c := getNB()
	return c
}

// GetWebPort return web server port
func GetWebPort() string {
	c := getWebConf()
	return ":" + c.Port
}

// GetConnectInfo return db info (dbms, connect info)
func GetConnectInfo() (string, string) {
	c := getDBConf()
	logger.Trace("c", c)
	DBMS := c.SQLDriver
	USER := c.DBUser
	PASS := c.DBPass
	PROTOCOL := "tcp(" + c.DBAddress + ":" + c.DBPort + ")"
	DBNAME := c.DBName
	PARSETIME := "parseTime=True"
	OPTION := "charset=utf8"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION + "&" + PARSETIME
	return DBMS, CONNECT
}
