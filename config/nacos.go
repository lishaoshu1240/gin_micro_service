/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-20 11:48:12
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-20 11:48:24
 * @FilePath: /gin_micro_service/config/nacos.go
 * @Description: nacos
 */

package config

type Nacos struct {
	RegisterIp          string `mapstructure:"register-ip" json:"registerIp" yaml:"register-ip"`
	RegisterPort        int64  `mapstructure:"register-port" json:"registerPort" yaml:"register-port"`
	RegisterServicename string `mapstructure:"register-servicename" json:"registerServicename" yaml:"register-servicename"`
	NacosIp             string `mapstructure:"nacos-ip" json:"nacosIp" yaml:"nacos-ip"`
	NacosPort           int64  `mapstructure:"nacos-port" json:"nacosPort" yaml:"nacos-port"`
	NacosUsername       string `mapstructure:"nacos-username" json:"nacosUsername" yaml:"nacos-username"`
	NacosPassword       string `mapstructure:"nacos-password" json:"nacosPassword" yaml:"nacos-password"`
	NacosLogdir         string `mapstructure:"nacos-logdir" json:"nacosLogdir" yaml:"nacos-logdir"`
}
