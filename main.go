package main

import (
   "fmt"
   "github.com/fatih/color"
   "go.uber.org/zap"
   "sample-go-service/global"
   "sample-go-service/initialize"   
   "sample-go-service/utils"
)

func main() {
   //1.初始化yaml配置
      initialize.InitConfig()

      r := gin.Default()
      r.GET("/ping", func(c *gin.Context) {
   	c.JSON(200, gin.H{
   		"message": "pong",
   	})
   })
   r.Run(fmt.Sprintf(":%d", global.Settings.Port))
}