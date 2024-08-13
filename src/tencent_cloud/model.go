package tencent_cloud

const (
	FaceSearch = "iai.tencentcloudapi.com"
)

type BaseInfo struct {
	SecretId  string
	SecretKey string
}

type CommonParam struct {
	Region string `json:"Region" name:"Region"` // 地区
}

// SearchFacesRequestParams 人脸搜索body参数体
type SearchFacesRequestParams struct {
	CommonParam

	// 希望搜索的人员库列表，上限100个。数组元素取值为创建人员库接口中的GroupId。
	// 不可同时搜索不同算法模型版本（FaceModelVersion）的人员库。
	GroupIds []string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url string `json:"Url,omitempty" name:"Url"`

	// 最多识别的人脸数目。默认值为1（仅检测图片中面积最大的那张人脸），最大值为10。
	// MaxFaceNum用于，当输入的待识别图片包含多张人脸时，设定要搜索的人脸的数量。
	// 例如：输入的Image或Url中的图片包含多张人脸，设MaxFaceNum=5，则会识别图片中面积最大的5张人脸。
	MaxFaceNum uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// 人脸长和宽的最小尺寸，单位为像素。默认为34。低于34的人脸图片无法被识别。建议设置为80。
	MinFaceSize uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// 单张被识别的人脸返回的最相似人员数量。默认值为5，最大值为100。
	// 例如，设MaxFaceNum为1，MaxPersonNum为8，则返回Top8相似的人员信息。
	// 值越大，需要处理的时间越长。建议不要超过10。
	MaxPersonNum uint64 `json:"MaxPersonNum,omitempty" name:"MaxPersonNum"`

	// 是否返回人员具体信息。0 为关闭，1 为开启。默认为 0。其他非0非1值默认为0
	NeedPersonInfo int64 `json:"NeedPersonInfo,omitempty" name:"NeedPersonInfo"`

	// 图片质量控制。
	// 0: 不进行控制；
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况；
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况；
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况；
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题；
	// 默认 0。
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 出参Score中，只有超过FaceMatchThreshold值的结果才会返回。默认为0。
	FaceMatchThreshold float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

// SearchFacesReturnsByGroupRequestParams 人脸搜索分库返回body参数体
type SearchFacesReturnsByGroupRequestParams struct {
	CommonParam

	// 希望搜索的人员库列表，上限60个。数组元素取值为创建人员库接口中的GroupId。
	// 不可同时搜索不同算法模型版本（FaceModelVersion）的人员库。
	GroupIds []string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。
	// jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url string `json:"Url,omitempty" name:"Url"`

	// 最多识别的人脸数目。默认值为1（仅检测图片中面积最大的那张人脸），最大值为10。
	// MaxFaceNum用于，当输入的待识别图片包含多张人脸时，设定要搜索的人脸的数量。
	// 例：输入的Image或Url中的图片包含多张人脸，设MaxFaceNum=5，则会识别图片中面积最大的5张人脸。
	MaxFaceNum uint64 `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// 人脸长和宽的最小尺寸，单位为像素。默认为34。低于34将影响搜索精度。建议设置为80。
	MinFaceSize uint64 `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// 被检测到的人脸，对应最多返回的最相似人员数目。默认值为5，最大值为10。
	// 例，设MaxFaceNum为3，MaxPersonNumPerGroup为5，GroupIds长度为3，则最多可能返回3*5*3=45个人员。
	MaxPersonNumPerGroup uint64 `json:"MaxPersonNumPerGroup,omitempty" name:"MaxPersonNumPerGroup"`

	// 是否返回人员具体信息。0 为关闭，1 为开启。默认为 0。其他非0非1值默认为0
	NeedPersonInfo int64 `json:"NeedPersonInfo,omitempty" name:"NeedPersonInfo"`

	// 图片质量控制。
	// 0: 不进行控制；
	// 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况；
	// 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，至少其中三种的情况；
	// 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况；
	// 4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题；
	// 默认 0。
	// 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl uint64 `json:"QualityControl,omitempty" name:"QualityControl"`

	// 出参Score中，只有大于等于FaceMatchThreshold值的结果才会返回。
	// 默认为0。
	// 取值范围[0.0,100.0) 。
	FaceMatchThreshold float64 `json:"FaceMatchThreshold,omitempty" name:"FaceMatchThreshold"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

// DetectLiveFaceAccurateRequestParams 人脸静态活体检测（高精度版）接口
type DetectLiveFaceAccurateRequestParams struct {
	CommonParam

	// 图片 base64 数据，base64 编码后大小不可超过5M。 jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。 图片的宽高比请接近 3:4，手机拍摄比例最佳。
	// 人脸尺寸大于100X100像素。 图片格式支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。 jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。 Url、Image必须提供一个，如果
	// 都提供，只使用 Url。 图片的宽高比请接近 3:4，手机拍摄比例最佳。 人脸尺寸大于100X100像素。 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议
	// 图片存储于腾讯云。非腾讯云存储的Url速度和稳定性可能受一定影响。 图片格式支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url string `json:"Url,omitempty" name:"Url"`

	//人脸识别服务所用的算法模型版本。目前入参支持“3.0“。
	FaceModelVersion string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`
}

// DetectFaceRequestParams 人脸检测与分析相关接口_人脸检测与分析
type DetectFaceRequestParams struct {
	CommonParam

	// 最多处理的人脸数目。默认值为1（仅检测图片中面积最大的那张人脸），最大值为120。 此参数用于控制处理待检测图片中的人脸个数，值越小，处理速度越快。
	MaxFaceNum string `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// 人脸长和宽的最小尺寸，单位为像素，低于MinFaceSize值的人脸不会被检测。 只支持设置34和20，建议使用34。
	MinFaceSize string `json:"MinFaceSize,omitempty" name:"MinFaceSize"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。 支持PNG、JPG、JPEG、BMP，
	// 不支持 GIF 图片。
	Image string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。 jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。 Url、Image必须提供一个，如果
	// 都提供，只使用 Url。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 非腾讯云存储的Url速度和稳定性可能受一定影响。 支持PNG、JPG、JPEG、
	// BMP，不支持 GIF 图片。
	Url string `json:"Url,omitempty" name:"Url"`

	// 是否需要返回人脸属性信息（FaceAttributesInfo）。0 为不需要返回，1 为需要返回。默认为 0。 非 1 值均视为不需要返回，此时 FaceAttributesInfo 不具备参考
	// 意义。
	// 最多返回面积最大的 5 张人脸属性信息，超过 5 张人脸（第 6 张及以后的人脸）的 FaceAttributesInfo 不具备参考意义。
	// 提取人脸属性信息较为耗时，如不需要人脸属性信息，建议关闭此项功能，加快人脸检测速度。
	NeedFaceAttributes string `json:"NeedFaceAttributes,omitempty" name:"NeedFaceAttributes"`

	// 是否开启质量检测。0 为关闭，1 为开启。默认为 0。 非 1 值均视为不进行质量检测。 最多返回面积最大的 30 张人脸质量分信息，超过 30 张人脸（第 31 张及
	// 以后的人脸）的 FaceQualityInfo不具备参考意义。
	// 建议：人脸入库操作建议开启此功能。
	NeedQualityDetection string `json:"NeedQualityDetection,omitempty" name:"NeedQualityDetection"`

	//人脸识别服务所用的算法模型版本。
	//目前入参支持 “2.0”和“3.0“ 两个输入。
	//2020年4月2日开始，默认为“3.0”，之前使用过本接口的账号若未填写本参数默认为“2.0”。
	//2020年11月26日后开通服务的账号仅支持输入“3.0”。
	//不同算法模型版本对应的人脸识别算法不同，新版本的整体效果会优于旧版本，建议使用“3.0”版本。
	FaceModelVersion string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	//是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。
	//若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection string `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

// DetectFaceAttributesRequestParams 人脸检测与分析相关接口_人脸检测与属性分析
type DetectFaceAttributesRequestParams struct {
	CommonParam

	// 最多处理的人脸数目。默认值为1（仅检测图片中面积最大的那张人脸），最大值为120。 此参数用于控制处理待检测图片中的人脸个数，值越小，处理速度越快。
	MaxFaceNum string `json:"MaxFaceNum,omitempty" name:"MaxFaceNum"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。 支持PNG、JPG、JPEG、BMP，
	// 不支持 GIF 图片。
	Image string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。 jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。 Url、Image必须提供一个，如果
	// 都提供，只使用 Url。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 非腾讯云存储的Url速度和稳定性可能受一定影响。 支持PNG、JPG、JPEG、
	// BMP，不支持 GIF 图片。
	Url string `json:"Url,omitempty" name:"Url"`

	// 是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。
	//若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection uint64 `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`

	//人脸识别服务所用的算法模型版本。目前入参支持“3.0“。
	FaceModelVersion string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	//是否返回年龄、性别、情绪等属性。 合法值为（大小写不敏感）：None、Age、Beauty、Emotion、Eye、Eyebrow、 Gender、Hair、Hat、Headpose、Mask、Mouth、Moustache、Nose、Shape、Skin、Smile。
	//None为不需要返回。默认为 None。即FaceAttributesType属性为空时，各属性返回值为0。 需要将属性组成一个用逗号分隔的字符串，属性之间的顺序没有要求。
	//关于各属性的详细描述，参见下文出参。 最多返回面积最大的 5 张人脸属性信息，超过 5 张人脸（第 6 张及以后的人脸）的 AttributesInfo 不具备参考意义。
	FaceAttributesType string `json:"FaceAttributesType,omitempty" name:"FaceAttributesType"`
}

// CompareFaceRequestParams 人脸比对相关接口_人脸比对
type CompareFaceRequestParams struct {
	CommonParam

	//A 图片 base64 数据，base64 编码后大小不可超过5M。 jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。
	//若图片中包含多张人脸，只选取其中人脸面积最大的人脸。 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	ImageA string `json:"ImageA,omitempty" name:"ImageA"`

	//A 图片 base64 数据，base64 编码后大小不可超过5M。 jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。
	//若图片中包含多张人脸，只选取其中人脸面积最大的人脸。 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	ImageB string `json:"ImageB,omitempty" name:"ImageB"`

	//A 图片的 Url ，对应图片 base64 编码后大小不可超过5M。 jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。 A 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。
	//图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	//非腾讯云存储的Url速度和稳定性可能受一定影响。 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	UrlA string `json:"UrlA,omitempty" name:"UrlA"`

	//A 图片的 Url ，对应图片 base64 编码后大小不可超过5M。 jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。 A 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。
	//图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	//非腾讯云存储的Url速度和稳定性可能受一定影响。 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	UrlB string `json:"UrlB,omitempty" name:"UrlB"`

	//人脸识别服务所用的算法模型版本。目前入参支持“3.0“。
	FaceModelVersion string `json:"FaceModelVersion,omitempty" name:"FaceModelVersion"`

	//图片质量控制。 0: 不进行控制； 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，
	//下巴遮挡，至少其中三种的情况； 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况；
	//4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 默认 0。 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl string `json:"QualityControl,omitempty" name:"QualityControl"`

	//是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。
	//若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection string `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}

// VerifyFaceRequestParams 人脸验证相关接口_人脸验证
type VerifyFaceRequestParams struct {
	CommonParam

	// 图片 base64 数据，base64 编码后大小不可超过5M。 jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。 支持PNG、JPG、JPEG、BMP，
	// 不支持 GIF 图片。
	Image string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。 jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。 Url、Image必须提供一个，如果
	// 都提供，只使用 Url。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 非腾讯云存储的Url速度和稳定性可能受一定影响。 支持PNG、JPG、JPEG、
	// BMP，不支持 GIF 图片。
	Url string `json:"Url,omitempty" name:"Url"`

	// 待验证的人员ID。人员ID具体信息请参考人员库管理相关接口。
	PersonId string `json:"PersonId,omitempty" name:"PersonId"`

	//图片质量控制。 0: 不进行控制； 1:较低的质量要求，图像存在非常模糊，眼睛鼻子嘴巴遮挡至少其中一种或多种的情况； 2: 一般的质量要求，图像存在偏亮，偏暗，模糊或一般模糊，眉毛遮挡，脸颊遮挡，
	//下巴遮挡，至少其中三种的情况； 3: 较高的质量要求，图像存在偏亮，偏暗，一般模糊，眉毛遮挡，脸颊遮挡，下巴遮挡，其中一到两种的情况；
	//4: 很高的质量要求，各个维度均为最好或最多在某一维度上存在轻微问题； 默认 0。 若图片质量不满足要求，则返回结果中会提示图片质量检测不符要求。
	QualityControl string `json:"QualityControl,omitempty" name:"QualityControl"`

	//是否开启图片旋转识别支持。0为不开启，1为开启。默认为0。本参数的作用为，当图片中的人脸被旋转且图片没有exif信息时，如果不开启图片旋转识别支持则无法正确检测、识别图片中的人脸。
	//若您确认图片包含exif信息或者您确认输入图中人脸不会出现被旋转情况，请不要开启本参数。开启后，整体耗时将可能增加数百毫秒。
	NeedRotateDetection string `json:"NeedRotateDetection,omitempty" name:"NeedRotateDetection"`
}
