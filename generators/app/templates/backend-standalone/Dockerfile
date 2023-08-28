# backend build
FROM golang:1.20-alpine AS backend
LABEL maintainer="hakurei <hfcj1839@gmail.com>"
WORKDIR /usr/src/app
# ENV GOPROXY=https://goproxy.cn
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN CARGO_NET_GIT_FETCH_WITH_CLI=1 && \
    apk --no-cache add \
        upx \
        ca-certificates\
        tzdata \
        sudo \
        python3\
        py3-pip \
        openssl \
        ca-certificates \
        sshpass \
        openssh-client \
        rsync \
        git && \
    apk --no-cache add --virtual build-dependencies \
        python3-dev \
        libffi-dev \
        musl-dev \
        gcc \
        cargo \
        openssl-dev \
        # libressl-dev \
        build-base

# RUN pip install -U pip wheel && \
#     pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple && \
#     pip install --upgrade cryptography cffi && \
#     pip install --ignore-installed jmespath && \
#     pip install --upgrade pywinrm
RUN apk del build-dependencies && \
    rm -rf /var/cache/apk/* && \
    rm -rf /root/.cache/pip && \
    rm -rf /root/.cargo

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" >  /etc/timezone

RUN apk update && apk upgrade && apk add --no-cache bash bash-doc bash-completion \
    && rm -rf /var/cache/apk/*

COPY ./go.mod ./
COPY ./go.sum ./

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release

RUN go mod download
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag
RUN swag init
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o server &&\
    upx --best server -o _upx_server && \
    mv -f _upx_server server
