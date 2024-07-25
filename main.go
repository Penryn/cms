package main

import (
	midwares "cms/internal/middleware"
	"cms/internal/pkg/database"
	"cms/internal/router"
	"cms/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db:=database.Init()
	service.ServiceInit(db)
	r := gin.Default()
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	router.Init(r)
	err:=r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}