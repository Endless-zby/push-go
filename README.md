# push-go


### 完整的config.yaml
```yaml

database:
  type: "mysql" # "sqlite" 或 "mysql"
  mysql:
    host: "192.168.192.36"
    port: 3306
    user: "byzhao"
    password: "zby123456"
    dbname: "byzhao"
  sqlite:
    file: "../push.db"
```

### 项目下构建 Docker 镜像

```shell
docker build -t push-server:latest .
```

### 运行 Docker 容器
```shell
docker run -v $(pwd)/configs:/app/configs -p 10002:10002 push-server
```


