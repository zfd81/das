package http

import (
	"net/http"

	"github.com/zfd81/das/conf"

	"github.com/gin-gonic/gin"
	"github.com/zfd81/das/dao"
	"github.com/zfd81/rooster/types/container"
	"github.com/zfd81/rooster/util"
)

const (
	jwtSecret               string = "zhang20181231jwtsecret"
	AuthenticationTokenName string = "zxcvb"
	tokenPrefix             string = "qwert"
	EffectiveDuration       int    = 18000
	HeaderUidName                  = "u_"
	HeaderTokenNameName            = "atn"
	HeaderTokenValueName           = "atv"
)

var (
	prefixLength  int
	config                          = conf.GetConfig()
	userDao       dao.UserDao       = dao.NewUserDao()
	projectDao    dao.ProjectDao    = dao.NewProjectDao()
	catalogDao    dao.CatalogDao    = dao.NewCatalogDao()
	connectionDao dao.ConnectionDao = dao.NewConnectionDao()
	serviceDao    dao.ServiceDao    = dao.NewServiceDao()
)

func init() {
	prefixLength = len(tokenPrefix)
	err := dao.Conf(config.Database.Dialect, config.Database.Address, config.Database.Port, config.Database.User, config.Database.Pwd, config.Database.Name)
	if err != nil {
		panic(err.Error())
	}
}
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(AuthenticationTokenName)
		if token != "" {
			tokenLength := len(token)
			if tokenLength > prefixLength {
				prefix := util.Left(token, prefixLength)
				if tokenPrefix == prefix {
					sec := NewSecurity(jwtSecret)
					claims, err := sec.ParseToken(util.Right(token, tokenLength-prefixLength))
					if err == nil {
						c.Request.Header.Add(HeaderUidName, claims.Uid)
						c.Next()
						return
					} else {
						if err == TokenExpired {
							c.JSON(http.StatusUnauthorized, gin.H{
								"msg": "授权已过期",
							})
							c.Abort()
							return
						}
						c.JSON(http.StatusUnauthorized, gin.H{
							"msg": err.Error(),
						})
						c.Abort()
						return
					}
				}
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "indicating that the request requires HTTP authentication.",
		})
		c.Abort()
		return
	}
}

func getUser(c *gin.Context) (uid string) {
	uid = c.Request.Header.Get(HeaderUidName)
	return
}

func param(c *gin.Context) container.Map {
	p := container.JsonMap{}
	c.ShouldBind(&p)
	return p
}

func Login(c *gin.Context) {
	p := param(c)
	name := p.GetString("name")
	pwd := p.GetString("password")
	if name == "" || pwd == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "请求参数错误",
		})
		return
	}
	user, err := userDao.FindByNameAndPwd(name, pwd)
	if err != nil {
		c.Request.Header.Add(HeaderUidName, name+"/"+pwd)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error,
		})
		return
	}
	if user == nil {
		c.Request.Header.Add(HeaderUidName, name+"/"+pwd)
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "用户名或密码错误",
		})
		return
	}
	sec := NewSecurity(jwtSecret)
	claims := PortalClaims{
		Uid: user.GetString("id"),
	}
	token, err := sec.CreateToken(&claims, EffectiveDuration)
	if err != nil {
		c.Request.Header.Add(HeaderUidName, name+"/"+pwd)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error,
		})
		return
	}
	c.Request.Header.Add(HeaderUidName, claims.Uid)
	c.Header(HeaderTokenNameName, AuthenticationTokenName)
	c.Header(HeaderTokenValueName, tokenPrefix+token)
	c.JSON(http.StatusOK, user)
}
