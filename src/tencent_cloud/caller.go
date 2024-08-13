package tencent_cloud

import (
	iai "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iai/v20200303"
	"log"
	"sdk-go/constant"
	"sync"
)

// 本模块封装了腾讯人脸识别调用api

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

// SearchFaces 人脸搜索
func (c *Caller) SearchFaces(where interface{}) (*iai.SearchFacesResponse, error) {
	if c.Client == nil {
		return nil, constant.ERR_INIT_CLIENT
	}
	return c.searchFaces(where)
}

// SearchFacesReturnsByGroup 人脸搜索分库返回
func (c *Caller) SearchFacesReturnsByGroup(where interface{}) (*iai.SearchFacesReturnsByGroupResponse, error) {
	if c.Client == nil {
		return nil, constant.ERR_INIT_CLIENT
	}
	return c.searchFacesReturnsByGroup(where)
}

// DetectLiveFaceAccurate 人脸静态活体检测（高精度版）接口
func (c *Caller) DetectLiveFaceAccurate(where interface{}) (*iai.DetectLiveFaceAccurateResponse, error) {
	if c.Client == nil {
		return nil, constant.ERR_INIT_CLIENT
	}
	return c.detectLiveFaceAccurate(where)
}

// DetectFace 人脸检测与分析相关接口_人脸检测与分析
func (c *Caller) DetectFace(where interface{}) (*iai.DetectFaceResponse, error) {
	if c.Client == nil {
		return nil, constant.ERR_INIT_CLIENT
	}
	return c.detectFace(where)
}

// DetectFaceAttributes 人脸检测与分析相关接口_人脸检测与属性分析
func (c *Caller) DetectFaceAttributes(where interface{}) (*iai.DetectFaceAttributesResponse, error) {
	if c.Client == nil {
		return nil, constant.ERR_INIT_CLIENT
	}
	return c.detectFaceAttributes(where)
}

// CompareFace 人脸比对相关接口_人脸比对
func (c *Caller) CompareFace(where interface{}) (*iai.CompareFaceResponse, error) {
	if c.Client == nil {
		return nil, constant.ERR_INIT_CLIENT
	}
	return c.compareFace(where)
}

// VerifyFace 人脸验证相关接口_人脸验证
func (c *Caller) VerifyFace(where interface{}) (*iai.VerifyFaceResponse, error) {
	if c.Client == nil {
		return nil, constant.ERR_INIT_CLIENT
	}
	return c.verifyFace(where)
}

func (c *Caller) searchFaces(where interface{}) (*iai.SearchFacesResponse, error) {
	resp, err := c.Client.SearchFaces(where)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Caller) searchFacesReturnsByGroup(where interface{}) (*iai.SearchFacesReturnsByGroupResponse, error) {
	resp, err := c.Client.SearchFacesReturnsByGroup(where)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DetectLiveFaceAccurate 人脸静态活体检测（高精度版）接口
func (c *Caller) detectLiveFaceAccurate(where interface{}) (*iai.DetectLiveFaceAccurateResponse, error) {
	resp, err := c.Client.DetectLiveFaceAccurate(where)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DetectFace 人脸检测与分析相关接口_人脸检测与分析
func (c *Caller) detectFace(where interface{}) (*iai.DetectFaceResponse, error) {
	resp, err := c.Client.DetectFace(where)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DetectFaceAttributes 人脸检测与分析相关接口_人脸检测与属性分析
func (c *Caller) detectFaceAttributes(where interface{}) (*iai.DetectFaceAttributesResponse, error) {
	resp, err := c.Client.DetectFaceAttributes(where)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CompareFace 人脸比对相关接口_人脸比对
func (c *Caller) compareFace(where interface{}) (*iai.CompareFaceResponse, error) {
	resp, err := c.Client.CompareFace(where)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// VerifyFace 人脸验证相关接口_人脸验证
func (c *Caller) verifyFace(where interface{}) (*iai.VerifyFaceResponse, error) {
	resp, err := c.Client.VerifyFace(where)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
