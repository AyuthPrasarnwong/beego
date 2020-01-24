package newrelic


import (
	//"fmt"
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/joho/godotenv"
	"github.com/newrelic/go-agent"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := newrelic.NewConfig(os.Getenv("NEW_RELIC_APP_NAME"), os.Getenv("NEW_RELIC_LICENSE_KEY"))
	config.HostDisplayName = "api-report.eggsmartpos.com"
	config.Labels = map[string]string{
		"Project": "horeca",
		"Environment": os.Getenv("APP_ENV"),
		"Server": "on-premise",
		"Tier": "backend",
		"Language": "golang",
	}
	app, err := newrelic.NewApplication(config)
	if err != nil {
		beego.Warn(err.Error())
		return
	}
	NewrelicAgent = app
	beego.InsertFilter("*", beego.BeforeRouter, StartTransaction, false)
	beego.InsertFilter("*", beego.AfterExec, NameTransaction, false)
	beego.InsertFilter("*", beego.FinishRouter, EndTransaction, false)
	beego.Info("NewRelic agent start")
}