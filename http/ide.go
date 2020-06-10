package http

import (
	"net/http"
	"time"

	"github.com/zfd81/das/util"

	"github.com/spf13/cast"

	"github.com/zfd81/das/meta"

	"github.com/zfd81/rooster/types/container"

	"github.com/zfd81/das/dao"

	"github.com/gin-gonic/gin"
)

func FindUserByName(c *gin.Context) {
	name := c.Param("name")
	user, err := userDao.FindByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
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

func FindProjectExist(c *gin.Context) {
	code := c.Param("code")
	project, err := projectDao.FindByCode(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
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

func FindProjectByCode(c *gin.Context) {
	code := c.Param("code")
	project, err := projectDao.FindByCode(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, project)
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
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func FindProjectsByUser(c *gin.Context) {
	uid := getUser(c) //用户编号
	codeOrName := c.Query("codeOrName")
	l, err := projectDao.FindAllByUser(uid, codeOrName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, l)
}

func ModifyProject(c *gin.Context) {
	uid := getUser(c) //用户编号
	p := param(c)
	t := time.Now()
	project := &dao.ProjectInfo{
		Code:        p.GetString("code"),
		Name:        p.GetString("name"),
		Description: p.GetString("desc"),
		Model: dao.Model{
			Modifier:     uid,
			ModifiedTime: t,
		},
	}
	err := projectDao.Modify(project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func FindProjectView(c *gin.Context) {
	code := c.Param("code")
	project, err := projectDao.FindByCode(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	root := TreeNode{
		Id:      project.Code,
		Label:   project.Name,
		Creator: project.Creator,
		Type:    "pro",
	}
	catalogs, err := catalogDao.FindAllByParent(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	nodes := make([]TreeNode, 0, 10)
	for _, catalog := range catalogs {
		nodes = append(nodes, TreeNode{
			Id:      catalog.GetString("catalog_code"),
			Label:   catalog.GetString("catalog_name"),
			Creator: catalog.GetString("creator"),
			Type:    "cat",
		})
	}
	nodes = append(nodes, TreeNode{
		Id:    "202004221429",
		Label: "数据库链接",
		Type:  "conn",
	})
	nodes = append(nodes, TreeNode{
		Id:    "202004221430",
		Label: "项目开发者",
		Type:  "user",
	})
	root.Children = nodes
	c.JSON(http.StatusOK, []TreeNode{root})
}

func FindProjectViewNode(c *gin.Context) {
	code := c.Param("code")
	catalogs, err := catalogDao.FindAllByParent(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	nodes := make([]TreeNode, 0, 10)
	for _, catalog := range catalogs {
		nodes = append(nodes, TreeNode{
			Id:    catalog.GetString("catalog_code"),
			Label: catalog.GetString("catalog_name"),
			Type:  "cat",
		})
	}
	c.JSON(http.StatusOK, nodes)
}

func SaveCatalog(c *gin.Context) {
	uid := getUser(c) //用户编号
	p := param(c)
	t := time.Now()
	catalog := &dao.CatalogInfo{
		Code:    t.Format("20060102150405"),
		Name:    p.GetString("name"),
		Order:   0,
		Parent:  p.GetString("parent"),
		Project: p.GetString("project"),
		Status:  "1",
		Model: dao.Model{
			Creator:      uid,
			CreatedTime:  t,
			Modifier:     uid,
			ModifiedTime: t,
		},
	}
	err := catalogDao.Save(catalog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func SaveConnection(c *gin.Context) {
	uid := getUser(c) //用户编号
	p := param(c)
	t := time.Now()
	conn := &dao.ConnectionInfo{
		ID:           t.Format("20060102150405"),
		Name:         p.GetString("name"),
		Driver:       p.GetString("driver"),
		Address:      p.GetString("address"),
		Port:         p.GetString("port"),
		UserName:     p.GetString("user"),
		Password:     p.GetString("password"),
		DatabaseName: p.GetString("db"),
		Project:      p.GetString("project"),
		Status:       "1",
		Model: dao.Model{
			Creator:      uid,
			CreatedTime:  t,
			Modifier:     uid,
			ModifiedTime: t,
		},
	}
	err := connectionDao.Save(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func TestConnection(c *gin.Context) {
	p := param(c)
	conn := &meta.Connection{
		Driver:       p.GetString("driver"),
		Address:      p.GetString("address"),
		Port:         cast.ToInt(p.GetString("port")),
		UserName:     p.GetString("user"),
		Password:     p.GetString("password"),
		DatabaseName: p.GetString("db"),
	}
	err := conn.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func RemoveConnectionById(c *gin.Context) {
	id := c.Query("id")
	err := connectionDao.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func ModifyConnection(c *gin.Context) {
	uid := getUser(c) //用户编号
	p := param(c)
	t := time.Now()
	conn := &dao.ConnectionInfo{
		ID:           t.Format("20060102150405"),
		Name:         p.GetString("name"),
		Driver:       p.GetString("driver"),
		Address:      p.GetString("address"),
		Port:         p.GetString("port"),
		UserName:     p.GetString("user_name"),
		Password:     p.GetString("password"),
		DatabaseName: p.GetString("db"),
		Model: dao.Model{
			Modifier:     uid,
			ModifiedTime: t,
		},
	}
	err := connectionDao.Modify(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func FindConnectionById(c *gin.Context) {
	id := c.Param("id")
	conn, err := connectionDao.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, conn)
}

func FindConnectionsByProject(c *gin.Context) {
	pro := c.Param("pro")
	l, err := connectionDao.FindAllByProject(pro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, l)
}

func SaveUserProject(c *gin.Context) {
	uid := getUser(c) //用户编号
	p := param(c)
	t := time.Now()
	code := p.GetString("code")
	ids, _ := p.Get("ids")
	relations := []container.JsonMap{}
	idSlice, ok := ids.([]interface{})
	if ok {
		for _, v := range idSlice {
			relations = append(relations, container.JsonMap{
				"id":   v,
				"code": code,
				"uid":  uid,
				"t":    t,
			})
		}
	}
	err := userDao.SaveUserProject(relations)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func RemoveUserProject(c *gin.Context) {
	p := param(c)
	code := p.GetString("code")
	ids, _ := p.Get("ids")
	relations := []container.JsonMap{}
	idSlice, ok := ids.([]interface{})
	if ok {
		for _, v := range idSlice {
			relations = append(relations, container.JsonMap{
				"id":   v,
				"code": code,
			})
		}
	}
	err := userDao.DeleteUserProject(relations)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func FindUsersByProject(c *gin.Context) {
	pro := c.Param("pro")
	l, err := userDao.FindAllByProject(pro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, l)
}

func FindUsersNotInProject(c *gin.Context) {
	uid := getUser(c) //用户编号
	pro := c.Param("pro")
	l, err := userDao.FindAllNotInProject(uid, pro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, l)
}

func FindServicesByCatalog(c *gin.Context) {
	catalog := c.Param("catalog")
	l, err := serviceDao.FindAllByCatalog(catalog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, l)
}

func RenameService(c *gin.Context) {
	uid := getUser(c) //用户编号
	p := param(c)
	t := time.Now()
	serv := &dao.ServiceInfo{
		Name:    p.GetString("name"),
		Version: p.GetString("ver"),
		Model: dao.Model{
			Modifier:     uid,
			ModifiedTime: t,
		},
	}
	err := serviceDao.ModifyName(serv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func MoveService(c *gin.Context) {
	uid := getUser(c) //用户编号
	p := param(c)
	t := time.Now()
	serv := &dao.ServiceInfo{
		Catalog: p.GetString("catalog"),
		Version: p.GetString("ver"),
		Model: dao.Model{
			Modifier:     uid,
			ModifiedTime: t,
		},
	}
	err := serviceDao.ModifyCatalog(serv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func FindServiceById(c *gin.Context) {
	id := c.Param("id")
	version := c.Param("ver")
	serv, err := serviceDao.FindById(id, version)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, serv)
}

func SaveService(c *gin.Context) {
	uid := getUser(c) //用户编号
	p := param(c)
	t := time.Now()
	serv := &dao.ServiceInfo{
		Code:    t.Format("20060102150405"),
		Name:    p.GetString("name"),
		Catalog: p.GetString("driver"),
		Type:    p.GetString("address"),
		Sql:     p.GetString("port"),
		Param:   p.GetString("user"),
		Version: "1",
		Status:  "1",
		Model: dao.Model{
			Creator:      uid,
			CreatedTime:  t,
			Modifier:     uid,
			ModifiedTime: t,
		},
	}
	err := serviceDao.Save(serv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func ModifyService(c *gin.Context) {
	uid := getUser(c) //用户编号
	p := param(c)
	t := time.Now()
	serv := &dao.ServiceInfo{
		Code:    p.GetString("code"),
		Sql:     p.GetString("port"),
		Param:   p.GetString("user"),
		Version: p.GetString("ver"),
		Model: dao.Model{
			Modifier:     uid,
			ModifiedTime: t,
		},
	}
	err := serviceDao.Modify(serv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func RemoveService(c *gin.Context) {
	uid := getUser(c) //用户编号
	p := param(c)
	t := time.Now()
	serv := &dao.ServiceInfo{
		Code:    p.GetString("code"),
		Version: p.GetString("ver"),
		Status:  "-1",
		Model: dao.Model{
			Modifier:     uid,
			ModifiedTime: t,
		},
	}
	err := serviceDao.ModifyStatus(serv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func ParsingParam(c *gin.Context) {
	p := param(c)
	ps, err := util.ParsingSql(p.GetString("sql"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	nodes := make([]TreeNode, 0, 10)
	for _, v := range ps {
		val, ok := v.Val.([]util.KV)
		if ok {
			node := TreeNode{
				Id:       v.Key,
				Label:    v.Key,
				Type:     "cat",
				Children: make([]TreeNode, 0, 10),
			}
			for _, v1 := range val {
				node.Children = append(node.Children, TreeNode{
					Id:    v1.Key,
					Label: cast.ToString(v1.Val),
					Type:  "param",
				})
			}
			nodes = append(nodes, node)
		} else {
			nodes = append(nodes, TreeNode{
				Id:    v.Key,
				Label: cast.ToString(v.Val),
				Type:  "param",
			})
		}
	}
	c.JSON(http.StatusOK, nodes)
}
