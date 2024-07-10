package service

import (
	"giligili/model"
	"giligili/serializer"
)

type DeleteVideoService struct {
}

func (svc *DeleteVideoService) DeleteVideo(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code:  serializer.CodeParamErr,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}
	err = model.DB.Delete(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  serializer.CodeDBError,
			Msg:   "视频删除失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: serializer.CodeSuccess,
		Msg:  "成功删除视频",
	}
}
