package oauth2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	GithubAppId        = "101827468"
	GithubUserInfoURL  = "https://api.github.com/user"
	GithubRedirectURI  = "http://127.0.0.1:8080/api/v1/oauth/github/callback"
	GithubAuthorizeURL = "https://github.com/login/oauth/authorize"
	GithubTokenURL     = "https://github.com/login/oauth/access_token"
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

type GithubToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段下面没用到
	Scope       string `json:"scope"`      // 这个字段下面也没用到
}

type UnmarshalGithubUserInfo struct {
	ID       string `json:"id"`
	NickName string `json:"name"`
	Avatar   string `json:"avatar_url"`
}

type githubService struct {
	PlatformType uint
}

func NewGithubService() *githubService {
	return &githubService{PlatformType: 4}
}

// 通过code获取token认证url
func (s *githubService) GetTokenAuthUrl(code string) string {
	return fmt.Sprintf(
		"%s?client_id=%s&client_secret=%s&code=%s",
		GithubTokenURL, conf.ClientId, conf.ClientSecret, code,
	)
}

// 获取 token
func (s *githubService) GetToken(code string) (*GithubToken, error) {
	// 形成请求
	var req *http.Request
	var err error
	if req, err = http.NewRequest(
		http.MethodGet,
		s.GetTokenAuthUrl(code), nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}

	// 将响应体解析为 token，并返回
	var token GithubToken
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *githubService) GenRedirectURL() string {
	// 跳转到第三方地址
	params := url.Values{}
	params.Add("client_id", GithubAppId)
	params.Add("state", "state")
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), GithubRedirectURI)
	loginURL := fmt.Sprintf("%s?%s", GithubAuthorizeURL, str)
	return loginURL
}

func (s *githubService) GetUserInfo(code string) (*UserInfo, error) {
	var unmarshalUserInfo *UnmarshalGithubUserInfo
	var req *http.Request
	var err error
	// 获取token
	token, err := s.GetToken(code)
	if err != nil {
		return nil, err
	}

	if req, err = http.NewRequest(http.MethodGet, GithubUserInfoURL, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))

	// 发送请求并获取响应
	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, err
	}

	defer res.Body.Close()
	bs, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(bs, &unmarshalUserInfo)
	if err != nil {
		return nil, err
	}

	return &UserInfo{
		AccessToken: token.AccessToken,
		OpenId:      unmarshalUserInfo.ID,
		Avatar:      unmarshalUserInfo.Avatar,
		NickName:    unmarshalUserInfo.NickName,
	}, nil
}

func (s *githubService) GetPlatformType() uint {
	return s.PlatformType
}
