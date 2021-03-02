package main

import (
	"github.com/MrEasonYang/letsencrypt-cloud/server/presentation"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", presentation.Login)
	r.POST("/logout", presentation.Logout)
	r.GET("/list-slaves", presentation.ListSlavesByPage)
	r.GET("/count-slaves", presentation.CountAllSlaves)
	r.POST("/create-slave", presentation.CreateSlave)
	r.POST("/remove-slave", presentation.RemoveSlave)
	r.POST("/update-slave", presentation.UpdateSlave)
	r.POST("/mannual-sync", presentation.MannualSync)
	r.Run()
}