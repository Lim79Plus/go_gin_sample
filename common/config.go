package common

import "github.com/go-ini/ini"
import "fmt"

var confname = "config.ini"

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

// Conf object
var Conf *ini.File

// InitConf generater
func InitConf(){
	cfg, err := ini.Load(confname)
	if err != nil{
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

func getWebConf() *ConfiglistWeb{
	return &ConfiglistWeb{
		Port:    Conf.Section("web").Key("port").String(),
	}
}

// GetWebPort return web server port
func GetWebPort() string{
	c := getWebConf()
	return ":" + c.Port
}
// GetConnectInfo return db info (dbms, connect info)
func GetConnectInfo() (string, string) {
	c := getDBConf()
	fmt.Println("c",c)
	DBMS := c.SQLDriver
	USER := c.DBUser
	PASS := c.DBPass
	PROTOCOL := "tcp(" + c.DBAddress + ":" + c.DBPort + ")"
	DBNAME := c.DBName
	OPTION := "charset=utf8"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	return DBMS, CONNECT
}
