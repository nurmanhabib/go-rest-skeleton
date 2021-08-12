package routers

import (
	"time"

	"github.com/nurmanhabib/go-rest-skeleton/pkg/response"

	"github.com/nurmanhabib/go-rest-skeleton/config"

	"github.com/gin-gonic/gin"
)

type pingResponse struct {
	Env       string    `json:"env"`
	Timestamp time.Time `json:"timestamp"`
}

func handlePing(e *gin.Engine, conf *config.Config) {
	e.GET("ping", func(c *gin.Context) {
		data := &pingResponse{Env: conf.Env, Timestamp: time.Now()}
		response.JSON(c.Writer, response.SuccessResponse{Data: data})
	})
}
