package dahua

import (
	"encoding/json"
	"log"
	"sdk-go/constant"
	"sync"
	"time"
)

// 本模块封装了大华api

type Caller struct {
	Client *Client

	Token string
	fn    *time.Timer
	mu    sync.Mutex
	last  time.Time
}

var (
	globalCaller = &Caller{}
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
		globalCaller.getLoginToken()
	})

	if globalCaller.Client == nil {
		log.Fatal(constant.ERR_INIT_CLIENT)
		return nil
	}

	return globalCaller
}

// GetLoginToken 获取登录的token
func (c *Caller) getLoginToken() string {
	now := time.Now()
	if c.Token == "" || now.Sub(c.last) >= time.Second*expiredTime {
		globalCaller.selfCertified()
	}

	return c.Token
}

// SelfCertified 自认证
func (c *Caller) selfCertified() {
	var (
		encryptInfo *EncryptInfoReturn
		authInfo    *AuthDataReturn

		err error
	)

	encryptInfo, err = c.getEncryptInfo()
	if err != nil {
		log.Print("第一次获取加密方式和信息失败：" + err.Error())
		return
	}

	//glog.Println("第一次获取内容：", encryptInfo)

	authInfo, err = c.authLogin(encryptInfo.Realm, encryptInfo.RandomKey, encryptInfo.EncryptType)
	if err != nil {
		log.Print("第二次进行登录鉴权失败：" + err.Error())
		return
	}

	//glog.Println("第二次获取内容：", authInfo)

	_, err = c.keepAliveLogin(authInfo.Token)
	if err != nil {
		log.Print("第三次进行登录信息保活失败：" + err.Error())
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.Token = authInfo.Token
	if c.fn != nil {
		c.fn.Stop()
	}
	c.fn = nil
	c.fn = time.AfterFunc(time.Second*(time.Duration(expiredTime)), c.selfCertified)
}

// getEncryptInfo 第一次获取加密方式和信息
func (c *Caller) getEncryptInfo() (*EncryptInfoReturn, error) {
	resp, err := c.Client.GetEncryptInfo()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var target EncryptInfoReturn

	if err = json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return nil, err
	}

	return &target, nil
}

// authLogin 第二次进行登录鉴权
func (c *Caller) authLogin(realm, randomKey, encryptType string) (*AuthDataReturn, error) {
	resp, err := c.Client.AuthLogin(realm, randomKey, encryptType)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var target AuthDataReturn

	if err = json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return nil, err
	}

	//todo test return
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return nil, err
	//}
	//
	//glog.Println("第二次body内容：", string(body))

	return &target, nil
}

// keepAliveLogin 第三次进行登录信息保活
func (c *Caller) keepAliveLogin(token string) (*KeepAliveDataReturn, error) {
	resp, err := c.Client.KeepAliveLogin(token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var target KeepAliveDataReturn

	if err = json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return nil, err
	}

	return &target, nil
}
