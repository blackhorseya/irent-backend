package main

import (
	"flag"
)

var path = flag.String("c", "configs/app.yaml", "set config file path")

func init() {
	flag.Parse()
}

// @title IRent API
// @version 0.0.1
// @description IRent API
//
// @contact.name Sean Cheng
// @contact.email blackhorseya@gmail.com
// @contact.url https://blog.seancheng.space
//
// @license.name GPL-3.0
// @license.url https://spdx.org/licenses/GPL-3.0-only.html
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
//
// @BasePath /api
func main() {
	app, err := CreateApp(*path)
	if err != nil {
		panic(err)
	}

	if err = app.Start(); err != nil {
		panic(err)
	}

	app.AwaitSignal()
}
