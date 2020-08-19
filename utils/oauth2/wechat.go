package oauth2

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// https://blog.csdn.net/qq_35781732/article/details/82662021

const (
	WechatAppId          = "wxbdc5610cc59c1631"
	WechatAppSecret      = "appsecret"
	WechatUserInfoURL    = "https://api.weixin.qq.com/sns/userinfo"
	WechatAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/access_token"
	WechatAuthorizeURL   = "https://open.weixin.qq.com/connect/oauth2/authorize"
	WechatRedirectURL    = "http://127.0.0.1:8080/api/v1/oauth/wechat/callback"
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

type wechatService struct {
	PlatformType uint
}

func NewWechatService() *wechatService {
	return &wechatService{PlatformType: 2}
}

func (s *wechatService) GenRedirectURL() string {
	params := url.Values{}
	params.Add("appid", WechatAppId)
	params.Add("state", "login")
	params.Add("response_type", "code")
	params.Add("scope", "snsapi_userinfo")
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), WechatRedirectURL)
	loginURL := fmt.Sprintf("%s?%s", WechatAuthorizeURL, str)
	return loginURL
}

/*
{
	"access_token":"ACCESS_TOKEN",
	"expires_in":7200,
	"refresh_token":"REFRESH_TOKEN",
	"openid":"OPENID",
	"scope":"SCOPE",
	"unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL"
}
	access_token的超时时间是32分钟
*/
func (s *wechatService) GetAccessToken(code string) (*WXBody, error) {
	params := url.Values{}
	params.Add("code", code)
	params.Add("appid", WechatAppId)
	params.Add("secret", WechatAppSecret)
	params.Add("grant_type", "authorization_code")
	resp, err := http.Get(fmt.Sprintf("%s?%s", WechatAccessTokenURL, params.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if !bytes.Contains(body, []byte("access_token")) {
		return nil, errors.New("get access_token fail")
	}

	atr := WXBody{}
	err = json.Unmarshal(body, &atr)
	if err != nil {
		return nil, err
	} else {
		return &atr, nil
	}
}

func (s *wechatService) GetUserInfo(code string) (*UserInfo, error) {
	token, err := s.GetAccessToken(code)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("lang", "zh_CN")
	params.Add("openid", token.Openid)
	params.Add("access_token", token.AccessToken)
	infoBody, err := http.Get(fmt.Sprintf("%s?%s", WechatUserInfoURL, params.Encode()))

	defer infoBody.Body.Close()
	info, _ := ioutil.ReadAll(infoBody.Body)

	updateUser := WXInfo{}
	err = json.Unmarshal(info, &updateUser)
	if err != nil {
		return nil, err
	}
	return &UserInfo{
		OpenId:      token.Openid,
		AccessToken: token.AccessToken,
		Avatar:      updateUser.HeadimgUrl.(string),
		NickName:    updateUser.Nickname.(string),
	}, nil
}

func (s *wechatService) GetPlatformType() uint {
	return s.PlatformType
}
