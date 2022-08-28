/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-21 11:01:42
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-21 11:06:09
 * @FilePath: /finance_service/model/orm/base_orm.go
 * @Description: base orm
 */

package model

import "time"

type BaseOrm struct {
	CreatedTime time.Time  `json:"created_time"`
	UpdatedTime *time.Time `json:"updated_time"`
}

type BaseOrmWithId struct {
	ID          uint32     `json:"id"`
	CreatedTime time.Time  `json:"created_time"`
	UpdatedTime *time.Time `json:"updated_time"`
}
