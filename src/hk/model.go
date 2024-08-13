package hk

const (
	BasePrefix = "/artemis"
	VideoUrl   = "/api/video/v1/cameras/previewURLs"

	SubscribeInfoUrl = "/api/common/v1/event/getTopicInfo" // 获取事件订阅信息
)

type BaseInfo struct {
	Key        string
	Secret     string
	BaseServer string
}

type CommonRsp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type VideoData struct {
	CommonRsp
	Data struct {
		URL string `json:"url"`
	} `json:"data"`
}

type SubscribeInfoData struct {
	CommonRsp
	Data struct {
		Host      string `json:"host"`
		ClientId  string `json:"clientId"`
		UserName  string `json:"userName"`
		Password  string `json:"password"`
		TopicName struct {
			Field1 string `json:"3187675137"`
			Field2 string `json:"3204452353"`
		} `json:"topicName"`
	} `json:"data"`
}
