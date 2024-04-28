package api

import (
	"fmt"
	"net/http"

	"github.com/calamity-m/fernio/pkg/middleware"
	"github.com/calamity-m/fernio/pkg/server"
	"github.com/calamity-m/fernio/recorder/food"
	"github.com/gin-gonic/gin"
)

func Serve(s *server.Server) error {
	s.Log.Info("Starting to serve")

	r := gin.New()
	r.Use(
		middleware.RequestId(s.Config.RequestIdHeader, true),
		middleware.Logger(s.Log),
		gin.Recovery(),
	)

	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"v1": "Loaded"})
	})

	err := food.AddGroup(r, s, "/v1")
	if err != nil {
		return err
	}

	err = r.Run(fmt.Sprintf("%s:%v", s.Config.Host, s.Config.Port))
	return err
}
