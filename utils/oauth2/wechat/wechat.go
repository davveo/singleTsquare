package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/davveo/singleTsquare/utils/oauth2/base"
)

// https://blog.csdn.net/qq_35781732/article/details/82662021
// https://blog.csdn.net/weixin_43851310/article/details/105816815
// https://github.com/yun-mu/weixin-login/blob/master/src/controller/weixin.go
// https://github.com/yizenghui/WechatQrcodeServe
// TODO 微信授权登录
// http://127.0.0.1:8080/api/v1/oauth/wechat/callback

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GenRedirectURL() string {
	params := url.Values{}
	params.Add("appid", AppId)
	params.Add("state", "login")
	params.Add("response_type", "code")
	params.Add("scope", "snsapi_userinfo")
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), RedirectURL)
	loginURL := fmt.Sprintf("%s?%s", AuthorizeURL, str)
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
func (s *Service) GetAccessToken(code string) (*WXBody, error) {
	params := url.Values{}
	params.Add("code", code)
	params.Add("appid", AppId)
	params.Add("secret", AppSecret)
	params.Add("grant_type", "authorization_code")
	resp, err := http.Get(fmt.Sprintf("%s?%s", AccessTokenURL, params.Encode()))
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

func (s *Service) GetUserInfo(code string) (*base.UserInfo, error) {
	token, err := s.GetAccessToken(code)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("lang", "zh_CN")
	params.Add("openid", token.Openid)
	params.Add("access_token", token.AccessToken)
	infoBody, err := http.Get(fmt.Sprintf("%s?%s", UserInfoURL, params.Encode()))

	defer infoBody.Body.Close()
	info, _ := ioutil.ReadAll(infoBody.Body)

	updateUser := WXInfo{}
	err = json.Unmarshal(info, &updateUser)
	if err != nil {
		return nil, err
	}
	return &base.UserInfo{
		OpenId:      token.Openid,
		AccessToken: token.AccessToken,
		Avatar:      updateUser.HeadimgUrl.(string),
		NickName:    updateUser.Nickname.(string),
	}, nil
}
