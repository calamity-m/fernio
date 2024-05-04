package food

import (
	"net/http"

	"github.com/calamity-m/fernio/pkg/server"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetFoodRecordById(s *server.Server, repo FoodRepository) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Request": "Missing ID"})
		}

		rtn, err := repo.GetById(id)
		if err != nil {
			// Should check if err was a not found one
			ctx.JSON(http.StatusInternalServerError, gin.H{"Internal server error": err})
		}

		filter := Filter{
			Limit:  0,
			Offset: 0,
			Terms:  map[string]any{"id": uuid.New()},
		}

		blublub, err := repo.GetByFilter(filter)

		ctx.JSON(http.StatusOK, gin.H{"get": rtn, "err": err, "secondget": blublub})
	}

}
