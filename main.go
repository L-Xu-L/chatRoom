package main

import (
	"chatRoom/models/data"
	"chatRoom/models/repositories"
	"chatRoom/models/services"
	"chatRoom/util/mysql"
	"chatRoom/util/redis"
	"chatRoom/web/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
)

func newApp() *iris.Application{
	app := iris.New()
	//配置全局跨越中间件
	app.Use(func(context context.Context) {
		context.Header("Access-Control-Allow-Origin","*")
		context.Header("Access-Control-Allow-Methods","POST,GET,OPTIONS,DELETE")
		context.Header("Access-Control-Allow-Headers","Content-Type,Content-Length,Accept-Encoding,X-Requested-with,Origin,token")
		context.Next()
	})
	mvc.Configure(app.Party("/login"), initLoginController)
	mvc.Configure(app.Party("/auth"), initAuthController)
	return app
}

func init(){
	mysql.NewPool() //初始化mysql连接池
	redis.NewConnetion() //初始化redis连接池
}

func main()  {
	app := newApp()
	app.RegisterView(iris.HTML("./web/views", ".html"))
	_ = app.Run(iris.Addr(":8000"))
}


/**
	注册中间件和完成各级依赖注入,注册登录控制器
 */
func initLoginController(app *mvc.Application) {
	//数据库访问层
	repo := repositories.NewUserRepository()
	//业务层
	userService := services.NewUserService(repo)
	//redis数据对象
	authData := data.NewAuthData()
	app.Register(userService)
	app.Register(authData)
	app.Handle(new(controllers.LoginController))
}

/**
	注册中间件和完成各级依赖注入,注册鉴权控制器
*/
func initAuthController(app *mvc.Application) {
	app.Handle(new(controllers.AuthController))
}