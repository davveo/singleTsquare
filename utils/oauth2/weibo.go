package oauth2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	WeiboAppId        = "101827468"
	WeiboAppKey       = "0d2d856e48e0ebf6b98e0d0c879fe74d"
	WeiboRedirectURL  = "http://127.0.0.1:8080/api/v1/oauth/weibo/callback" // TODO host动态获取
	WeiboTokenURL     = "https://api.weibo.com/oauth2/access_token"
	WeiboAuthorizeURL = "https://api.weibo.com/oauth2/authorize"
	WeiboUserInfoURL  = "https://api.weibo.com/2/users/show.json"
	WeiboOpenIdURL    = "https://api.weibo.com/oauth2/get_token_info"
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

type weiboService struct {
	PlatformType uint
}

func NewWeiboService() *weiboService {
	return &weiboService{PlatformType: 3}
}

func (s *weiboService) TokenParams(code string) string {
	params := url.Values{}
	params.Add("code", code)
	params.Add("client_id", WeiboAppId)
	params.Add("client_secret", WeiboAppKey)
	params.Add("grant_type", "authorization_code")
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), WeiboRedirectURL)
	return fmt.Sprintf("%s?%s", WeiboTokenURL, str)
}

func (s *weiboService) GetToken(code string) (*AccessTokenInfo, error) {
	var accessTokenInfo *AccessTokenInfo
	loginUrl := s.TokenParams(code)
	response, err := http.Get(loginUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bs, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bs, accessTokenInfo)

	return accessTokenInfo, nil
}

func (s *weiboService) GetUserInfo(code string) (*UserInfo, error) {
	var unmarshalUserInfo *UnmarshalUserInfo
	accessTokenInfo, err := s.GetToken(code)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("access_token", accessTokenInfo.AccessToken)
	params.Add("uid", accessTokenInfo.Uid)
	uri := fmt.Sprintf("%s?%s", WeiboUserInfoURL, params.Encode())
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bs, &unmarshalUserInfo)
	if err != nil {
		return nil, err
	}
	return &UserInfo{
		OpenId:      accessTokenInfo.Uid,
		Avatar:      unmarshalUserInfo.Avatar,
		NickName:    unmarshalUserInfo.NickName,
		AccessToken: accessTokenInfo.AccessToken,
	}, nil
}

func (s *weiboService) GenRedirectURL() string {
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", WeiboAppId)
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), WeiboRedirectURL)
	loginURL := fmt.Sprintf("%s?%s", WeiboAuthorizeURL, str)
	return loginURL
}

func (s *weiboService) GetPlatformType() uint {
	return s.PlatformType
}
