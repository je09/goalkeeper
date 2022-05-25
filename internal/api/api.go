package api

import "github.com/gin-gonic/gin"

var engine *gin.Engine

func Init() {
	engine = gin.Default()
}

func initRoutes() {
	engine.POST("/register")
}
