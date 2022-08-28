/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-21 11:13:43
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-22 09:50:19
 * @FilePath: /gin_micro_service/service/fees_info_crud.go
 * @Description: 财务资料-》费用信息*/

package service

import (
	"fmt"
	"gin_micro_service/global"
	model "gin_micro_service/model/orm"
)

//base finance operator,egs: add finance item; receivable/payable price
type BaseFinanceOpt struct {
}

func (b *BaseFinanceOpt) CreateFeesInfo() {

}

/**
 * @Description: all fees info get
 * @Author: stefan
 * @Date: 2022-08-21 11:35:06
 * @LastEditors: stefan
 * @LastEditTime: Do not edit
 * @param {uint32} user_id
 */
func (b *BaseFinanceOpt) FeesInfoGet(user_id uint32) (res model.F_fees_info) {

	global.GVA_DB.Select("*").Find(&res)
	fmt.Printf("res: %v\n", res)
	return res
}

func (b *BaseFinanceOpt) FeesInfoGetById(user_id, id uint32) {

}
