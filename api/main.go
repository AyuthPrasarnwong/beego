package main

import (
	_ "api/routers"

	"api/newrelic"

	validators "api/validation"

	"fmt"
	//"regexp"
	//"strings"
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/joho/godotenv"


	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/newrelic/go-agent/_integrations/nrmysql"
)



func init() {
	validators.Run()
	newrelic.Run()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConnection := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	sqlConnect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, dbHost, dbPort, dbName)
	orm.RegisterDriver("nrmysql", orm.DRMySQL)
	orm.RegisterDataBase("default", dbConnection, sqlConnect)
	beego.BConfig.CopyRequestBody = true

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		//orm.Debug = true
	}
	beego.Run()
}



