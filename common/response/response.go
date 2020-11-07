package response

import (
	"goBlog/config/code"
	"strconv"
)

//http响应体

type response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// return response struct pointer.
func Response() *response {
	return &response{}
}

// return error response struct by code.
func NewErrResponse(c int) *response {
	if _, ok := codeCfg.Lang[c]; ok {
		return &response{Code: c, Msg: codeCfg.Lang[c]}
	}

	return &response{Code: c, Msg: "error: no find msg in code:" + strconv.Itoa(c)}
}

// return success response struct by result data
func NewSuccessResponse(d interface{}) *response {
	return &response{Code: codeCfg.Success, Msg: codeCfg.Lang[codeCfg.Success], Data: d}
}

// return a one page response struct.
func NewPageResponse(d interface{}, total int) *response {
	res := map[string]interface{}{}
	res["total"] = total
	res["data"] = d

	return &response{Code: codeCfg.Success, Msg: codeCfg.Lang[codeCfg.Success], Data: res}
}

// return response struct by code and result data.
func NewResponse(c int, d interface{}) *response {
	return &response{Code: c, Msg: codeCfg.Lang[c], Data: d}
}

// set response struct data.
func (r *response) SetData(d interface{}) *response {
	r.Data = d
	return r
}

// set response struct code.
func (r *response) SetCode(c int) *response {
	r.Code = c
	return r
}

// set response struct msg.
func (r *response) SetMsg(m string) *response {
	r.Msg = m
	return r
}
