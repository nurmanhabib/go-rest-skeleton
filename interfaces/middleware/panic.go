package middleware

import (
	"net/http"

	"github.com/nurmanhabib/go-rest-skeleton/infrastructure/message/exception"

	"github.com/gin-gonic/gin"
	"github.com/rollbar/rollbar-go"
)

type Panic struct {
}

func NewPanic() *Panic {
	return &Panic{}
}

func (p Panic) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer p.recover(c)
		c.Next()
	}
}

func (Panic) recover(c *gin.Context) {
	if err := recover(); err != nil {
		rollbar.Critical(err, c.Request)
		rollbar.Wait()

		_ = c.AbortWithError(http.StatusInternalServerError, exception.ErrorTextInternalServerError)
		return
	}
}
