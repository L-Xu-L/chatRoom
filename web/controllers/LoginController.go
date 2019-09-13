package controllers

import (
	"chatRoom/models/logical"
	"chatRoom/models/services"
	"chatRoom/response"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"time"
)

type LoginController struct {
	Ctx iris.Context
	services.UserService
}

func (c *LoginController) Get() mvc.Result {
	//todo 回去使用service层调用model实现把数据写入数据库中,然后转发到聊天室主页,把websocket的数据放入redis或者map中,构建组号
	code := c.Ctx.URLParam("code")	//获取code
	state := c.Ctx.URLParam("state") //获取state
	if code == "" && state == "" {
		return response.InvalidParam("error param","")
	}
	userData := logical.BuildUserInserData(code)
	if user,err := c.UserService.GetUserByOpenid(userData.Openid);err != nil {
		if user == nil {
			user, err := c.UserService.Login(*userData)
			if err != nil {
				return response.System("system error","")
			}
			return response.CreateSuccess("登录成功!",user)
		}
		user.LoginAt = time.Now().Unix()
		return response.CreateSuccess("登录成功!",user)
	} else {
		log.Fatalf("have error where getUserByOpenid #%v",err)
		return response.System("system error","")
	}
}

