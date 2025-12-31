# YaegiDemo

演示如何使用 [Yaegi](https://github.com/traefik/yaegi) 在运行时加载并执行 Go 代码。

## 项目简介

HTTP 处理逻辑位于 `rule/rule.go`，由 Yaegi 热加载。

## 前置条件

- Go 1.22+

## 快速开始

启动示例服务：

```bash
make run
```

发送请求：

```bash
curl localhost:8080 -i
```

示例响应（节选）：

```bash
HTTP/1.1 200 OK
Cache-Control: no-cache
Content-Type: text/plain; charset=utf-8
Via: useless-helper
Date: Thu, 23 Nov 2023 12:24:31 GMT
Content-Length: 241

Hello!
Your Content-Type is
Your User-Agent is curl/8.1.2
Your Host is localhost:8080
Your RemoteAddr is 127.0.0.1:61784
Your RequestURI is /
Your Method is GET
Your Proto is HTTP/1.1
Your URL is /
Code above is wrote by GitHub Copilot.

:)
```

## 热加载

编辑 `rule/rule.go` 并保存，服务端响应会自动更新，无需重启。

## 配置

- `--rule`（默认：`./rule/rule.go`）：规则文件路径

## 开发说明

- `make generate`：重新生成 Yaegi 符号表
- `make build`：构建二进制到 `output/yaegi_demo`

## 项目结构

```text
.
├── cmd
│   └── main.go         # 程序入口
├── internal
│   ├── engine          # Yaegi 执行与热加载
│   ├── helper          # 示例辅助函数
│   └── symbol          # 解释器符号注册
├── rule
│   └── rule.go          # 热加载的处理逻辑
├── Makefile             # 开发命令
├── README.md            # 项目说明
├── README.zh_CN.md      # 中文说明
├── go.mod               # 模块定义
└── go.sum               # 依赖校验
```
