package oauth2

import (
	"errors"
	"github.com/davveo/singleTsquare/utils/oauth2/facebook"
	"github.com/davveo/singleTsquare/utils/oauth2/github"
	"github.com/davveo/singleTsquare/utils/oauth2/qq"
	"github.com/davveo/singleTsquare/utils/oauth2/wechat"
	"github.com/davveo/singleTsquare/utils/oauth2/weibo"
)

type FuncRedirectURL func() string

func ServiceRedirectURL(serviceTag string) (string, error) {
	// TODO 可能存在其他优雅的注册方式
	// TODO 后续在改造
	if service, ok := map[string]FuncRedirectURL{
		"qq":       qq.GenRedirectURL,
		"weibo":    weibo.GenRedirectURL,
		"wechat":   wechat.GenRedirectURL,
		"github":   github.GenRedirectURL,
		"facebook": facebook.GenRedirectURL,
	}[serviceTag]; ok {
		return service(), nil
	} else {
		return "", errors.New("service does not exist")
	}
}
