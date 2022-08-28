/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-21 10:37:51
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-22 10:04:12
 * @FilePath: /finance_service/model/orm/f_fees_info.go
 * @Description: orm table
 */

package model

type Tabler interface {
	TableName() string
}

func (F_fees_info) TableName() string {
	return "f_fees_info"
}

type F_fees_info struct {
	BaseOrmWithId
	UserId            uint32 `json:"user_id"`
	FeeClassification uint8  `json:"fee_classification"`
	IsCommon          uint8  `json:"is_common"`
	FeeName           string `json:"fee_name"`
	FeeCode           string `json:"fee_code"`
	FeeType           uint8  `json:"fee_type"`
	FeeProperty       uint8  `json:"fee_property"`
	PaymentUnit       uint8  `json:"payment_uint"`
	OtherPaymentObjId uint32 `json:"other_payment_obj_id"`
	IsActive          uint8  `json:"is_active"`
}
