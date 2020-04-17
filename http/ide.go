package http

import (
	"net/http"
	"time"

	"github.com/zfd81/das/dao"

	"github.com/gin-gonic/gin"
)

func FindUserByName(c *gin.Context) {
	name := c.Param("name")
	user, err := userDao.FindByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error,
		})
		return
	}
	exist := true
	if user.ID == "" {
		exist = false
	}
	c.JSON(http.StatusOK, gin.H{
		"exist": exist,
	})
}

func FindProjectByCode(c *gin.Context) {
	code := c.Param("code")
	project, err := projectDao.FindByCode(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error,
		})
		return
	}
	exist := true
	if project.Code == "" {
		exist = false
	}
	c.JSON(http.StatusOK, gin.H{
		"exist": exist,
	})
}

func SaveProject(c *gin.Context) {
	uid := getUser(c) //用户编号
	p := param(c)
	t := time.Now()
	project := &dao.ProjectInfo{
		Code:        p.GetString("code"),
		Name:        p.GetString("name"),
		Description: p.GetString("desc"),
		Status:      "1",
		Model: dao.Model{
			Creator:      uid,
			CreatedTime:  t,
			Modifier:     uid,
			ModifiedTime: t,
		},
	}
	err := projectDao.Save(project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
