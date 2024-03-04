package main

import (
	// "github.com/daniadjt/task-5-pbi-btpns-dania-djatniko/seed"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nicobrahmantiyo/task-5-pbi-btpns-nico-brahmantiyo/helpers"
	"github.com/nicobrahmantiyo/task-5-pbi-btpns-nico-brahmantiyo/models"
	"github.com/nicobrahmantiyo/task-5-pbi-btpns-nico-brahmantiyo/router/photoRouter"
	"github.com/nicobrahmantiyo/task-5-pbi-btpns-nico-brahmantiyo/router/userRouter"
)

func init() {
	helpers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	models.ConnectDatabase()
	userRouter.UserRouter(r)
	photoRouter.PhotoRouter(r)

	// seed.Load(models.DB)

	r.Run()
}