package routes

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spinel/gophkeeper-storage-svc/pkg/config"
	"github.com/spinel/gophkeeper-storage-svc/pkg/services"
)

// Booking contains binded and validated data.
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func GetEntity(c *gin.Context) {
	err := errors.New("er")
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

}

func RegisterRoutes(r *gin.Engine, c *config.Config, svc services.Server) {
	routes := r.Group("/entity")
	routes.POST("/create", svc.CreateEntityHttp)
	routes.GET("/:id", svc.FindOne)
}
