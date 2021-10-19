## init

```
cd $GOPATH

cd src/github.com/tinywan/

mkdire gin.tinywan.cn

cd gin.tinywan.cn

go mod init gin.tinywan.cn
```

main.go
```golang
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
```

add required module
```
go get github.com/gin-gonic/gin
```
运行代码
```
go run main.go
```
访问地址
```
http://127.0.0.1:8080/ping
```

使用 go 工具构建和安装该程序

```
go install gin.tinywan.cn
```
[更多](https://golang.org/doc/code)