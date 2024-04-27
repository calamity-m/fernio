package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestId(header string) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		reqId := ctx.Request.Header.Get(header)
		if reqId == "" {
			reqId = uuid.NewString()
			ctx.Request.Header.Set(header, reqId)
		}

		ctx.Set(header, reqId)

		ctx.Next()
	}

}
