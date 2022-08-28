/*
 * @Author: stefans
 * @Date: 2021-12-08 16:11:18
 * @LastEditTime: 2022-08-20 17:49:32
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @Description: nacos 相关代码；
 * @FilePath: /gw_micro_service/utils/nacos/nacos.go
 */

package nacos

import (
	"gin_micro_service/global"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type ServiceInstance struct {
	Ip   string `json:"ip"`
	Port uint64 `json:"port"`
}

func NacosInit() *naming_client.INamingClient {

	client, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": []constant.ServerConfig{
			{
				IpAddr: global.GVA_CONFIG.Nacos.NacosIp,
				Port:   uint64(global.GVA_CONFIG.Nacos.NacosPort),
			},
		},
		"clientConfig": constant.ClientConfig{
			TimeoutMs:           5000,
			ListenInterval:      10000,
			NotLoadCacheAtStart: true,
			LogDir:              global.GVA_CONFIG.Nacos.NacosLogdir,
			//Username:			 "nacos",
			//Password:			 "nacos",
		},
	})
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}
	return &client
}

/**
 * @Description: 服务注册
 * @Author: stefans
 * @Date: 2021-12-08 19:45:03
 * @LastEditors: stefans
 * @LastEditTime: Do not edit
 * @param {naming_client.INamingClient} client
 */
func RegisterService(client naming_client.INamingClient) bool {
	success, err := client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          global.GVA_CONFIG.Nacos.RegisterIp,
		Port:        uint64(global.GVA_CONFIG.System.Addr),
		ServiceName: global.GVA_CONFIG.Nacos.RegisterServicename,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		//ClusterName: "cluster-a", // default value is DEFAULT
		//GroupName:   "group-a",   // default value is DEFAULT_GROUP
	})
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return false
	}
	return success
}

func GetService(client naming_client.INamingClient, serviceName string) ServiceInstance {
	serviceInstance := ServiceInstance{
		Ip:   "0.0.0.0",
		Port: 9888,
	}
	services, err := client.GetService(vo.GetServiceParam{
		ServiceName: serviceName,
		//Clusters:    []string{"cluster-a"}, // default value is DEFAULT
		//GroupName:   "group-a",             // default value is DEFAULT_GROUP
	})
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return serviceInstance
	}
	Hosts := services.Hosts
	for _, host := range Hosts {
		if host.Healthy {
			serviceInstance.Ip = host.Ip
			serviceInstance.Port = host.Port
			break
		}
	}
	return serviceInstance
}
