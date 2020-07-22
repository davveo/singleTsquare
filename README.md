### singleTsquare
  单体版本
### usage
> 执行数据库迁移
> make migrate

> 运行服务
> make run

> 清理服务
> make clean

### 接口规范
返回参数
    code number  结果状态码  成功=0 失败=-1 未登录=401 无权限=403
    showMsg string 显示结果 系统繁忙, 稍后重试
    errorMsg string 错误信息 便于研发定位
    data object 数据 json格式
    
分页
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
    
   