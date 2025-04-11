package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/Remrinrin/bot-youmiya-main-go/internal/controller/hello"
	"github.com/Remrinrin/bot-youmiya-main-go/internal/controller/onebot"
	"github.com/Remrinrin/bot-youmiya-main-go/internal/middleware"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(middleware.TokenMiddleware)
				group.Bind(
					hello.NewV1(),
					onebot.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
