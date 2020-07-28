package github

const (
	AppId        = "101827468"
	UserInfoURL  = "https://api.github.com/user"
	RedirectURI  = "http://127.0.0.1:8080/api/v1/oauth/github/callback"
	AuthorizeURL = "https://github.com/login/oauth/authorize"
	TokenURL     = "https://github.com/login/oauth/access_token"
)

type Conf struct {
	ClientId     string // 对应: Client ID
	ClientSecret string // 对应: Client Secret
	RedirectUrl  string // 对应: Authorization callback URL
}

var conf = Conf{
	ClientId:     "7e5fe351bc9b131c6f2a",
	ClientSecret: "9fd22c13ae790685c59e3fb4a9b444b75b506a5b",
	RedirectUrl:  "http://localhost:9090/oauth/redirect",
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段下面没用到
	Scope       string `json:"scope"`      // 这个字段下面也没用到
}

type UnmarshalUserInfo struct {
	ID       string `json:"id"`
	NickName string `json:"name"`
	Avatar   string `json:"avatar_url"`
}
