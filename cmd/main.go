package main

import (
	"time"

	"architecture.com/api/route"
	"architecture.com/bootstrap"
	"architecture.com/domain"
	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	db := app.MySQL.DB

	db.AutoMigrate(&domain.User{})

	defer app.CloseDBConnection()

	timeout := time.Duration(app.Env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(app.Env, timeout, db, gin)

	gin.Run(app.Env.ServerAddress)

}