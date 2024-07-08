package serializer

import "giligili/model"

// video 视频序列化器
type Video struct {
	Title string `json:"title"`
	Info  string `json:"info"`
}

// Buildvideo 序列化视频
func BuildVideo(video model.Video) Video {
	return Video{
		Title: video.Title,
		Info:  video.Info,
	}
}

// BuildvideoResponse 序列化视频响应
func BuildVideoResponse(video model.Video) Response {
	return Response{
		Data: BuildVideo(video),
	}
}
