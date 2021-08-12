package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nurmanhabib/go-rest-skeleton/pkg/response"
)

func handleError(e *gin.Engine) {
	e.GET("error", func(c *gin.Context) {
		response.JSON(c.Writer, response.ErrorResponse{
			Message: "Sample Error Message",
		})
	})
}
