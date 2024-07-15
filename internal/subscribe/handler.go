package subscribe

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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

func (h *Handler) Subscribe(ctx *gin.Context) {
	var req struct {
		Address string `json:"address"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := h.Service.Subscribe(ctx, req.Address)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "action": "subscribe", "address": req.Address})
}

func (h *Handler) UnSubscribe(ctx *gin.Context) {
	var req struct {
		Address string `json:"address"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := h.Service.UnSubscribe(ctx, req.Address)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Address %s haven't subcribe yet", req.Address)})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "action": "unsubscribe", "address": req.Address})
}

func (h *Handler) GetAll(ctx *gin.Context) {
	subs, err := h.Service.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "action": "GetAll", "subscribes": subs})
}
