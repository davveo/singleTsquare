package oauth2

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
	QQAppId        = "101827468"
	QQAppKey       = "0d2d856e48e0ebf6b98e0d0c879fe74d"
	QQRedirectURL  = "http://127.0.0.1:8080/api/v1/oauth/qq/callback" // TODO host动态获取
	QQTokenURL     = "https://graph.qq.com/oauth2.0/token"
	QQAuthorizeURL = "https://graph.qq.com/oauth2.0/authorize"
	QQOpenidUrl    = "https://graph.qq.com/oauth2.0/me"
	QQUserInfoURL  = "https://graph.qq.com/user/get_user_info"
)

type PrivateInfo struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
}

type UnmarshalQQUserInfo struct {
	NickName string `json:"nickname"`
	Avatar   string `json:"figureurl_qq_2"`
}

type qqService struct {
	PlatformType uint
}

func NewQQService() *qqService {
	return &qqService{PlatformType: 1}
}

func (s *qqService) RequestAccessToken(code string) (*PrivateInfo, error) {
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", QQAppId)
	params.Add("client_secret", QQAppKey)
	params.Add("code", code)
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), QQRedirectURL)
	loginUrl := fmt.Sprintf("%s?%s", QQTokenURL, str)
	response, err := http.Get(loginUrl)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bs, _ := ioutil.ReadAll(response.Body)
	body := string(bs)

	resultMap := s.convertToMap(body)

	info := &PrivateInfo{}
	info.AccessToken = resultMap["access_token"]
	info.RefreshToken = resultMap["refresh_token"]
	info.ExpiresIn = resultMap["expires_in"]

	return info, nil
}

func (s *qqService) GetOpenId(accessToken string) (string, error) {
	resp, err := http.Get(fmt.Sprintf(
		"%s?access_token=%s", QQOpenidUrl, accessToken))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	return string(bs)[45:77], nil
}

func (s *qqService) GetUserInfo(code string) (*UserInfo, error) {
	var unmarshaluserInfo *UnmarshalQQUserInfo
	tokenInfo, err := s.RequestAccessToken(code)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	accessToken := tokenInfo.AccessToken
	openid, _ := s.GetOpenId(accessToken)
	params.Add("openid", openid)
	params.Add("oauth_consumer_key", QQAppId)
	params.Add("access_token", accessToken)

	resp, err := http.Get(fmt.Sprintf(
		"%s?%s", QQUserInfoURL, params.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bs, &unmarshaluserInfo)
	if err != nil {
		return nil, err
	}
	// 将用户的标示写入

	return &UserInfo{
		OpenId:      openid,
		AccessToken: accessToken,
		NickName:    unmarshaluserInfo.NickName,
		Avatar:      unmarshaluserInfo.Avatar,
	}, nil
}

func (s *qqService) convertToMap(str string) map[string]string {
	var resultMap = make(map[string]string)
	values := strings.Split(str, "&")
	for _, value := range values {
		vs := strings.Split(value, "=")
		resultMap[vs[0]] = vs[1]
	}
	return resultMap
}

func (s *qqService) GenRedirectURL() string {
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", QQAppId)
	params.Add("state", "test")
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), QQRedirectURL)
	loginURL := fmt.Sprintf("%s?%s", QQAuthorizeURL, str)
	return loginURL
}

func (s *qqService) GetPlatformType() uint {
	return s.PlatformType
}
