/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-20 11:45:17
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-20 11:45:26
 * @FilePath: /finance_service/config/redis.go
 * @Description: redis配置读取
 */
package config

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
}
