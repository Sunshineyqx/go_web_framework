package serializer

import "giligili/model"

// video 视频序列化器
type Video struct {
	Title string `json:"title"`
	Info  string `json:"info"`
}

// BuildVideo 序列化视频
func BuildVideo(video model.Video) Video {
	return Video{
		Title: video.Title,
		Info:  video.Info,
	}
}

// BuildVideoResponse 序列化视频响应
func BuildVideoResponse(video model.Video) Response {
	return Response{
		Data: BuildVideo(video),
	}
}

// BuildVideoList 序列化视频列表
func BuildVideoListResponse(videos []model.Video) Response {
	var re_videos []Video
	for i := 0; i < len(videos); i++ {
		re_videos = append(re_videos, BuildVideo(videos[i]))
	}
	return Response{
		Data: re_videos,
	}
}
