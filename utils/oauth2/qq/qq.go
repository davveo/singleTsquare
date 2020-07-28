package qq

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/davveo/singleTsquare/utils/oauth2/base"
)

type Service struct {
	PlatformType uint
}

func NewService() *Service {
	return &Service{PlatformType: 1}
}

func (s *Service) RequestAccessToken(code string) (*PrivateInfo, error) {
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

	resultMap := s.convertToMap(body)

	info := &PrivateInfo{}
	info.AccessToken = resultMap["access_token"]
	info.RefreshToken = resultMap["refresh_token"]
	info.ExpiresIn = resultMap["expires_in"]

	return info, nil
}

func (s *Service) GetOpenId(accessToken string) (string, error) {
	resp, err := http.Get(fmt.Sprintf(
		"%s?access_token=%s", OpenidUrl, accessToken))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	return string(bs)[45:77], nil
}

func (s *Service) GetUserInfo(code string) (*base.UserInfo, error) {
	var unmarshaluserInfo *UnmarshalUserInfo
	tokenInfo, err := s.RequestAccessToken(code)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	accessToken := tokenInfo.AccessToken
	openid, _ := s.GetOpenId(accessToken)
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

	err = json.Unmarshal(bs, &unmarshaluserInfo)
	if err != nil {
		return nil, err
	}
	// 将用户的标示写入

	return &base.UserInfo{
		OpenId:      openid,
		AccessToken: accessToken,
		NickName:    unmarshaluserInfo.NickName,
		Avatar:      unmarshaluserInfo.Avatar,
	}, nil
}

func (s *Service) convertToMap(str string) map[string]string {
	var resultMap = make(map[string]string)
	values := strings.Split(str, "&")
	for _, value := range values {
		vs := strings.Split(value, "=")
		resultMap[vs[0]] = vs[1]
	}
	return resultMap
}

func (s *Service) GenRedirectURL() string {
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", AppId)
	params.Add("state", "test")
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), RedirectURL)
	loginURL := fmt.Sprintf("%s?%s", AuthorizeURL, str)
	return loginURL
}

func (s *Service) GetPlatformType() uint {
	return s.PlatformType
}
