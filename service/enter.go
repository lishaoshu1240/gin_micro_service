/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-21 11:09:30
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-21 12:15:31
 * @FilePath: /finance_service/service/enter.go
 * @Description: enter.go
 */

package service

type servie struct {
	BaseFinanceOptService BaseFinanceOpt
}

var ServiceGroup = new(servie)
