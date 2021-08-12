package routers

import (
	"net/http"

	"github.com/nurmanhabib/go-rest-skeleton/infrastructure/message/exception"

	"github.com/gin-gonic/gin"
)

func handle404(e *gin.Engine) {
	e.NoRoute(func(c *gin.Context) {
		err := exception.ErrorTextNotFound
		_ = c.AbortWithError(http.StatusNotFound, err)
	})
}
