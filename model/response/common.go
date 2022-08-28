/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-21 12:05:40
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-21 12:07:54
 * @FilePath: /gin_micro_service/model/response/common.go
 * @Description: encasulate return function
 */

package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = "-1"
	SUCCESS = "0"
)

func Result(code string, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func FailWithMessageV2(code string, message string, c *gin.Context) {
	Result(code, map[string]interface{}{}, message, c)
}

func FailWithDetailedV2(code string, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}
