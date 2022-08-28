/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-20 19:04:49
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-21 10:20:49
 * @FilePath: /finance_service/initialize/gin_server.go
 * @Description: gin server init
 */

package initialize

import (
	"finance_service/global"
	"fmt"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server interface {
	ListenAndServe() error
}

func InitServer(address string, router *gin.Engine) {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20

	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	global.GVA_LOG.Info(fmt.Sprintf(`
	Welecome to use: finance_service
	Current version: V1.0.0 Beta
	Listen port:%v
`, global.GVA_CONFIG.System.Addr))
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
