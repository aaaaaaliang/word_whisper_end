package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"word_whisper_end/internal/controller/fronted"
	"word_whisper_end/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "word_whisper_end",
		Usage: "word_whisper_end",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 未验证的组
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware().MiddlewareHandlerResponse)
				group.Bind(
					fronted.NewUserV1(),
				)
			})
			// 已验证的组
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().JWTMiddleware,
				)
				group.Middleware(service.Middleware().MiddlewareHandlerResponse)
				group.Bind(
					fronted.NewHomeV1(),
				)
			})

			s.Run()
			return nil
		},
	}
)
