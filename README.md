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

创建表结构 (临时)
```sql
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `email` varchar(64) NOT NULL DEFAULT '',
  `password` varchar(64) NOT NULL DEFAULT '',
  `avater` varchar(255) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
```

+ https://github.com/silenceper/wechat  (go 微信包)