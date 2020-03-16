# logger
logrus,lfshook,rotatelogs for logs

# install
```shell
go get github.com/ganodermaking/logger
```

# used
```go
package main

import (
	"github.com/ganodermaking/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// 初始化
	logger.NewLogger("./logs/debug.log")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		message := "hello world"
		logrus.Info(message)
		c.JSON(200, gin.H{
			"message": message,
		})
	})
	r.Run(":8081")
}
```