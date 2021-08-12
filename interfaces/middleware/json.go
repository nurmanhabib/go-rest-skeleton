package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nurmanhabib/go-rest-skeleton/config"
	"github.com/nurmanhabib/go-rest-skeleton/pkg/response"
)

type JSONResponse struct {
	Conf *config.Config
}

func NewJSONResponse(c *config.Config) *JSONResponse {
	return &JSONResponse{Conf: c}
}

func (r JSONResponse) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset: utf-8")
		c.Next()

		if c.Errors.Last() == nil {
			return
		}

		err := c.Errors.Last().Err
		c.Errors = c.Errors[:0]
		message := err.Error()

		if r.Conf.Env == "production" && c.Writer.Status() == 500 {
			message = "an_error_occurred"
			response.JSON(c.Writer, response.FailureResponse{Message: message})
			return
		}

		response.JSON(c.Writer, response.FailureResponse{
			Message: message,
		})
	}
}
