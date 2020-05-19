# logger

logrus for gin

## install

```shell
go get github.com/ganodermaking/logger
```

## used

```go
package main

import (
 "github.com/ganodermaking/logger"
 "github.com/gin-gonic/gin"
 "github.com/sirupsen/logrus"
)

func main() {
 logger.NewLogger("./logs/debug.log")  // 错误日志

 r := gin.Default()
 r.Use(logger.Middleware("./logs/debug.access.log")) // 访问日志
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
