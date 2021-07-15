# wx服务

### 初始化项目
- 容器及代理：dockerfile docker-compose.yml
- 创建入口文件，开启http（gin），设计路由
- docker-compose up 启动项目
- 添加项目githuh持续集成 github-workflows
- 创建 handler - service - repository - handler(controller)/test
- 引入gorm
- 引入jwt
- 编写injection依赖注入
- 注册登陆模块
- ***退出登陆 （未完成）***
- 校验中文
- 权限模块 进行中 （通用分页优化；role，permisiion）
- ---- 日志统一管理？？？？
- ----casbin权限？？？？




```code
-------

 复制新模块：（替换时注意大小写，替换大写，替换小写）
 - 创建model
 - 复制iface-role（service，repository） 改对应的名字
 - 复制dto-role 改对应的名字
- 复制repository-role 改对应的名字
- 复制service-role 改对应的名字
- 复制handle-role 改对应的名字
- 其他：更改request对应字段，更改update能更新的字段，更新inject
----完事大吉

-------
```




创建表结构 (临时)
```sql

```

+ https://github.com/silenceper/wechat  (go 微信包)