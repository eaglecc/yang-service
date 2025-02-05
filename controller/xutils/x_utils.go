package xutils

import (
	"github.com/gin-gonic/gin"
	"yang-service/service/xutils"
)

type XUtilsHandler struct {
}

func (api *XUtilsHandler) GetAngryHistory(c *gin.Context) {
	xutilserv := xutils.XUtilServ{}
	recordsVo, err := xutilserv.GetAngryHistory()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "获取历史记录失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "获取历史记录成功",
		"data":    recordsVo,
	})
}
