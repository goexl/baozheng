# baozheng

[![Sourcegraph](https://sourcegraph.com/github.com/goexl/baozheng/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/goexl/baozheng?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/goexl/baozheng)
[![Go Report Card](https://goreportcard.com/badge/github.com/goexl/baozheng?style=flat-square)](https://goreportcard.com/report/github.com/goexl/baozheng)
[![Build Status](https://github.ruijc.com:20443/api/badges/goexl/baozheng/status.svg)](https://github.ruijc.com:20443/goexl/baozheng)
[![Codecov](https://img.shields.io/codecov/c/github/goexl/baozheng.svg?style=flat-square)](https://codecov.io/gh/goexl/baozheng)
[![License](https://img.shields.io/github/license/goexl/baozheng)](https://raw.githubusercontent.com/goexl/baozheng/master/LICENSE)

Golang数据验证方法，包括如下验证器

- 密码
- 特殊符号

## 快速开始

`Baozheng`使用非常简单

```go
package main

import (
  `github.com/goexl/baozheng`
)

func main() {
  _ = baozheng.Password(`te88@@`)
}
```

> 推荐使用`Xiren`集成更多的验证器，请参看[**使用文档**](https://xiren.storezhang.tech)

## 捐助

![支持宝](doc/docs/.vuepress/public/donate/alipay-small.jpg)
![微信](doc/docs/.vuepress/public/donate/weipay-small.jpg)
