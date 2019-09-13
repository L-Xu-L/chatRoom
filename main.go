package main

import (
	"chatRoom/models/repositories"
	"chatRoom/models/services"
	"chatRoom/util/mysql"
	"chatRoom/web/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func newApp() *iris.Application{
	app := iris.New()
	mvc.Configure(app.Party("/login"), initLoginController)
	mvc.Configure(app.Party("/auth"), initAuthController)
	return app
}

func init(){
	mysql.NewPool() //初始化mysql连接池
}

func main()  {
	app := newApp()
	_ = app.Run(iris.Addr(":8000"))
}

/**
	注册中间件和完成各级依赖注入
 */
func initLoginController(app *mvc.Application) {
	// 使用数据源中的一些（内存）数据创建 movie 的数据库。
	repo := repositories.NewUserRepository()
	// 创建 movie 的服务，我们将它绑定到 movie 应用程序。
	userService := services.NewUserService(repo)
	app.Register(userService)
	app.Handle(new(controllers.LoginController))
}

/**
	注册中间件和完成各级依赖注入
*/
func initAuthController(app *mvc.Application) {
	app.Handle(new(controllers.LoginController))
}