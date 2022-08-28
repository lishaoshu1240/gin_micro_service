/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-20 11:57:24
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-21 12:43:57
 * @FilePath: /gin_micro_service/global/global.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package global

import (
	"gin_micro_service/config"

	"github.com/go-redis/redis/v8"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB          *gorm.DB
	GVA_DB_ORDER    *gorm.DB
	GVA_DB_CUSTOMER *gorm.DB
	GVA_REDIS       *redis.Client
	GVA_CONFIG      config.Server
	GVA_VP          *viper.Viper

	GVA_LOG   *zap.Logger
	GVA_NACOS *naming_client.INamingClient
)
