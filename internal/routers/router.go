package routers

import (
	"github.com/gin-gonic/gin"
)

/**
* @program: src
* @description:
* @author: 占翔昊
* @create 2020-10-05 18:52
**/

func NewRouter() (*gin.Engine) {
	r := gin.New()
	r.Use(gin.Logger(),gin.Recovery())
	v1 := r.Group("/index")
	{
		v1.GET("/hello")
	}
	return r
}