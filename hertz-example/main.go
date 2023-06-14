package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	jsonp "github.com/li-jin-gou/hertz-jsonp"
)

func main() {
	r := server.Default()
	r.Use(jsonp.JsonPHertz())
	r.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(200, utils.H{
			"message": "pong",
		})
	})
	r.Spin()
}
