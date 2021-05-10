package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "everything is ok!")
	}
}
