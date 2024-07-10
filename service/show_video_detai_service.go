package service

import (
	"giligili/model"
	"giligili/serializer"
)

type ShowVideoDetailService struct {
}

func (svc *ShowVideoDetailService) ShowVideoDetail(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.DBErr("视频获取失败...", err)
	}
	return serializer.BuildVideoResponse(video)
}
