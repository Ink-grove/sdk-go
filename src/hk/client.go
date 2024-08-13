package hk

import (
	"bytes"
	"crypto/tls"
	"github.com/google/uuid"
	"net"
	"net/http"
	"sdk-go/constant"
	"sdk-go/utils"
	"strconv"
	"time"
)

// Client http请求客户端
type Client struct {
	*http.Client
	*BaseInfo
}

func NewClient(info *BaseInfo) (*Client, error) {
	if info == nil {
		return nil, constant.ERR_INIT_CLIENT
	}

	dial := &net.Dialer{
		Timeout:   15 * time.Second,
		KeepAlive: 15 * time.Second,
	}

	transport := http.Transport{
		DialContext:       dial.DialContext,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true,
	}

	c := &Client{
		BaseInfo: info,
		Client: &http.Client{
			Transport: &transport,
			Timeout:   60 * time.Second,
		},
	}

	return c, nil
}

// GetVideoUrl 获取视频流地址
func (c *Client) GetVideoUrl(cameraIndexCode string, streamType, transMode int) (*http.Response, error) {
	path := BasePrefix + VideoUrl

	params := map[string]interface{}{
		"cameraIndexCode": cameraIndexCode,
		"streamType":      streamType,
		"protocol":        "rtsp",
		"transmode":       transMode,
		"expand":          "streamform=rtp&standard=rtsp",
	}

	return c.Do(c.newRequest(http.MethodPost, path, params))
}

func (c *Client) GetSubscribeInfo() (*http.Response, error) {
	path := BasePrefix + VideoUrl

	params := map[string]interface{}{
		"eventTypes": []int64{187675137},
	}

	return c.Do(c.newRequest(http.MethodPost, path, params))
}

func (c *Client) newRequest(method, path string, params map[string]interface{}) *http.Request {
	req, _ := http.NewRequest(method, c.BaseServer+path, bytes.NewBuffer(utils.MapToJson(params)))

	headers := ""
	httpHeader := method + "\n"
	acceptStr := "*/*"
	req.Header.Set("Accept", acceptStr)
	httpHeader += acceptStr + "\n"

	contentTypeStr := "application/json"
	req.Header.Set("Content-Type", contentTypeStr)
	httpHeader += contentTypeStr + "\n"

	req.Header.Set("x-ca-key", c.Key)
	httpHeader += "x-ca-key:" + c.Key + "\n"
	uuidStr := uuid.New().String()
	req.Header.Set("x-ca-nonce", uuidStr)
	httpHeader += "x-ca-nonce:" + uuidStr + "\n"

	timestampStr := strconv.Itoa(int(time.Now().UnixNano() / 1e6))
	req.Header.Set("x-ca-timestamp", timestampStr)
	httpHeader += "x-ca-timestamp:" + timestampStr + "\n" + path
	headers = "x-ca-key,x-ca-nonce,x-ca-timestamp"

	signature := utils.ComputeHmac256(httpHeader, c.Secret)
	req.Header.Set("x-ca-signature-headers", headers)
	req.Header.Set("x-ca-signature", signature)

	//content-md5,

	//glog.Printf("\nhttpHeader:\n%s\nsignature:\n%s\n", httpHeader, signature)

	return req
}
