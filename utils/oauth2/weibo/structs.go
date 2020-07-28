package weibo

// TODO 配置的处理
const (
	AppId        = "101827468"
	AppKey       = "0d2d856e48e0ebf6b98e0d0c879fe74d"
	RedirectURL  = "http://127.0.0.1:8080/api/v1/oauth/weibo/callback" // TODO host动态获取
	TokenURL     = "https://api.weibo.com/oauth2/access_token"
	AuthorizeURL = "https://api.weibo.com/oauth2/authorize"
	UserInfoURL  = "https://api.weibo.com/2/users/show.json"
	OpenIdURL    = "https://api.weibo.com/oauth2/get_token_info"
)

type AccessTokenInfo struct {
	AccessToken string `json:"access_token"`
	RemindIn    int    `json:"remind_in"`
	ExpiresIn   int    `json:"expires_in"`
	Uid         string `json:"uid"`
}

type UnmarshalUserInfo struct {
	NickName string `json:"name"`
	Avatar   string `json:"avatar_large"`
}
