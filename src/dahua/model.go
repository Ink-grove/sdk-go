package dahua

const (
	AuthUrl  = "/videoService/accounts/authorize"
	TokenUrl = "/videoService/accounts/token/keepalive"

	expiredTime = 50 // 每50s刷新一次token
)

type BaseInfo struct {
	Key        string
	Secret     string
	BaseServer string
}

type EncryptInfoReturn struct {
	RandomKey   string `json:"randomKey"`
	Realm       string `json:"realm"`
	EncryptType string `json:"encryptType"`
	Method      string `json:"method"`
}

type AuthDataReturn struct {
	Duration      int      `json:"duration"`
	Token         string   `json:"token"`
	UserName      string   `json:"userName"`
	UserId        string   `json:"userId"`
	ServiceAbilty []string `json:"serviceAbilty"`
	LastLoginIp   string   `json:"lastLoginIp"`
	VersionInfo   struct {
		LastVersion string `json:"lastVersion"`
		UpdateUrl   string `json:"updateUrl"`
	} `json:"versionInfo"`
}

type KeepAliveDataReturn struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
