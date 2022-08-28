/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-21 11:50:23
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-21 12:19:54
 * @FilePath: /gin_micro_service/api/v1/finance/fees_info_api.go
 * @Description: fees_info api */

package v1

import (
	"gin_micro_service/model/response"
	"gin_micro_service/service"

	"github.com/gin-gonic/gin"
)

type FeesInfoApi struct {
}

//define service
var baseFinanceOptService = service.ServiceGroup.BaseFinanceOptService

func (f *FeesInfoApi) GetFeesInfo(c *gin.Context) {
	ret := baseFinanceOptService.FeesInfoGet(1)
	response.OkWithData(ret, c)
}
