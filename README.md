# push-go

### push是什么？
类似于ntfy，ntfy有丰富的功能和详细的文档，但是受限于本地部署后ios设备无法实时接收到推送的消息，所以有了push项目，旨在开发更接近ntfy的实用功能和更适用于国内的网络环境

[目前比较好用的推送服务](https://byzhao.cn/2024/07/15/clyqro3d40000fcjbhimh3t0f/)


> 支持`sqlite` 和 `mysql` 默认使用`sqlite`

> [此项目对应ios客户端app项目](https://github.com/Endless-zby/push-swift)

> [此项目的java版服务端](https://github.com/Endless-zby/push)


### 完整的config.yaml
```yaml
server:
  port: 10002
database:
  type: "sqlite" # 选择持久化数据库"sqlite" 或 "mysql"
  mysql:
    host: "192.168.192.36"
    port: 3306
    user: "byzhao"
    password: "zby123456"
    dbname: "byzhao"
  sqlite:
    file: "push.db"
apns: # Apns的证书使用的bark项目作者公开证书
  keyId: "LH4T9V5U4R"
  teamId: "5U8LBRXG3A"
  authKeyFile: "AuthKey_LH4T9V5U4R_5U8LBRXG3A.p8"
```

### 项目下构建 Docker 镜像

```shell
docker build -t push-server:latest .
```

### 运行 Docker 容器
```shell
docker run -v $(pwd)/configs:/app/configs -p 10002:10002 push-server
```


