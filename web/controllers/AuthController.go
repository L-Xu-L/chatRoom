package controllers

import (
	"chatRoom/conventions"
	"chatRoom/response"
	"chatRoom/util/wechatPublic"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type AuthController struct {
	Ctx iris.Context
}

/**
	获取微信公众平台第三方授权的url
*/
func (c *AuthController) GetUrl() mvc.Result{
	state := c.Ctx.URLParam("state")
	//不存在state直接报错
	if state == "" {
		return response.InvalidParam("param state should be required",nil)
	} else {
		return response.QuerySuccess("",&response.Data{
			Item: map[string]interface{}{
				"url":wechatPublic.GetCodeUrl("http://"+ conventions.AUTH_REDIRECT_URL + "/login",state),
			},
		})
	}
}

