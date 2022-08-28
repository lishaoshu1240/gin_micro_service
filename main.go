/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-20 10:48:13
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-22 09:39:46
 * @FilePath: /finance_service/main.go
 * @Description: main.go
 */

package main

import (
	"finance_service/global"
	"finance_service/initialize"
	"finance_service/utils/nacos"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	//init viper
	global.GVA_VP = initialize.Viper()
	//init zap log
	global.GVA_LOG = initialize.Zap()
	//init redis
	global.GVA_REDIS = initialize.Redis()

	//init nacos client
	global.GVA_NACOS = nacos.NacosInit()
	//register servie to nacos
	if !nacos.RegisterService(*global.GVA_NACOS) {
		global.GVA_LOG.Error("Faild, register service to nacos fail.")
	} else {
		global.GVA_LOG.Info("register nacos success.", zap.String("service name:", global.GVA_CONFIG.Nacos.RegisterServicename))
	}

	//start mysql
	global.GVA_DB = initialize.OpenMysql("finance_center")
	if global.GVA_DB != nil {
		global.GVA_LOG.Debug("init finance_center db done.")
		// close db before server close.
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

	//start gin
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	initialize.InitServer(address, Router)

}
