package food

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/calamity-m/fernio/pkg/server"
	"github.com/gin-gonic/gin"
)

func baseHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Base")
	}
}

func testHandler(s *server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		s.Log.Debug("Hit testing handler")
		// ctx.Request.Response.
		ctx.JSON(http.StatusOK, gin.H{
			"response": "Testing fuction hit",
		})
	}
}

func AddGroup(r *gin.Engine, s *server.Server, prefix string) error {
	if !strings.HasPrefix(prefix, "/") {
		return fmt.Errorf("prefix must start with a slash")
	}

	g := r.Group(prefix + "/food")
	g.GET("", baseHandler())
	g.GET("/test", testHandler(s))

	return nil
}
