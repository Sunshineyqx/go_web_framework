package service

import (
	"giligili/model"
	"giligili/serializer"
)

type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=1,max=30"`
	Info  string `form:"info" json:"info" binding:"min=0,max=200"`
}

func (svc *UpdateVideoService) UpdateVideo(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code:  serializer.CodeNotFound,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}
	video.Info = svc.Info
	video.Title = svc.Title
	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  serializer.CodeDBError,
			Msg:   "视频更新失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: serializer.CodeSuccess,
		Msg:  "成功更新视频",
	}
}
