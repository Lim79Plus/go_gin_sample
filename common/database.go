package common

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	*gorm.DB
}

// DB is db instance
var DB *gorm.DB

// Init initialize db instance
func Init() *gorm.DB {
	DBMS, CONNECT := GetConnectInfo()
	fmt.Println("gormConnect CONNECT", CONNECT)
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	DB = db
	return db
}

// GetDB Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
