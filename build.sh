#!/bin/bash

# 定义目标平台  linux/armv7 为玩客云设备编译
platforms=("windows/amd64" "windows/386" "linux/amd64" "linux/386" "linux/armv7" "darwin/amd64" "darwin/arm64")

for platform in "${platforms[@]}"
do
    # 拆分平台变量
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    # 输出文件名
    output_name='push-server-'$GOOS'-'$GOARCH

    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    # 设置环境变量并编译
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name

    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done