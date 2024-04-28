package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Grabs the request id header from incoming request, then appends it to the gin context and response headers.
// If the header is not present or empty, a UUID formatted id will be generated and used instead.
// If mut is true any newly generaded request id will appended to the incoming request
func RequestId(header string, mut bool) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		reqId := ctx.Request.Header.Get(header)
		if reqId == "" {
			reqId = uuid.NewString()
			if mut {
				ctx.Request.Header.Set(header, reqId)
			}
		}

		ctx.Header(header, reqId)
		ctx.Set(header, reqId)

		ctx.Next()
	}

}
