package response

import (
	"chatRoom/conventions"
	"github.com/kataras/iris/mvc"
)

func QuerySuccess(message string,data interface{}) mvc.Response{
	response := new(baseResponse)
	return response.setCode(conventions.OK).setMessage(message).setData(data).responseJson()
}

func CreateSuccess(message string,data interface{}) mvc.Response{
	response := new(baseResponse)
	return response.setCode(conventions.CREATED).setMessage(message).setData(data).responseJson()
}

func AuthFailure(message string,data interface{}) mvc.Response{
	response := new(baseResponse)
	return response.setCode(conventions.FORBIDDEN).setMessage(message).setData(data).responseJson()
}

func InvalidParam(message string,data interface{}) mvc.Response{
	response := new(baseResponse)
	return response.setCode(conventions.BADREQUEST).setMessage(message).setData(data).responseJson()
}

func NotFound(message string,data interface{}) mvc.Response{
	response := new(baseResponse)
	return response.setCode(conventions.NOTFOUND).setMessage(message).setData(data).responseJson()
}

func NotModified(message string,data interface{}) mvc.Response{
	response := new(baseResponse)
	return response.setCode(conventions.NOTMODIFIED).setMessage(message).setData(data).responseJson()
}

func System(message string,data interface{}) mvc.Response{
	response := new(baseResponse)
	return response.setCode(conventions.SYSTEM).setMessage(message).setData(data).responseJson()
}