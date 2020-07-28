package wechat

const (
	AppId          = "wxbdc5610cc59c1631"
	AppSecret      = "appsecret"
	UserInfoURL    = "https://api.weixin.qq.com/sns/userinfo"
	AccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/access_token"
	AuthorizeURL   = "https://open.weixin.qq.com/connect/oauth2/authorize"
	RedirectURL    = "http://127.0.0.1:8080/api/v1/oauth/wechat/callback"
)

type WXBody struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
}

type WXInfo struct {
	Openid     string      `json:"openid"`
	Nickname   interface{} `json:"nickname"`
	City       interface{} `json:"city"`
	Country    interface{} `json:"country"`
	Province   interface{} `json:"province"`
	HeadimgUrl interface{} `json:"headimgurl"`
}

type WXUser struct {
	Id         int    `orm:"column(id);pk;auto"`
	Name       string `orm:"column(name)"`
	CreateTime int64  `orm:"column(create_time)"`
	Openid     string `orm:"column(open_id)"`
	City       string `orm:"column(city)"`
	Country    string `orm:"column(country)"`
	Province   string `orm:"column(province)"`
	HeadimgUrl string `orm:"column(headimg_url)"`
}
