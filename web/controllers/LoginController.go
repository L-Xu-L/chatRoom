package controllers

import (
	"chatRoom/models/data"
	"chatRoom/models/logical"
	"chatRoom/models/services"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"time"
)

type LoginController struct {
	Ctx iris.Context
	services.UserService
	data.AuthData
}

func (c *LoginController) Get() {
	//todo 回去使用service层调用model实现把数据写入数据库中,然后转发到聊天室主页,把websocket的数据放入redis或者map中,构建组号
	code := c.Ctx.URLParam("code")	//获取code
	state := c.Ctx.URLParam("state") //获取state
	if code == "" && state == "" {
		_ = c.Ctx.View("login_error.html")
	}

	userData := logical.BuildUserInserData(code)
	if userData == nil {
		_ = c.Ctx.View("login_error.html")
		return
	}

	if user,err := c.UserService.GetUserByOpenid(userData.Openid);err == nil {
		//根据openid获取用户信息,如果用户不在数据库中则实现注册
		if user == nil {
			user, err = c.UserService.Login(*userData)
			if err != nil {
				golog.Infof("have error where user login #%v",err)
				_ = c.Ctx.View("login_error.html")
			}
		}
		//更新最近一次登录时间
		go func() {
			user.LoginAt = time.Now().Unix()
			c.UserService.Save(user)
		}()
		//注册token
		token := c.AuthData.GenerateToken(user)
		// 渲染模板文件： ./views/hello.html
		c.Ctx.ViewData("token",token)
		c.Ctx.ViewData("state",state)
		_ = c.Ctx.View("login_success.html")
	} else {
		golog.Warnf("have error where getUserByOpenid #%v",err)
		_ = c.Ctx.View("login_error.html")
	}
}

