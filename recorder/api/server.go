package api

import (
	"net/http"

	"github.com/calamity-m/fernio/lib/middleware"
	"github.com/calamity-m/fernio/lib/server"
	"github.com/gin-gonic/gin"
)

func Serve(s *server.Server) {
	s.Log.Info("Starting to serve")

	r := gin.New()
	r.Use(middleware.LoggingMiddleware(s.Log))

	r.GET("/pong", func(c *gin.Context) {
		s.Log.Debug("whooooo i am in pong")
		c.String(http.StatusOK, "pong")
		s.Log.Debug("whoooooooooo i have set c.String")
	})

	r.Run(":8000")

}
