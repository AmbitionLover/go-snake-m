package cmd

import (
	"context"
	"github.com/gogf/gf/v2/protocol/goai"
	"go-snake/app/api/internal/consts"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"go-snake/app/api/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "api",
		Usage: "main",
		Brief: "The system mainly provides services",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Hello,
				)
			})
			// Custom enhance API document.
			enhanceOpenAPIDoc(s)
			// Just run the server.
			s.Run()
			return nil
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Version:     consts.OpenAPIVERSION,
		Contact: &goai.Contact{
			Name: "GoSnake",
			URL:  "https://ambi.vip",
		},
	}
}
