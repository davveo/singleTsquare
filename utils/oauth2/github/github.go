package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/davveo/singleTsquare/utils/oauth2/base"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// 通过code获取token认证url
func (s *Service) GetTokenAuthUrl(code string) string {
	return fmt.Sprintf(
		"%s?client_id=%s&client_secret=%s&code=%s",
		TokenURL, conf.ClientId, conf.ClientSecret, code,
	)
}

// 获取 token
func (s *Service) GetToken(code string) (*Token, error) {
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
	var token Token
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *Service) GenRedirectURL() string {
	// 跳转到第三方地址
	params := url.Values{}
	params.Add("client_id", AppId)
	params.Add("state", "state")
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), RedirectURI)
	loginURL := fmt.Sprintf("%s?%s", AuthorizeURL, str)
	return loginURL
}

func (s *Service) GetUserInfo(code string) (*base.UserInfo, error) {
	var unmarshalUserInfo *UnmarshalUserInfo
	var req *http.Request
	var err error
	// 获取token
	token, err := s.GetToken(code)
	if err != nil {
		return nil, err
	}

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

	defer res.Body.Close()
	bs, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(bs, &unmarshalUserInfo)
	if err != nil {
		return nil, err
	}

	return &base.UserInfo{
		AccessToken: token.AccessToken,
		OpenId:      unmarshalUserInfo.ID,
		Avatar:      unmarshalUserInfo.Avatar,
		NickName:    unmarshalUserInfo.NickName,
	}, nil
}
