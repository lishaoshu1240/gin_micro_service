/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-20 11:44:23
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-20 11:44:52
 * @FilePath: /gin_micro_service/config/system.go
 * @Description:
 */
package config

type System struct {
	Env    string `mapstructure:"env" json:"env" yaml:"env"`            // 环境值
	Addr   int    `mapstructure:"addr" json:"addr" yaml:"addr"`         // 端口值
	DbType string `mapstructure:"db-type" json:"dbType" yaml:"db-type"` // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
}
