package qq

import (
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
	MeURL        = "https://graph.qq.com/oauth2.0/me"
	UserInfoURL  = "https://graph.qq.com/user/get_user_info"
)

type PrivateInfo struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
}

func TokenParams(code string) string {
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", AppId)
	params.Add("client_secret", AppKey)
	params.Add("code", code)
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), RedirectURL)
	return fmt.Sprintf("%s?%s", TokenURL, str)
}

func GetToken(code string) (*PrivateInfo, error) {
	loginUrl := TokenParams(code)
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

// 根据access_token获取openid
func GetOpenId(info *PrivateInfo) (*PrivateInfo, error) {
	resp, err := http.Get(fmt.Sprintf(
		"%s?access_token=%s", MeURL, info.AccessToken))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	body := string(bs)
	info.OpenId = body[45:77]
	return info, nil
}

func GetUserInfo(info *PrivateInfo) (string, error) {
	params := url.Values{}
	params.Add("access_token", info.AccessToken)
	params.Add("openid", info.OpenId)
	params.Add("oauth_consumer_key", AppId)

	uri := fmt.Sprintf("%s?%s", UserInfoURL, params.Encode())
	resp, err := http.Get(uri)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	return string(bs), nil
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
