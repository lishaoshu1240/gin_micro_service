/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-20 18:58:49
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-21 12:25:37
 * @FilePath: /gin_micro_service/initialize/router.go
 * @Description: 总路由
 */

package initialize

import (
	"gin_micro_service/global"
	"gin_micro_service/router"

	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()

	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开

	//get route handle
	apiRouter := router.RouterGroup.ApiRouter

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		apiRouter.InitApiRouter(PublicGroup)

	}

	global.GVA_LOG.Info("router register success")
	return Router
}
