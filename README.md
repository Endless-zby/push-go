# push-go

### 项目下构建 Docker 镜像
```shell
docker build -t push-server:latest .
```

### 运行 Docker 容器
```shell
docker run -v $(pwd)/configs:/app/configs -p 10002:10002 push-server
```