package controllers

import (
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
func (c *LoginController) GetUrl() mvc.Result{
	state := c.Ctx.URLParam("state") //获取state
	//不存在state直接报错
	if state == "" {
		return response.InvalidParam("param state should be required","")
	} else {
		return response.QuerySuccess("{\"message\":\""+ wechatPublic.GetCodeUrl("http://yry.chatroom.top:8000/login",state) +"\",\"data\":{\"userId\":5}}","")
	}
}

