package oauth2

import (
	"errors"
)

type UserInfo struct {
	NickName    string
	OpenId      string
	Avatar      string
	AccessToken string
}

type ServiceInterface interface {
	GenRedirectURL() string
	GetUserInfo(code string) (*UserInfo, error)
	GetPlatformType() uint
}

func OauthService(serviceTag string) (ServiceInterface, error) {
	var service ServiceInterface
	switch serviceTag {
	case "qq":
		service = NewQQService()
	case "github":
		service = NewGithubService()
	case "weibo":
		service = NewWeiboService()
	case "wechat":
		service = NewWechatService()
	default:
		return nil, errors.New("service does not exist")

	}
	return service, nil
}
