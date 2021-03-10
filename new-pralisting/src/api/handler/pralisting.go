package handler

import (
	"new-pralisting/src/api/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PralistingHandler ...
type PralistingHandler struct {
	PralistingEntity domain.PralistingEntity
}

// NewPralistingHandler ...
func NewPralistingHandler(r *gin.RouterGroup, us domain.PralistingEntity) {
	handler := &PralistingHandler{
		PralistingEntity: us,
	}
	pralisting := r.Group("/pra-listing")
	pralisting.GET("/", handler.Fetch)
}

// Fetch ...
func (a *PralistingHandler) Fetch(c *gin.Context) {
	pralisting, _ := a.PralistingEntity.Fetch(c.Request.Context())
	c.JSON(200, gin.H{"data": pralisting})
}

// GetByID ...
func (a *PralistingHandler) GetByID(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	pralisting, _ := a.PralistingEntity.GetByID(c, id)
	c.JSON(200, gin.H{"data": pralisting})
}
