package weibo

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

func (s *Service) TokenParams(code string) string {
	params := url.Values{}
	params.Add("code", code)
	params.Add("client_id", AppId)
	params.Add("client_secret", AppKey)
	params.Add("grant_type", "authorization_code")
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), RedirectURL)
	return fmt.Sprintf("%s?%s", TokenURL, str)
}

func (s *Service) GetToken(code string) (*AccessTokenInfo, error) {
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

func (s *Service) GetUserInfo(code string) (*base.UserInfo, error) {
	var unmarshalUserInfo *UnmarshalUserInfo
	accessTokenInfo, err := s.GetToken(code)
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
	err = json.Unmarshal(bs, &unmarshalUserInfo)
	if err != nil {
		return nil, err
	}
	return &base.UserInfo{
		OpenId:      accessTokenInfo.Uid,
		Avatar:      unmarshalUserInfo.Avatar,
		NickName:    unmarshalUserInfo.NickName,
		AccessToken: accessTokenInfo.AccessToken,
	}, nil
}

func (s *Service) GenRedirectURL() string {
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", AppId)
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), RedirectURL)
	loginURL := fmt.Sprintf("%s?%s", AuthorizeURL, str)
	return loginURL
}
