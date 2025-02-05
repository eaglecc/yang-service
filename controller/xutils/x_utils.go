package xutils

import (
	"github.com/gin-gonic/gin"
	"yang-service/service/xutils"
	"yang-service/vo"
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

func (api *XUtilsHandler) AddAngryHistory(c *gin.Context) {
	var record vo.AngryHistoryRecordVo
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(400, gin.H{
			"message": "请求参数错误",
		})
		return
	}
	xutilserv := xutils.XUtilServ{}
	res, err := xutilserv.AddAngryHistory(record)
	if res == 0 || err != nil {
		c.JSON(500, gin.H{
			"message": "记录保存失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "记录保存成功",
	})
}
