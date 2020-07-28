package base

import (
	"errors"

	"github.com/davveo/singleTsquare/utils/oauth2/github"
	"github.com/davveo/singleTsquare/utils/oauth2/qq"
	"github.com/davveo/singleTsquare/utils/oauth2/wechat"
	"github.com/davveo/singleTsquare/utils/oauth2/weibo"
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
}

func OauthService(serviceTag string) (ServiceInterface, error) {
	var service ServiceInterface
	switch serviceTag {
	case "qq":
		service = qq.NewService()
	case "github":
		service = github.NewService()
	case "weibo":
		service = weibo.NewService()
	case "wechat":
		service = wechat.NewService()
	default:
		return nil, errors.New("service does not exist")

	}
	return service, nil
}
