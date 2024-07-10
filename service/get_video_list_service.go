package service

import (
	"giligili/model"
	"giligili/serializer"
)

type GetVideoListService struct {
}

func (svc *GetVideoListService) GetVideoList() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.DBErr("视频列表查找错误...", err)
	}
	return serializer.BuildVideoListResponse(videos)
}
