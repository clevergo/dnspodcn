# DNSPod.cn Go SDK DNSPod.cn Go 开发工具包
[![Build Status](https://img.shields.io/travis/clevergo/dnspodcn?style=for-the-badge)](https://travis-ci.org/clevergo/dnspodcn)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/dnspodcn?style=for-the-badge)](https://coveralls.io/github/clevergo/dnspodcn)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/clevergo.tech/dnspodcn?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/dnspodcn?style=for-the-badge)](https://goreportcard.com/report/github.com/clevergo/dnspodcn)
[![Release](https://img.shields.io/github/release/clevergo/dnspodcn.svg?style=for-the-badge)](https://github.com/clevergo/dnspodcn/releases)
[![Downloads](https://img.shields.io/endpoint?url=https://pkg.clevergo.tech/api/badges/downloads/total/clevergo.tech/dnspodcn?style=flat-square)](https://pkg.clevergo.tech/clevergo.tech/dnspodcn)
[![Chat](https://img.shields.io/badge/chat-telegram-blue?style=flat-square)](https://t.me/clevergotech)
[![Community](https://img.shields.io/badge/community-forum-blue?style=flat-square&color=orange)](https://forum.clevergo.tech)

`dnspodcn` 是 [DNSPod](https://www.dnspod.cn/docs/index.html) API 的 Go SDK。其中 `Provider` 实现了 [libdns](https://github.com/libdns/libdns/) 接口。

## 安装

```shell
$ go get -u clevergo.tech/dnspodcn
```

## 使用

```go
import "clevergo.tech/dnspodcn"

client := dnspodcn.NewClient("APP_ID", "APP_TOKEN")
```

> `APP_ID` 和 `APP_TOKEN` 为 DNSPod 后台生成的 API 的 ID 和 Token。

## 接口

- [ ] DNS 记录
    - [x] 新增记录：`client.CreateRecord`
    - [x] 删除记录：`client.DeleteRecord`
    - [x] 查询记录：`client.QueryRecords`
    - [x] 修改记录：`client.UpdateRecord`
- [ ] ...

目前仅实现了 DNS 记录的相关接口，如果你有其他接口要求，可以通过 [issue](https://github.com/clevergo/dnspodcn/issues/new) 反馈，也可以自行实现并通过 PR 提交。
