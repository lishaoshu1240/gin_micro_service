/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-20 11:20:12
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-21 12:19:48
 * @FilePath: /gin_micro_service/api/v1/enter.go
 * @Description: */
package v1

type api struct {
	FeesInfoApi FeesInfoApi
}

var ApiGroup = new(api)
