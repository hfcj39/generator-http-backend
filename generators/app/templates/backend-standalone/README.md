# <%= displayName %>

## 开发前的准备

1. 初次运行项目需生成 swag 文档

   `$ swag init` 或者 `$ make doc`

## 开发项目

推荐 air 启动 `air -c .air.toml`

## Unit test

`$ make test`

单元测试文件在 tests 目录内，单元测试文件以`_test`结尾

手动执行单测

`$ OPS_CONFIG=../config.test.yaml go test ./... -v`

You can also generate coverage profile using Cover tool

`$ go test ./... -v -coverpkg=./... -coverprofile=coverage.out`

查看去除 docs 的覆盖率

`$ cat coverage.out | grep -v docs.go > cc.out | go tool cover -func=cc.out`

查看全部覆盖率

`$ go tool cover -func=coverage.out`

To analyze coverage via a browser, you can also use

`$ go tool cover -html=coverage.out`

Here are various options supported by Cover command

Usage of 'go tool cover':
Given a coverage profile produced by ‘go test’:

`go test -coverprofile=c.out`

Open a web browser displaying annotated source code:

`go tool cover -html=c.out`

Write out an HTML file instead of launching a web browser:

`go tool cover -html=c.out -o coverage.html`

Display coverage percentages to stdout for each function:

`go tool cover -func=c.out`
