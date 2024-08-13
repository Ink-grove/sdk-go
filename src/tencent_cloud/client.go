package tencent_cloud

import (
	tcm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	iai "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iai/v20200303"
	"sdk-go/constant"
)

// Client http请求客户端
type Client struct {
	*tcm.Credential
	*profile.ClientProfile
	*BaseInfo
}

func NewClient(info *BaseInfo) (*Client, error) {
	if info == nil {
		return nil, constant.ERR_INIT_CLIENT
	}

	c := &Client{
		Credential:    tcm.NewCredential(info.SecretId, info.SecretKey),
		ClientProfile: profile.NewClientProfile(),
		BaseInfo:      info,
	}

	return c, nil
}

func (c *Client) SearchFaces(where interface{}) (*iai.SearchFacesResponse, error) {
	if p := where.(*SearchFacesRequestParams); p == nil {
		return nil, constant.ERR_PARAMS_INPUT
	} else {
		// 实例化要请求产品的client对象,clientProfile是可选的
		client, _ := iai.NewClient(c.Credential, p.Region, c.creatProfile(FaceSearch))

		// 实例化一个请求对象,每个接口都会对应一个request对象
		request := iai.NewSearchFacesRequest()
		if len(p.Url) <= 2 {
			request.Image = &p.Image
		} else {
			request.Url = &p.Url
		}
		request.GroupIds = tcm.StringPtrs(p.GroupIds)
		request.MaxFaceNum = &p.MaxFaceNum
		request.MaxPersonNum = &p.MaxPersonNum

		// 返回的resp是一个DescribeRegionsResponse的实例，与请求对象对应
		return client.SearchFaces(request)
	}
}

func (c *Client) SearchFacesReturnsByGroup(where interface{}) (*iai.SearchFacesReturnsByGroupResponse, error) {
	if p := where.(*SearchFacesReturnsByGroupRequestParams); p == nil {
		return nil, constant.ERR_PARAMS_INPUT
	} else {
		// 实例化要请求产品的client对象,clientProfile是可选的
		client, _ := iai.NewClient(c.Credential, p.Region, c.creatProfile(FaceSearch))

		// 实例化一个请求对象,每个接口都会对应一个request对象
		request := iai.NewSearchFacesReturnsByGroupRequest()
		if len(p.Url) <= 2 {
			request.Image = &p.Image
		} else {
			request.Url = &p.Url
		}
		request.GroupIds = tcm.StringPtrs(p.GroupIds)
		request.MaxFaceNum = &p.MaxFaceNum

		// 返回的resp是一个DescribeRegionsResponse的实例，与请求对象对应
		return client.SearchFacesReturnsByGroup(request)
	}
}

func (c *Client) creatProfile(path string) *profile.ClientProfile {
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	c.ClientProfile.HttpProfile.Endpoint = path
	return c.ClientProfile
}

// DetectLiveFaceAccurate 人脸静态活体检测（高精度版）接口
func (c *Client) DetectLiveFaceAccurate(where interface{}) (*iai.DetectLiveFaceAccurateResponse, error) {
	if p := where.(*DetectLiveFaceAccurateRequestParams); p == nil {
		return nil, constant.ERR_PARAMS_INPUT
	} else {
		// 实例化要请求产品的client对象,clientProfile是可选的
		client, _ := iai.NewClient(c.Credential, where.(*DetectLiveFaceAccurateRequestParams).Region, c.creatProfile(FaceSearch))
		// 实例化一个请求对象,每个接口都会对应一个request对象
		request := iai.NewDetectLiveFaceAccurateRequest()
		if len(p.Url) <= 2 {
			request.Image = &p.Image
		} else {
			request.Url = &p.Url
		}
		return client.DetectLiveFaceAccurate(request)
	}

}

// DetectFace 人脸检测与分析相关接口_人脸检测与分析
func (c *Client) DetectFace(where interface{}) (*iai.DetectFaceResponse, error) {
	if p := where.(*DetectFaceRequestParams); p == nil {
		return nil, constant.ERR_PARAMS_INPUT
	} else {
		// 实例化要请求产品的client对象,clientProfile是可选的
		client, _ := iai.NewClient(c.Credential, where.(*DetectFaceRequestParams).Region, c.creatProfile(FaceSearch))
		// 实例化一个请求对象,每个接口都会对应一个request对象
		request := iai.NewDetectFaceRequest()
		if len(p.Url) <= 2 {
			request.Image = &p.Image
		} else {
			request.Url = &p.Url
		}
		return client.DetectFace(request)
	}

}

// DetectFaceAttributes 人脸检测与分析相关接口_人脸检测与属性分析
func (c *Client) DetectFaceAttributes(where interface{}) (*iai.DetectFaceAttributesResponse, error) {
	if p := where.(*DetectFaceAttributesRequestParams); p == nil {
		return nil, constant.ERR_PARAMS_INPUT
	} else {
		// 实例化要请求产品的client对象,clientProfile是可选的
		client, _ := iai.NewClient(c.Credential, where.(*DetectFaceAttributesRequestParams).Region, c.creatProfile(FaceSearch))
		// 实例化一个请求对象,每个接口都会对应一个request对象
		request := iai.NewDetectFaceAttributesRequest()
		if len(p.Url) <= 2 {
			request.Image = &p.Image
		} else {
			request.Url = &p.Url
		}
		return client.DetectFaceAttributes(request)
	}
}

// CompareFace 人脸比对相关接口_人脸比对
func (c *Client) CompareFace(where interface{}) (*iai.CompareFaceResponse, error) {
	if p := where.(*CompareFaceRequestParams); p == nil {
		return nil, constant.ERR_PARAMS_INPUT
	} else {
		// 实例化要请求产品的client对象,clientProfile是可选的
		client, _ := iai.NewClient(c.Credential, where.(*CompareFaceRequestParams).Region, c.creatProfile(FaceSearch))
		// 实例化一个请求对象,每个接口都会对应一个request对象
		request := iai.NewCompareFaceRequest()
		//if len(p.Url) <= 2 {
		//	request.Image = &p.Image
		//} else {
		//	request.Url = &p.Url
		//}
		request.ImageA = &p.ImageA
		request.ImageB = &p.ImageB
		return client.CompareFace(request)
	}
}

// VerifyFace 人脸验证相关接口_人脸验证
func (c *Client) VerifyFace(where interface{}) (*iai.VerifyFaceResponse, error) {
	if p := where.(*VerifyFaceRequestParams); p == nil {
		return nil, constant.ERR_PARAMS_INPUT
	} else {
		// 实例化要请求产品的client对象,clientProfile是可选的
		client, _ := iai.NewClient(c.Credential, where.(*VerifyFaceRequestParams).Region, c.creatProfile(FaceSearch))
		// 实例化一个请求对象,每个接口都会对应一个request对象
		request := iai.NewVerifyFaceRequest()
		if len(p.Url) <= 2 {
			request.Image = &p.Image
		} else {
			request.Url = &p.Url
		}
		request.PersonId = &p.PersonId
		return client.VerifyFace(request)
	}
}
