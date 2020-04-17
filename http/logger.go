package http

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Log struct {
	// Start time
	StartTime time.Time

	// EndTime shows the time after the server returns a response.
	EndTime time.Time

	// Latency is how much time the server cost to process a certain request.
	Latency time.Duration

	RequestURI string

	// 传入服务器请求的协议版本
	Proto string

	// 用户代理
	UserAgent string

	// ClientIP equals Context's ClientIP method.
	ClientIP string

	// Method is the HTTP method given to the request.
	Method string

	// Path is a path the client requests.
	Path string

	// StatusCode is HTTP response code.
	StatusCode int

	// BodySize is the size of the Response Body
	BodySize int

	//用户标识
	userId string
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		log := &Log{}
		end := time.Now()
		log.StartTime = start
		log.EndTime = end
		log.Latency = end.Sub(start)
		log.RequestURI = c.Request.RequestURI
		log.Proto = c.Request.Proto
		log.UserAgent = c.Request.UserAgent()
		log.ClientIP = c.ClientIP()
		log.Method = c.Request.Method
		log.StatusCode = c.Writer.Status()
		log.BodySize = c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		log.Path = path
		uid := getUser(c)
		log.userId = uid
		fmt.Println(log)
	}
}
