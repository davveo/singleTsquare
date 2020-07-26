package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	AuthorizeURL = "https://github.com/login/oauth/authorize"
	AppId        = "101827468"
	RedirectURI  = "http://127.0.0.1:8080/oauth2/"
	TokenURL     = "https://github.com/login/oauth/access_token"
	UserInfoURL  = "https://api.github.com/user"
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

// 认证并获取用户信息
func Oauth(code string) (userInfo map[string]interface{}, err error) {
	// 通过 code, 获取 token
	var token *Token
	var tokenAuthUrl = GetTokenAuthUrl(code)
	if token, err = GetToken(tokenAuthUrl); err != nil {
		return nil, errors.New(fmt.Sprintf("获取Token失败，错误信息为:%s", err))
	}

	// 通过token，获取用户信息
	userInfo, err = GetUserInfo(token)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("获取用户信息失败，错误信息为:%s", err))
	}
	return userInfo, nil
}

// 通过code获取token认证url
func GetTokenAuthUrl(code string) string {
	return fmt.Sprintf(
		"%s?client_id=%s&client_secret=%s&code=%s",
		TokenURL, conf.ClientId, conf.ClientSecret, code,
	)
}

// 获取 token
func GetToken(code string) (*Token, error) {
	// 形成请求
	var req *http.Request
	var err error
	if req, err = http.NewRequest(
		http.MethodGet,
		GetTokenAuthUrl(code), nil); err != nil {
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
	var token Token
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}
	return &token, nil
}

// 获取用户信息
func GetUserInfo(token *Token) (map[string]interface{}, error) {

	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, UserInfoURL, nil); err != nil {
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

	// 将响应的数据写入 userInfo 中，并返回
	var userInfo = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}

func GenRedirectURL() string {
	// 跳转到第三方地址
	params := url.Values{}
	params.Add("client_id", AppId)
	params.Add("state", "state")
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), RedirectURI)
	loginURL := fmt.Sprintf("%s?%s", AuthorizeURL, str)
	return loginURL
}
