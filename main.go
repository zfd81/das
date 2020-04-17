package main

import (
	"fmt"

	"github.com/zfd81/das/conf"
	"github.com/zfd81/das/http"

	"github.com/gin-gonic/gin"
)

var (
	config = conf.GetConfig()
)

func main() {
	r := gin.Default()
	r.Use(http.Logger())
	r.POST("/login", http.Login)
	authorized := r.Group("/auth", http.Authentication())
	{
		das := authorized.Group("/das")
		{
			das.GET("/test", http.FindUserByName)
		}

		if config.Mode == conf.ModeDevelop {
			ide := authorized.Group("/ide")
			{
				ide.GET("/user/name/:name", http.FindUserByName)
				ide.GET("/project/code/:code", http.FindProjectByCode)
				ide.POST("/project", http.SaveProject)
			}
		}
	}
	fmt.Println(config.Name)
	fmt.Println(config.Database.Dialect)
	r.Run(fmt.Sprintf(":%d", config.Http.Port))
}
