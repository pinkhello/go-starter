# go-starter ![](https://img.shields.io/github/stars/pinkhello/go-starter?color=0088ff) ![](https://img.shields.io/github/forks/pinkhello/go-starter?color=0088ff) ![](https://img.shields.io/github/issues/pinkhello/go-starter?color=0088ff)

## github activity

[![Stargazers repo roster for @pinkhello/go-starter](https://reporoster.com/stars/pinkhello/go-starter)](https://github.com/pinkhello/go-starter/stargazers)

[![Stargazers over time](https://starchart.cc/pinkhello/go-starter.svg)](https://starchart.cc/pinkhello/go-starter)


## Demo
Demo: [房产CRM](http://121.4.242.26) 
- url: http://121.4.242.26 (test/123456)。
- BE [go-starter](https://github.com/pinkhello/go-starter)
- FE [ant-design-vue-pro](https://pro.antdv.com/)




## thirdparty
- [X] Github Actions
- [X] Custom Logger [logrus](https://github.com/sirupsen/logrus)
- [X] CLI Command [cobra](https://github.com/spf13/cobra)
- [X] Configuration [viper](https://github.com/spf13/viper)
- [X] Web [echo](https://github.com/labstack/echo)
- [X] DI/IOC [fx](https://github.com/uber-go/fx)
- [X] ORM [xorm](https://github.com/go-xorm/xorm)
- [X] Swagger generator [swag](https://github.com/swaggo/swag) [echo-swagger](https://github.com/swaggo/echo-swagger)
- [X] Messaging [NSQ](https://github.com/nsqio/nsq)
- [X] golangci-lint [golangci-lint](https://github.com/golangci/golangci-lint)
- [ ] Migrate [migrate](https://github.com/golang-migrate/migrate)
- [ ] ......

## code layer
```shell
   - app        # application main
     - cmd
     - ... 
   - config       # config
   - deploy       # ci/cd
     - mysql      # mysql docker-compose
     - nsq        # nsq docker-compose
     - ...        # other     
   - docs         # swag gen swagger2.0 doc
   - internal     # core 
     - controller # http handler（controller）
     - http       # http sever startup
     - lib        # lib
     - models     # models
     - nsq        # nsq producer & nsq consumer startup
     - repository # repository 
     - service    # service
   - utils        # util
     - ... 
   - ...
```

## Build & Publish & Deploy

### swag tips
```shell
> swag init -g app/main.go
```
`swagger url: http://{IP}:{PORT}/swagger/index.html`

### build
- local
    ```shell
    > cd .
    > docker build . --file deploy/Dockerfile --tag {ImageTag}
    ```
- github action 
  ```shell
  > { secrets.ACCESS_USERNAME }: `your docker hub username`
  ```

### be deploy
1. `docker network`:`go_starter_network`
    ```shell
    > docker network create go_starter_network
    ```
2. `mysql` & `nsq`
    ```shell
    # MYSQL start
    > cd deploy/mysql
    > docker-compose up -d
    # NSQ start
    > cd deploy/nsq
    > docker-compose up -d
    ```
3. `be server`
    ```shell
    # go-starter start
    > cd deploy
    > docker-compose up -d
    ```
4. `Health`
    ```shell
    http://{IP}:{PORT}/
    ```


### Other
- Dockerfile 
  ```dockerfile
    # build go 
    FROM golang:1.16.1-alpine3.13 as builder
    ......
    RUN CGO_ENABLED=0 GOOS=linux go build -o go_starter app/main.go
    
    # package stage
    FROM alpine
    ......
    # copy from builder
    COPY --from=builder /app/go_starter /app
    # ......
  ```
  
- Uber IOC/DI: [fx](https://pkg.go.dev/go.uber.org/fx)
  
    ```go
    //other code
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


