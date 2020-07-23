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
    
    
### feature
    1. 使用高性能的本地缓存bigcache
    2. oauth2第三方登录接入
    3. 多渠道第三方短信接入, 异常自动切换短信服务商
