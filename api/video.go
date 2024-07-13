package api

import (
	"giligili/service"

	"github.com/gin-gonic/gin"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	var service service.CreateVideoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateVideo()
		c.JSON(200, res)
	} else {
		c.JSON(444, ErrorResponse(err))
	}
}

// ShowVideoDetail 查看某个视频详情
func ShowVideoDetail(c *gin.Context) {
	var service service.ShowVideoDetailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ShowVideoDetail(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(444, ErrorResponse(err))
	}
}

// GetVideosList 获取全部视频信息
func GetVideosList(c *gin.Context) {
	var service service.GetVideoListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetVideoList()
		c.JSON(200, res)
	} else {
		c.JSON(444, ErrorResponse(err))
	}
}

// UpdateVideo 更新视频详情
func UpdateVideo(c *gin.Context) {
	var service service.UpdateVideoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateVideo(c.Params.ByName("id"))
		c.JSON(200, res)
	} else {
		c.JSON(444, ErrorResponse(err))
	}
}

// DeleteVideo 删除视频
func DeleteVideo(c *gin.Context) {
	var service service.DeleteVideoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DeleteVideo(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(444, ErrorResponse(err))
	}
}
