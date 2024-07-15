package cmd

import (
	"context"
	"github.com/ntc-goer/parser-exercise/config"
	"github.com/ntc-goer/parser-exercise/internal/subscribe"
	"github.com/ntc-goer/parser-exercise/internal/transaction"
	"github.com/ntc-goer/parser-exercise/internal/worker"
	"github.com/ntc-goer/parser-exercise/pkg/database"
	"github.com/ntc-goer/parser-exercise/pkg/eth"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

func workerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "worker",
		Short: "",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancelFn := context.WithCancel(context.Background())
			defer cancelFn()
			container := provideWorkerDependencies()

			err := container.Invoke(func(manager *worker.GetTransactionManager) {
				manager.Run()
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

func provideWorkerDependencies() *dig.Container {
	c := dig.New()
	_ = c.Provide(context.Background)
	c.Provide(func() *config.Config {
		return cfg
	})
	c.Provide(database.NewMongoDB)
	c.Provide(database.InitMongoDB)

	c.Provide(eth.NewETH)

	// Subscriber
	c.Provide(subscribe.NewService)
	c.Provide(subscribe.NewRepository)

	// Transaction
	c.Provide(transaction.NewService)
	c.Provide(transaction.NewRepository)

	c.Provide(worker.NewGetTransactionManager)
	c.Provide(worker.NewGetTransactionWorker)
	return c
}
