# go-starter

## GoPush ![](https://img.shields.io/github/stars/pinkhello/go-starter?color=0088ff) ![](https://img.shields.io/github/forks/pinkhello/go-starter?color=0088ff) ![](https://img.shields.io/github/issues/pinkhello/go-starter?color=0088ff)

--------------
## 集成的第三方库与技
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
--------------
## 项目代码结构分层
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
--------------
## 构建发布依赖

### swag 生成最新的Swagger文档
```shell
> swag init -g app/main.go
```
产生的swagger地址为 http://{IP}:{PORT}/swagger/index.html

### 构建和发布镜像
- 本地
    ```shell
    > cd .
    > docker build . --file deploy/Dockerfile --tag {ImageTag}
    ```
- github action 
  > 在CI内部已经做了自动化构建和发布镜像（ `github` 当前项目下`secrets`配置的 { secrets.ACCESS_USERNAME } 的值是 `docker hub 用户名`）

### BE初始化和启动部署
1. 创建`docker`自定义网络`go_starter_network`
    ```shell
    > docker network create go_starter_network
    ```
2. `mysql`启动 与 `nsq`启动
    ```shell
    # MYSQL 启动
    > cd deploy/mysql
    > docker-compose up -d
    # NSQ 启动
    > cd deploy/nsq
    > docker-compose up -d
    ```
3. BE后端服务启动
    ```shell
    # go-starter 启动
    > cd deploy
    > docker-compose up -d
    ```
4. 访问健康坚持接口
    ```shell
    http://{IP}:{PORT}/
    ```
--------------

### 其他文件相关描述

- Dockerfile -- Docker镜像分为两阶段构建, Builder阶段 与 打包阶段。
  ```dockerfile
    # Builder阶段构建二进制可执行文件
    FROM golang:1.16.1-alpine3.13 as builder
    ......
    RUN CGO_ENABLED=0 GOOS=linux go build -o go_starter app/main.go
    
    # 打包阶段
    FROM alpine
    ......
    # 从构建阶段COPY生成的可执行文件
    COPY --from=builder /app/go_starter /app
    # ......
  ```
  
- Uber 依赖框架 fx.   [fx 开发包详细说明](https://pkg.go.dev/go.uber.org/fx)
  
    ```go
    //上面的代码不写了
    ......
    
    var (
        httpCmd = &cobra.Command{
            Use:   "http",
            Short: "Start Http REST API",
            Run:   initHTTP,
        }
    )
    
    func initHTTP(cmd *cobra.Command, args []string) {
        fx.New(inject()).Run()
    }
    
    // 注意前后依赖关系，顺序启动。
    func inject() fx.Option {
        return fx.Options(
            fx.Provide(
                config.NewConfig,
                utils.NewTimeoutContext,
            ),
            libs.XormModule,
            repository.Module,
            service.Module,
            controller.Module,
            nsq.ProducerModule,
            nsq.ConsumerModule,
            http.Module,
        )
    }
    
    ```
  
- swag 说明 
  `doc 目录下为自动生成的，不需要更改，需要更改的化话需要更高主要上面的注释`


# Star 趋势

[![Stargazers over time](https://starchart.cc/PinkHello/go-starter.svg)](https://starchart.cc/PinkHello/go-starter)
