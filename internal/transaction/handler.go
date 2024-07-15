package transaction

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) GetByAddress(ctx *gin.Context) {
	addr := ctx.Query("address")
	if addr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing address parameter"})
		return
	}
	transactions, err := h.Service.GetByAddress(ctx, addr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing address parameter"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": transactions})
}
