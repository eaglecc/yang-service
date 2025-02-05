package routers

import (
	"github.com/gin-gonic/gin"
	"yang-service/routers/xutils"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	xutil := xutils.XUtils{}

	xutil.InitXUtils(r)
	return r
}
