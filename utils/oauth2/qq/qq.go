package qq

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// TODO 配置的处理
const (
	AppId        = "101827468"
	AppKey       = "0d2d856e48e0ebf6b98e0d0c879fe74d"
	RedirectURL  = "http://127.0.0.1:9090/api/v1/qqLogin" // TODO host动态获取
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

type UserInfo struct {
	NickName    string `json:"nickname"`
	OpenId      string `json:"openid"`
	Avatar      string `json:"figureurl_qq_2"`
	AccessToken string `json:"access_token"`
}

func RequestAccessToken(code string) (*PrivateInfo, error) {
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", AppId)
	params.Add("client_secret", AppKey)
	params.Add("code", code)
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), RedirectURL)
	loginUrl := fmt.Sprintf("%s?%s", TokenURL, str)
	response, err := http.Get(loginUrl)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bs, _ := ioutil.ReadAll(response.Body)
	body := string(bs)

	resultMap := convertToMap(body)

	info := &PrivateInfo{}
	info.AccessToken = resultMap["access_token"]
	info.RefreshToken = resultMap["refresh_token"]
	info.ExpiresIn = resultMap["expires_in"]

	return info, nil
}

func GetOpenId(accessToken string) (string, error) {
	resp, err := http.Get(fmt.Sprintf(
		"%s?access_token=%s", OpenidUrl, accessToken))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	return string(bs)[45:77], nil
}

func GetUserInfo(code string) (*UserInfo, error) {
	var userInfo *UserInfo
	tokenInfo, err := RequestAccessToken(code)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	accessToken := tokenInfo.AccessToken
	openid, _ := GetOpenId(accessToken)
	params.Add("openid", openid)
	params.Add("oauth_consumer_key", AppId)
	params.Add("access_token", accessToken)

	resp, err := http.Get(fmt.Sprintf(
		"%s?%s", UserInfoURL, params.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bs, &userInfo)
	if err != nil {
		return nil, err
	}
	// 将用户的标示写入
	userInfo.OpenId = openid
	userInfo.AccessToken = accessToken

	return userInfo, nil
}

func convertToMap(str string) map[string]string {
	var resultMap = make(map[string]string)
	values := strings.Split(str, "&")
	for _, value := range values {
		vs := strings.Split(value, "=")
		resultMap[vs[0]] = vs[1]
	}
	return resultMap
}

func GenRedirectURL() string {
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", AppId)
	params.Add("state", "test")
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), RedirectURL)
	loginURL := fmt.Sprintf("%s?%s", AuthorizeURL, str)
	return loginURL
}
