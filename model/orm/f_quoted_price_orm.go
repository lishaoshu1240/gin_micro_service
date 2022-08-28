/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-22 10:03:48
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-22 10:03:51
 * @FilePath: /finance_service/model/orm/f_quoted_price_orm.go
 * @Description: quoted price
 */

package model

type F_quoted_price struct {
	BaseOrmWithId
	UserId           uint32  `json:"user_id"`
	OrderType        uint8   `json:"order_type"`
	QuotedType       uint8   `json:"quoted_type"`
	FeesInfoId       uint32  `json:"fees_info_id"`
	QuotedPrice      float32 `json:"quoted_price"`
	CustomerCode     string  `json:"customer_code"`
	PickupTerminalId uint32  `json:"pickup_terminal_id"`
	LoadFactoryId    uint32  `json:"load_factory_id"`
	UnloadFactoryId  uint32  `json:"unload_factory_id"`
	ReturnTerminalId uint32  `json:"return_terminal_id"`
	ContainerId      uint32  `json:"container_id"`
	MotorcadeId      uint32  `json:"motorcade_id"`
	IsActive         uint8   `json:"is_active"`
}
