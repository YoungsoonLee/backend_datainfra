package main

import (
	"os"
	"strconv"

	"github.com/YoungsoonLee/backend_datainfra/models"
	_ "github.com/YoungsoonLee/backend_datainfra/routers"
	"github.com/joho/godotenv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	models.RegisterDB()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		beego.Error("Error loading .env file")
	}

	// PORT
	PORT, _ := strconv.Atoi(os.Getenv("PORT"))
	if PORT == 0 {
		PORT, _ = strconv.Atoi(beego.AppConfig.String("httpport"))
	}
	beego.BConfig.Listen.HTTPPort = PORT

	// RUNMODE
	RUNMODE := os.Getenv("BEEGO_RUNMODE")
	if RUNMODE == "" {
		RUNMODE = beego.AppConfig.String("runmode")
	}
	beego.BConfig.RunMode = RUNMODE

	beego.Info("beego runmode: ", RUNMODE)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		orm.Debug = true // !
	}

	orm.RunSyncdb("default", false, true)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{

		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},

		//AllowOrigins:     []string{"http://localhost", "http://*.naddic.com"},
		AllowOrigins: []string{"*"},

		AllowCredentials: true,
	}))

	/*
		beego.InsertFilter("*", beego.BeforeExec, func(ctx *context.Context){
			ctx.Input.Data["requestid"] = UUID() // generate uuid
		})
		// print
		beego.Info(ctx.Input.Data["requestid"], xxx)
	*/

	beego.SetLogger(logs.AdapterFile, `{"filename":"./logs/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)

	beego.Run()
}
