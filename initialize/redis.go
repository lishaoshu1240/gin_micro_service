/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-20 16:41:10
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-20 17:47:29
 * @FilePath: /gin_micro_service/initialize/redis.go
 * @Description: redis
 */

package initialize

import (
	"context"

	"gin_micro_service/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() *redis.Client {
	redisCfg := global.GVA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GVA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.GVA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
	}
	return client
}
