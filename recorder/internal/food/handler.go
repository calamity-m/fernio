package food

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/calamity-m/fernio/pkg/persistence"
	"github.com/calamity-m/fernio/pkg/server"
	"github.com/gin-gonic/gin"
)

func GetFoodById(s *server.Server, repo persistence.StorageHandler[FoodDao]) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Request": "Missing ID"})
		}

		filt := persistence.Filter{Terms: map[string]any{"Hello": "there"}}

		ctx.JSON(http.StatusOK, gin.H{"get": "rtn", "err": "err", "fakefilt": filt})
	}

}

func baseHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"endpoint": map[string]interface{}{"food": "storage of food records and food items"}})
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
