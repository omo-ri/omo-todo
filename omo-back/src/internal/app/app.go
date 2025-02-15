package app

import (
	"github.com/gin-gonic/gin"

	"omo-back/src/internal/handler"
)

type App struct {
	router      *gin.Engine
	authHandler *handler.AuthHandler
}

type Handlers struct {
	AuthHandler *handler.AuthHandler
}

func NewApp(handlers Handlers) *App {
	r := gin.Default()
	r.Use(gin.Logger())

	// 注册自定义中间件
	//r.Use(middleware.Logger()) // 示例日志中间件

	// 注册路由
	registerRoutes(r, handlers)

	return &App{router: r}
}

func (a *App) Run(addr string) error {
	return a.router.Run(addr)
}

func registerRoutes(r *gin.Engine, handlers Handlers) {
	// 公开路由
	public := r.Group("/api")
	{
		public.GET("/ping", handler.Ping)
		public.POST("/register", handlers.AuthHandler.Register)
		public.POST("/login", handler.Login)
	}

	// 受保护路由（需JWT认证）
	//protected := r.Group("/api")
	//protected.Use(middleware.JWTAuth())
	//{
	//	protected.GET("/profile", handler.Profile)
	//}
}
