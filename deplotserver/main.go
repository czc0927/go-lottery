package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"os"
	"os/exec"
)
// 文件日志
var logger *log.Logger


type baseController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&baseController{})
	initLog()
	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8000"))
}

func reLaunch()  {
	cmd := exec.Command("sh", "./deploy.sh")
	err := cmd.Start()
	if err != nil{
		log.Fatal(err)
	}
	err = cmd.Wait()
}

// Post http://localhost:8000/
func (c *baseController) Post() string {
	reLaunch()
	return fmt.Sprintf("<h1>Hello , this is my reploy!</h1>")
}

// 初始化日志信息
func initLog() {
	f, _ := os.Create("/tmp/lottery_deploy.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

