package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nurmanhabib/go-rest-skeleton/config"
	"github.com/nurmanhabib/go-rest-skeleton/interfaces/middleware"
)

type Router struct {
	conf *config.Config
}

func New(conf *config.Config) *Router {
	return &Router{
		conf: conf,
	}
}

func (r Router) Init() *gin.Engine {
	if r.conf.IsDebugMode() == false {
		gin.SetMode(gin.ReleaseMode)
	}

	e := gin.Default()
	e.Use(middleware.NewJSONResponse(r.conf).Middleware())
	e.Use(middleware.NewPanic().Middleware())

	handle404(e)
	handlePing(e, r.conf)
	handleError(e)

	// Your routes here
	// handleAwesome(e)
	handleUsers(e)

	return e
}
