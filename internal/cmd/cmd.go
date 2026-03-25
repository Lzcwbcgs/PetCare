package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gcfg"

	"PetCare/internal/controller/auth"
	"PetCare/internal/middleware"
	"PetCare/internal/model"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("manifest/config/config.yaml")

			var (
				s              = g.Server()
				publicAuth     = auth.NewPublic()
				privateAuth    = auth.NewPrivate()
			)

			s.SetAddr(g.Cfg().MustGet(ctx, "server.address", ":8000").String())
			s.SetOpenApiPath(g.Cfg().MustGet(ctx, "server.openapiPath", "/api.json").String())
			s.SetSwaggerPath(g.Cfg().MustGet(ctx, "server.swaggerPath", "/swagger").String())
			s.Use(middleware.ResponseHandler)

			s.BindHandler("/", func(r *ghttp.Request) {
				r.Response.RedirectTo("/swagger/")
			})

			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Group("/auth", func(group *ghttp.RouterGroup) {
					group.Bind(
						publicAuth,
					)
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.Auth)
						group.Bind(
							privateAuth,
						)
					})
				})
			})
			enhanceOpenAPIDoc(s)
			s.Run()
			return nil
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = model.Response{}
	openapi.Config.CommonResponseDataField = `Data`
	openapi.Info = goai.Info{
		Title:       "PetCare API",
		Description: "宠物医疗协同管理平台接口文档",
	}
}
