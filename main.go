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
				ide.POST("/user_pro", http.SaveUserProject)
				ide.DELETE("/user_pro", http.RemoveUserProject)
				ide.GET("/user/pro/:pro", http.FindUsersByProject)
				ide.GET("/user/nipro/:pro", http.FindUsersNotInProject)
				ide.GET("/project/exist/:code", http.FindProjectExist)
				ide.GET("/project/code/:code", http.FindProjectByCode)
				ide.POST("/project", http.SaveProject)
				ide.PUT("/project", http.ModifyProject)
				ide.GET("/project/user", http.FindProjectsByUser)
				ide.GET("/project_view/:code", http.FindProjectView)
				ide.GET("/project_view_node/:code", http.FindProjectViewNode)
				ide.POST("/catalog", http.SaveCatalog)
				ide.POST("/conn_test", http.TestConnection)
				ide.POST("/conn", http.SaveConnection)
				ide.DELETE("/conn", http.RemoveConnectionById)
				ide.PUT("/conn", http.ModifyConnection)
				ide.GET("/conn/id/:id", http.FindConnectionById)
				ide.GET("/conn/pro/:pro", http.FindConnectionsByProject)
				ide.POST("/serv/param/parsing", http.ParsingParam)
			}
		}
	}
	fmt.Println(config.Name)
	fmt.Println(config.Database.Dialect)
	r.Run(fmt.Sprintf(":%d", config.Http.Port))
}
