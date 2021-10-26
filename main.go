package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "pong",
		})
	})
	err := r.Run("0.0.0.0:8081")
	if err != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
