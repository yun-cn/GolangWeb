package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	_ = app.Run(iris.Addr("127.0.0.1:8080"))
}
