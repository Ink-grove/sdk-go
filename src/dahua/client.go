package dahua

import (
	"crypto/tls"
	"net"
	"net/http"
	"sdk-go/constant"
	"sdk-go/utils"
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

// GetEncryptInfo 第一次获取加密方式和信息
func (c *Client) GetEncryptInfo() (*http.Response, error) {
	path := c.BaseServer + AuthUrl

	content := map[string]interface{}{
		"userName":   c.Key,
		"clientType": "winpc",
		//"expiredTime": expiredTime,
	}

	buffer, _ := utils.ValueToBuffer(content)
	req, _ := http.NewRequest(http.MethodPost, path, buffer)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	return c.Do(req)
}

// AuthLogin 第二次进行登录鉴权
func (c *Client) AuthLogin(realm, randomKey, encryptType string) (*http.Response, error) {
	path := c.BaseServer + AuthUrl

	content := map[string]interface{}{
		"userName":    c.Key,
		"signature":   c.signature(realm, randomKey),
		"randomKey":   randomKey,
		"encryptType": encryptType,
		"clientType":  "winpc",
		//"expiredTime": expiredTime,
	}

	buffer, _ := utils.ValueToBuffer(content)
	req, _ := http.NewRequest(http.MethodPost, path, buffer)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	return c.Do(req)
}

// KeepAliveLogin 保活token
func (c *Client) KeepAliveLogin(token string) (*http.Response, error) {
	path := c.BaseServer + TokenUrl

	content := map[string]interface{}{
		"token": token,
	}

	buffer, _ := utils.ValueToBuffer(content)
	req, _ := http.NewRequest(http.MethodPut, path, buffer)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("X-Subject-Token", token)
	return c.Do(req)
}

func (c *Client) signature(realm, randomKey string) string {
	s := utils.Md5Encrypt(c.Secret)
	s = utils.Md5Encrypt(c.Key + s)
	s = utils.Md5Encrypt(s)
	s = utils.Md5Encrypt(c.Key + ":" + realm + ":" + s)
	return utils.Md5Encrypt(s + ":" + randomKey)
}
