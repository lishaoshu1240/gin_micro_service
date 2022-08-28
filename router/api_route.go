/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-21 12:12:02
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-21 12:27:25
 * @FilePath: /finance_service/router/api_route.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

package router

import (
	v1 "finance_service/api/v1"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	//apiRouter := Router.Group("api").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("finance")
	var feesInfoApi = v1.ApiGroup.FeesInfoApi
	{
		apiRouterWithoutRecord.POST("getAllApis", feesInfoApi.GetFeesInfo) // 获取所有api
	}
}
