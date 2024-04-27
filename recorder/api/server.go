package api

import (
	"fmt"
	"net/http"

	"github.com/calamity-m/fernio/pkg/middleware"
	"github.com/calamity-m/fernio/pkg/server"
	"github.com/gin-gonic/gin"
)

func Serve(s *server.Server) {
	s.Log.Info("Starting to serve")

	r := gin.New()
	r.Use(
		middleware.RequestId(s.Config.RequestIdHeader),
		middleware.Logger(s.Log),
	)

	r.GET("/pong", func(c *gin.Context) {
		s.Log.DebugContext(c, "Pong endpoint hit")
		c.String(http.StatusOK, "ping")
	})

	r.Run(fmt.Sprintf("%s:%v", s.Config.Host, s.Config.Port))

}
