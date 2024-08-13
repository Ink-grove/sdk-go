package hk

import (
	"encoding/json"
	"log"
	"sdk-go/constant"
	"sync"
)

// 本模块封装了海康api

type Caller struct {
	Client *Client
}

var (
	globalCaller *Caller
	one          sync.Once
)

func InitCaller(info *BaseInfo) *Caller {
	one.Do(func() {
		client, err := NewClient(info)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		globalCaller.Client = client

	})
	return globalCaller
}

// GetVideoUrl 获取视频流地址
func (c *Caller) GetVideoUrl(cameraIndexCode string, streamType, transMode int) (*VideoData, error) {
	if c.Client == nil {
		return nil, constant.ERR_INIT_CLIENT
	}

	return c.getVideoUrl(cameraIndexCode, streamType, transMode)
}

// GetSubscribeInfo 获取事件订阅信息
func (c *Caller) GetSubscribeInfo() (*SubscribeInfoData, error) {
	if c.Client == nil {
		return nil, constant.ERR_INIT_CLIENT
	}

	return c.getSubscribeInfo()
}

func (c *Caller) getVideoUrl(cameraIndexCode string, streamType, transMode int) (*VideoData, error) {
	resp, err := c.Client.GetVideoUrl(cameraIndexCode, streamType, transMode)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var target VideoData

	if err = json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return nil, err
	}

	return &target, nil
}

func (c *Caller) getSubscribeInfo() (*SubscribeInfoData, error) {
	resp, err := c.Client.GetSubscribeInfo()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var target SubscribeInfoData

	if err = json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return nil, err
	}

	return &target, nil
}
