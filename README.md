### singleTsquare
单体版本

### 本地开发
> step one: 执行数据库迁移
> make migrate

> step two: 运行服务
> make run

> 清理服务
> make clean

> 本地快捷提交
> make push msg="commit msg"

### 接口规范
返回参数
> code number  结果状态码  成功=0 失败=-1 未登录=401 无权限=403
> showMsg string 显示结果 系统繁忙, 稍后重试
> errorMsg string 错误信息 便于研发定位
> data object 数据 json格式
 
不带分页
```json
{
        "code": 0,
        "showMsg": "success",
        "errorMsg": "",
        "data": {
            
        }
}
```  
带分页
```json
{
        "code": 0,
        "showMsg": "success",
        "errorMsg": "",
        "data": {
            "list": [],
            "pagination": {
                "total": 100,
                "currentPage": 1,
                "prePageCount": 10
            }
        }
}
```
### TODO
    1. 统一状态码以及错误信息返回
    2. 第三方短信服务商接入
    3. fix TODO以及一些代码优化
    4. 用户在使用第三方登录授权后，需要进行绑定
    
### feature
    1. 使用高性能的本地缓存bigcache
    2. 多种登录方式
       - 用户名/邮箱 + 密码
       - 手机号+验证码
       - oauth2(qq/wechat/weibo/github)快速登录
       - 扫码登录
       - 扫脸登录
    3. 多渠道第三方短信接入, 异常自动切换短信服务商
    
### Oauth2
    1. 用户点击第三方登录图标，发送请求到后端
    2. 后端拼装请求参数, 重定向到第三方进行授权
    3. 第三方授权通过后，会根据参数的redirect_uri重定向到指定uri,并且带上code
      > 如: http://xxx.xxx.com?code=xxxxxxx
    5. 后端捕获这个请求, 获取code, 拼装参数重新请求第三方获取token
    6. 然后根据token获取用户信息，完成登录流程。
