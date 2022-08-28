/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-20 12:00:36
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-20 12:02:47
 * @FilePath: /finance_service/config/enter.go
 * @Description: enter.go
 */
package config

type Server struct {
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// oss
	Nacos Nacos `mapstructure:"nacos" json:"nacos" yaml:"nacos"`
}
