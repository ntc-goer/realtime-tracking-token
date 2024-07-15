package cmd

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ntc-goer/parser-exercise/config"
	"github.com/ntc-goer/parser-exercise/internal/server"
	"github.com/ntc-goer/parser-exercise/internal/subscribe"
	"github.com/ntc-goer/parser-exercise/internal/transaction"
	"github.com/ntc-goer/parser-exercise/pkg/database"
	"github.com/ntc-goer/parser-exercise/pkg/eth"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

func serverCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancelFn := context.WithCancel(context.Background())
			defer cancelFn()
			container := provideCoreDependencies()

			// Start server
			err := container.Invoke(func(s *server.CoreHTTPServer) {
				s.AddCoreRouter()
				err := s.Start()
				if err != nil {
					return
				}
			})
			if err != nil {
				return err
			}

			var client *mongo.Client
			err = container.Invoke(func(db *mongo.Database) {
				client = db.Client()
			})
			if err != nil {
				return err
			}
			defer func() {
				if err = client.Disconnect(ctx); err != nil {
					panic(err)
				}
			}()

			return nil
		},
	}
}

func provideCoreDependencies() *dig.Container {
	c := dig.New()
	err := c.Provide(gin.New)
	_ = c.Provide(context.Background)
	c.Provide(server.NewHTTPServer)
	c.Provide(server.NewCoreHTTPServer)
	if err != nil {
		return nil
	}
	c.Provide(func() *config.Config {
		return cfg
	})
	c.Provide(database.NewMongoDB)
	c.Provide(database.InitMongoDB)

	// Subscriber
	c.Provide(subscribe.NewHandler)
	c.Provide(subscribe.NewService)
	c.Provide(subscribe.NewRepository)

	// Transaction
	c.Provide(transaction.NewHandler)
	c.Provide(transaction.NewService)
	c.Provide(transaction.NewRepository)

	c.Provide(eth.NewETH)
	return c
}
