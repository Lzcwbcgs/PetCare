package cmd

import (
	"context"

	adminctl "PetCare/internal/controller/admin"
	appointmentctl "PetCare/internal/controller/appointment"
	commonctl "PetCare/internal/controller/common"
	doctorctl "PetCare/internal/controller/doctor"
	medicalrecordctl "PetCare/internal/controller/medicalrecord"
	notificationctl "PetCare/internal/controller/notification"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"

	"PetCare/internal/controller/auth"
	"PetCare/internal/controller/pet"
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
				s                       = g.Server()
				appointmentController   = appointmentctl.New()
				doctorAppointmentCtl    = doctorctl.NewAppointment()
				doctorMedicalRecordCtl  = doctorctl.NewMedicalRecord()
				doctorMedicalReportCtl  = doctorctl.NewReport()
				doctorNotificationCtl   = doctorctl.NewNotification()
				publicAuth              = auth.NewPublic()
				privateAuth             = auth.NewPrivate()
				userController          = auth.NewUser()
				doctorController        = auth.NewDoctor()
				adminController         = auth.NewAdmin()
				adminDoctorController   = adminctl.NewDoctor()
				adminHospitalController = adminctl.NewHospital()
				adminNotificationCtl    = adminctl.NewNotification()
				commonController        = commonctl.New()
				medicalRecordController = medicalrecordctl.New()
				notificationController  = notificationctl.New()
				petController           = pet.New()
			)

			s.SetAddr(g.Cfg().MustGet(ctx, "server.address", ":8000").String())
			s.SetOpenApiPath(g.Cfg().MustGet(ctx, "server.openapiPath", "/api.json").String())
			s.SetSwaggerPath(g.Cfg().MustGet(ctx, "server.swaggerPath", "/swagger").String())
			s.AddStaticPath("/reports", "resource/public/uploads/reports")
			s.AddStaticPath("/uploads", "resource/public/uploads")
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
				group.Group("/users", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth, middleware.Role("user"))
					group.Bind(
						userController,
					)
				})
				group.Group("/doctors", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth, middleware.Role("doctor"))
					group.Bind(
						doctorController,
					)
				})
				group.Group("/doctor/appointments", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth, middleware.Role("doctor"))
					group.Bind(
						doctorAppointmentCtl,
					)
				})
				group.Group("/doctor", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth, middleware.Role("doctor"))
					group.Bind(
						doctorMedicalRecordCtl,
						doctorMedicalReportCtl,
						doctorNotificationCtl,
					)
				})
				group.Group("/admin", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth, middleware.Role("admin"))
					group.Bind(
						adminController,
						adminDoctorController,
						adminHospitalController,
						adminNotificationCtl,
					)
				})
				group.Group("/pets", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth)
					group.Bind(
						petController,
					)
				})
				group.Group("/appointments", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth)
					group.Bind(
						appointmentController,
					)
				})
				group.Group("/common", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth)
					group.Bind(
						commonController,
					)
				})
				group.Group("/medical-records", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth)
					group.Bind(
						medicalRecordController,
					)
				})
				group.Group("/notifications", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth)
					group.Bind(
						notificationController,
					)
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
