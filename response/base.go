package response

import "github.com/kataras/iris/mvc"

type baseResponse struct {
	code int
	data interface{}
	message string
}

/**
	响应json
 */
func (this *baseResponse) responseJson() mvc.Response{
	return func(code int,message string,data interface{}) mvc.Response{
		return mvc.Response{
			Code:code,
			ContentType:"application/json",
			Content:[]byte(message),
		}
	}(this.code,this.message,this.data)
}

/**
	响应json
*/
func (this *baseResponse) setCode(code int) *baseResponse{
	this.code = code
	return this

}

/**
	响应json
*/
func (this *baseResponse) setMessage(message string) *baseResponse{
	this.message = message
	return this
}

/**
	响应json
*/
func (this *baseResponse) setData(data interface{}) *baseResponse{
	this.data = data
	return this
}