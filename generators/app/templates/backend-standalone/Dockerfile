# backend build
FROM golang:1.22-alpine AS backend
LABEL maintainer="hakurei <hfcj1839@gmail.com>"

WORKDIR /usr/src/app

# 设置时区和代理
ENV GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release \
    GO111MODULE=on \
    TZ=Asia/Shanghai

# 替换源以提高速度
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update && \
    apk upgrade && \
    apk add --no-cache \
        upx \
        ca-certificates \
        tzdata \
        sudo \
        python3 \
        py3-pip \
        openssl \
        sshpass \
        openssh-client \
        rsync \
        git \
        bash bash-doc bash-completion && \
    apk add --no-cache --virtual build-dependencies \
        python3-dev \
        libffi-dev \
        musl-dev \
        gcc \
        cargo \
        openssl-dev \
        build-base && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# 安装Python包
RUN pip install --upgrade pip wheel && \
    pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple && \
    pip install --upgrade cryptography cffi pywinrm

# 清理安装后的残留文件
RUN apk del build-dependencies && \
    rm -rf /var/cache/apk/* && \
    rm -rf /root/.cache/pip && \
    rm -rf /root/.cargo

# 复制 Go 模块相关文件
COPY go.mod go.sum ./

# 下载 Go 依赖
RUN go mod download

# 复制项目源码
COPY . .

# 构建 Go 应用
RUN go install github.com/swaggo/swag/cmd/swag && \
    swag init && \
    CGO_ENABLED=0 go build -ldflags "-s -w" -o server && \
    upx --best server -o _upx_server && \
    mv -f _upx_server server
