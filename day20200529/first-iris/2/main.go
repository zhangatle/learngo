package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.Use(recover2.New())
	app.Use(logger.New())

	mvc.New(app).Handle(new(ExampleController))
	//app.Run(iris.Addr(":8080"))
	app.Run(iris.AutoTLS(":443","zhangatle.imwork.net","admin@zhangatle.imwork.net"))
}

type ExampleController struct {}

func (c *ExampleController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text: "<h1>Welcome</h1>",
	}
}

func (c *ExampleController) GetPing() string {
	return "pong"
}

func (c *ExampleController) GetHello() interface{}  {
	return map[string]string{"message": "hello Iris!"}
}