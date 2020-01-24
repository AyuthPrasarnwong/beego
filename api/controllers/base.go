package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

//Response 结构体
type Response struct {
	Errcode int         `json:"errcode,omitempty"`
	Errmsg  string      `json:"errmsg,omitempty"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Errors struct {
		Code    int         `json:"status_code"`
		Message string      `json:"message"`
		Errors  interface{} `json:"errors,omitempty"`
	} `json:"errors"`
}

type Errors struct {
	Code    int         `json:"status_code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

type PaginateResponse struct {
	Data    interface{} `json:"data"`
	Meta 	interface{} `json:"meta,omitempty"`
}