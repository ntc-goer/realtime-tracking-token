package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ntc-goer/parser-exercise/internal/subscribe"
	"github.com/ntc-goer/parser-exercise/internal/transaction"
	"github.com/ntc-goer/parser-exercise/pkg/eth"
	"net/http"
)

type CoreHTTPServer struct {
	*HTTPServer
	ETHHandler         *eth.ETH
	SubscribeHandler   *subscribe.Handler
	TransactionHandler *transaction.Handler
}

func NewCoreHTTPServer(httpServer *HTTPServer, ethHandler *eth.ETH, subscribeHandler *subscribe.Handler, transactionHandler *transaction.Handler) *CoreHTTPServer {
	return &CoreHTTPServer{
		HTTPServer:         httpServer,
		ETHHandler:         ethHandler,
		SubscribeHandler:   subscribeHandler,
		TransactionHandler: transactionHandler,
	}
}

func (c *CoreHTTPServer) AddCoreRouter() {
	c.Engine.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to parser")
	})
	c.Engine.GET("current-block", c.TransactionHandler.GetCurrentBlockNumber)
	c.Engine.GET("subscribes", c.SubscribeHandler.GetAll)
	c.Engine.POST("subscribe", c.SubscribeHandler.Subscribe)
	c.Engine.POST("unsubscribe", c.SubscribeHandler.UnSubscribe)

	c.Engine.GET("transactions", c.TransactionHandler.GetByAddress)
}
