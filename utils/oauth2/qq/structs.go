package qq

// TODO 配置的处理
const (
	AppId        = "101827468"
	AppKey       = "0d2d856e48e0ebf6b98e0d0c879fe74d"
	RedirectURL  = "http://127.0.0.1:8080/api/v1/oauth/qq/callback" // TODO host动态获取
	TokenURL     = "https://graph.qq.com/oauth2.0/token"
	AuthorizeURL = "https://graph.qq.com/oauth2.0/authorize"
	OpenidUrl    = "https://graph.qq.com/oauth2.0/me"
	UserInfoURL  = "https://graph.qq.com/user/get_user_info"
)

type PrivateInfo struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
}

type UnmarshalUserInfo struct {
	NickName string `json:"nickname"`
	Avatar   string `json:"figureurl_qq_2"`
}
