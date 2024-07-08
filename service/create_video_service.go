package service

import (
	"giligili/model"
	"giligili/serializer"
)

type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=1,max=30"`
	Info  string `form:"info" json:"info" binding:"min=0,max=200"`
}

func (svc *CreateVideoService) CreateVideo() serializer.Response {
	video := model.Video{
		Title: svc.Title,
		Info:  svc.Info,
	}
	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.DBErr("视频保存失败...", err)
	}
	return serializer.BuildVideoResponse(video)
}
