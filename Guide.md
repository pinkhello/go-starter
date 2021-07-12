# 使用指南
  这是一个快速开发的脚手架, 它这边是最小化的. 这边有依赖的有 MySQL、Nsq
  
1. 下载代码
   ```shell
   git clone git@github.com:pinkhello/go-starter.git
   ```
2. Docker 启动 MySQL
   ```shell
   docker-compose -f deploy/mysql/docker-compose.yml up -d
   ```
   备注:
   > docker-compose 启动MySQL的时候 会通过 init.sql 初始化数据库和表
3. Docker 启动 NSQ
   ```shell
   docker-compose -f deploy/nsq/docker-compose.yml up -d
   ```
4. 分解项目结构
   ```
     - app 
       - cmd 
         - http.go         http 启动方法
         - root.go   
       - main.go           主入口
     - config              配置
     - docs                swag 自动生成的代码
     - internal
       - controller        controller层, http restful等映射层        
       - http
         - middleware.go   echo 中间件代码（包含了 日志、JWT验证、CROS等）
         - server.go       http server 启动代码
       - lib
         - xorm.go         xorm 适配
       - models            模型层 映射数据库层面
       - nsq               NSQ 生产者、消费者代码
       - repository
         - mysql            mysql repo层
         - mongo           mongo repo层（参照mysql层）                   
       - service
     - utils   error通用和工具通用代码
   ```
   所有的  `moudel.go` 都是 fx 依赖注入的代码, 为了结构清晰进行了单独编写.

5. 如何进行业务开发的
  + 编写 model 
    > internal/model/business_group.go
  + 编写 repo层
    > internal/repository/mysql/business_group.go
  + 编写 service层
    > internal/service/business_group.go
  + 编写 controller层
    > internal/controller/business_group_handler.go
  
6. 其他-参照 [README.md](README.md)
  
