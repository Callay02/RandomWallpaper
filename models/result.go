/*
 * @Author: Callay 2415993100@qq.com
 * @Date: 2024-05-07 14:28:51
 * @LastEditors: Callay 2415993100@qq.com
 * @LastEditTime: 2024-05-07 22:07:37
 * @FilePath: \checkIn\go\model\result.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package models

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Result) OK(message string) Result {
	r.Code = 200
	r.Message = message
	r.Data = ""
	return *r
}

func (r *Result) Error(message string) Result {
	r.Code = 400
	r.Message = message
	r.Data = ""
	return *r
}

func (r *Result) OKSetData(data interface{}) Result {
	r.Code = 200
	r.Message = "success"
	r.Data = data
	return *r
}

func (r *Result) OKSetMsgData(message string, data interface{}) Result {
	r.Code = 200
	r.Message = message
	r.Data = data
	return *r
}
