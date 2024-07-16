package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/ntc-goer/parser-exercise/internal/subscribe"
	"github.com/ntc-goer/parser-exercise/pkg/eth"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type Handler struct {
	Service          *Service
	SubscribeService *subscribe.Service
	ETH              *eth.ETH
}

func NewHandler(s *Service, ss *subscribe.Service, eth *eth.ETH) *Handler {
	return &Handler{
		Service:          s,
		SubscribeService: ss,
		ETH:              eth,
	}
}

func (h *Handler) GetByAddress(ctx *gin.Context) {
	addr := ctx.Query("address")
	if addr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing address parameter"})
		return
	}
	// Check the address still be subscribed
	sub, err := h.SubscribeService.GetOne(ctx, bson.M{"address": addr})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !sub.DeletedAt.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "This address is not subscribed"})
		return
	}
	transactions, err := h.Service.GetByAddress(ctx, addr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": transactions})
}

func (h *Handler) GetCurrentBlockNumber(ctx *gin.Context) {
	blockNum, err := h.ETH.GetLatestBlockNumber()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"blockNumber": blockNum})
}
