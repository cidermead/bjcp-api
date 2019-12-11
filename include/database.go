package include

import (
	"fmt"
  "os"
  "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/cidermead/bjcp-api/config"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB
var err error

// InitDB opens a database and saves the reference to `Database` struct.
func InitDB() *gorm.DB {
	var db = DB

	config := config.InitConfig()

  driver := config.Database.Driver
	database := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := config.Database.Host
	port := os.Getenv("DB_PORT")

	if driver == "postgres" {

		db, err = gorm.Open("postgres", "host="+host+" port="+port+" user="+username+" dbname="+database+"  sslmode=disable password="+password)
		if err != nil {
			fmt.Println("db err: ", err)
		}

	} else if driver == "mysql" {

		db, err = gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			fmt.Println("db err: ", err)
		}
	}

	db.LogMode(true)
	DB = db

	return DB
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}
