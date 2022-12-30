package serializer

import "net/http"

// RTSPPlayPath RTSP play 信息序列化器
type RTSPPlayPath struct {
	Path string `json:"path"`
}

// BuildRTSPPlayPathResponse 序列化 RTSP play 响应
func BuildRTSPPlayPathResponse(path string) *Response {
	return &Response{
		Code: http.StatusOK,
		Data: &RTSPPlayPath{
			Path: path,
		},
		Message: "success",
	}
}
