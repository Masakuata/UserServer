package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"xgUserServer/routes"
)

func main() {
	var engine *gin.Engine

	if _, isDebug := os.LookupEnv("DEBUG"); isDebug {
		engine = gin.Default()
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
	}

	engine.GET("/", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})
	routes.UserRoutes(engine)
	err := engine.Run("0.0.0.0:42102")
	if err != nil {
		return
	}
}
