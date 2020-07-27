package weibo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// TODO 配置的处理
const (
	AppId        = "101827468"
	AppKey       = "0d2d856e48e0ebf6b98e0d0c879fe74d"
	RedirectURL  = "http://127.0.0.1:9090/api/v1/wbLogin" // TODO host动态获取
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

type UserInfo struct {
	NickName    string `json:"name"`
	OpenId      string `json:"openid"`
	Avatar      string `json:"avatar_large"`
	AccessToken string `json:"access_token"`
}

func TokenParams(code string) string {
	params := url.Values{}
	params.Add("code", code)
	params.Add("client_id", AppId)
	params.Add("client_secret", AppKey)
	params.Add("grant_type", "authorization_code")
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), RedirectURL)
	return fmt.Sprintf("%s?%s", TokenURL, str)
}

func GetToken(code string) (*AccessTokenInfo, error) {
	var accessTokenInfo *AccessTokenInfo
	loginUrl := TokenParams(code)
	response, err := http.Get(loginUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bs, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bs, accessTokenInfo)

	return accessTokenInfo, nil
}

func GetUserInfo(code string) (*UserInfo, error) {
	var userInfo *UserInfo
	accessTokenInfo, err := GetToken(code)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("access_token", accessTokenInfo.AccessToken)
	params.Add("uid", accessTokenInfo.Uid)
	uri := fmt.Sprintf("%s?%s", UserInfoURL, params.Encode())
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bs, &userInfo)
	if err != nil {
		return nil, err
	}

	userInfo.OpenId = accessTokenInfo.Uid
	userInfo.AccessToken = accessTokenInfo.AccessToken
	return userInfo, nil
}

func GenRedirectURL() string {
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", AppId)
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), RedirectURL)
	loginURL := fmt.Sprintf("%s?%s", AuthorizeURL, str)
	return loginURL
}
