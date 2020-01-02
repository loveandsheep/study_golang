package database

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type People struct {
	gorm.Model
	Name     string    `json:"name"`
	Age      int       `json:age`
	Birthday time.Time `json:birthday`
}

func GormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "Aidkun05230523"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "gotest"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	return db
}
