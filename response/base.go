package response

import (
	"encoding/json"
	"github.com/kataras/iris/mvc"
)

type Data struct {
	Item map[string]interface{} //总体值
	Counting int64 //总计
	HasMore bool //是否还有更多数据
}

func (this *Data) GetItem() map[string]interface{}{
	if this.Item != nil {
		return this.Item
	}
	return nil
}

func (this *Data) GetCounting() int64{
	return this.Counting
}

func (this *Data) GetHasMore() bool {
	return this.HasMore
}

type baseResponse struct {
	code int
	data *Data
	message string
}

/**
	响应json
 */
func (this *baseResponse) responseJson() mvc.Response{
	return func(code int,message string,data *Data) mvc.Response{
		dataMap := make(map[string]interface{}) //构造响应给前端的data数据
		if this.data != nil {
			if this.data.GetCounting() > 0 {
				dataMap["counting"] = this.data.Counting
			}
			if this.data.GetHasMore() {
				dataMap["hasMore"] = this.data.HasMore
			}
			if this.data.GetItem() != nil {
				for key,value := range this.data.Item {
					dataMap[key] = value
				}
			}
		}
		responseData := map[string]interface{}{
			"message":message,
			"data":dataMap,
		}
		jsons, _ := json.Marshal(responseData)
		return mvc.Response{
			Code:code,
			ContentType:"application/json",
			Content:[]byte(jsons),
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
func (this *baseResponse) setData(data *Data) *baseResponse{
	this.data = data
	return this
}