package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTMLFormDetails struct {
	Message string
}

func preRunTasks() {

}

func main() {
	preRunTasks()
	router := gin.Default()
	router.LoadHTMLGlob("web/*")
	router.GET("/", func(c *gin.Context) {
		details := HTMLFormDetails{
			Message: "hello",
		}
		c.HTML(http.StatusOK, "index.html", details)
	})
	router.StaticFS("/css", http.Dir("/assets/css"))
	router.StaticFS("/js", http.Dir("/assets/js"))
	router.Use(cors.Default())
	router.Run(":80")
}
