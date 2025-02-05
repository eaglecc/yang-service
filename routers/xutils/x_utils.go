package xutils

import (
	"github.com/gin-gonic/gin"
	"yang-service/controller/xutils"
)

type XUtils struct {
}

func (*XUtils) InitXUtils(Router *gin.Engine) {
	handler := xutils.XUtilsHandler{}
	groups := Router.Group("api/v1/xutils")
	{
		groups.GET("/getAngryHistory", handler.GetAngryHistory)
		groups.POST("/addAngryHistory", handler.AddAngryHistory)

	}
}
