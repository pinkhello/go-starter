# go-starter
## Go微服务框架整合使用的第三方依赖库与技术
- [X] Github Actions
- [X] 自定义Logger [logrus](https://github.com/sirupsen/logrus)
- [X] CLI命令 [cobra](https://github.com/spf13/cobra)
- [X] 配置读取器 [viper](https://github.com/spf13/viper)
- [X] Web框架 [echo](https://github.com/labstack/echo)
- [X] 依赖注入 [fx](https://github.com/uber-go/fx)
- [X] ORM [xorm](https://github.com/go-xorm/xorm)
- [X] Swagger generator [swag](https://github.com/swaggo/swag) [echo-swagger](https://github.com/swaggo/echo-swagger)
- [X] Messaging [NSQ](https://github.com/nsqio/nsq)
- [X] golangci-lint [golangci-lint](https://github.com/golangci/golangci-lint)
- [ ] Migrate [migrate](https://github.com/golang-migrate/migrate)
- [ ] ......

## 代码结构分层
```shell
   - app        # 主程序入口
     - cmd
     - ... 
   - config       # 配置相关
   - deploy       # ci/cd相关, 镜像相关
     - mysql      # mysql docker-compose
     - nsq        # nsq docker-compose
     - ...        # 其他部署相关     
   - docs         # swag 生成的 swagger2.0 文档目录
   - internal     # 核心业务逻辑
     - controller # http handler（controller层）
     - http       # http sever 启动入口
     - lib        # 核心基础
     - models     # 模型
     - nsq        # nsq producer 和 nsq consumer 启动入口与业务
     - repository # repository 层
     - service    # service层
   - utils        # 工具通用(工具方法、常量、通用错误定义)
     - ... 
   - ...
```

# swag 生成最新的Swagger文档
```shell
> swag init -g app/main.go
```
产生的swagger地址为 http://{IP}:{PORT}/swagger/index.html

# 构建和发布镜像

## 本地
```shell
> cd .
> docker build . --file deploy/Dockerfile --tag {ImageTag}
```
## github action
在CI内部已经做了自动化构建和发布镜像（ `github` 当前项目下`secrets`配置的 { secrets.ACCESS_USERNAME } 的值是 `docker hub 用户名`）

## BE部署
1. 创建 `docker` 自定义网络`go_starter_network`
```shell
> docker network create go_starter_network
```
2. `mysql` 启动
```shell
> cd deploy/mysql
> docker-compose up -d
```
3. `nsq` 启动
```shell
> cd deploy/nsq
> docker-compose up -d
```
4. BE后端服务启动
```shell
> cd deploy
> docker-compose up -d
```
