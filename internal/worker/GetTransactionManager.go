package worker

import (
	"context"
	"github.com/ntc-goer/parser-exercise/config"
	"github.com/ntc-goer/parser-exercise/internal/subscribe"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type GetTransactionManager struct {
	Config           *config.Config
	Worker           *GetTransactionWorker
	SubscribeService *subscribe.Service
}

func NewGetTransactionManager(sb *subscribe.Service, wk *GetTransactionWorker, cfg *config.Config) *GetTransactionManager {
	return &GetTransactionManager{
		Config:           cfg,
		SubscribeService: sb,
		Worker:           wk,
	}
}

func (wk *GetTransactionManager) Run() {
	// Init WorkerQuantity worker
	ctx := context.TODO()
	var wg sync.WaitGroup

	for {
		// get all valid subscriber
		subscribers, err := wk.SubscribeService.GetAll(ctx)
		logrus.Infof("There are %d subcriber to track", len(subscribers))
		if err != nil {
			logrus.Error("Get subscribers fail")
			continue
		}
		if len(subscribers) == 0 {
			time.Sleep(15 * time.Second)
			continue
		}
		for i := 0; i < len(subscribers); i += wk.Config.WorkerQuantity {
			low := i
			high := low + wk.Config.WorkerQuantity
			if low+wk.Config.WorkerQuantity > len(subscribers) {
				high = len(subscribers)
			}
			partition := subscribers[low:high]
			for _, subscriber := range partition {
				wg.Add(1)
				go func(w *sync.WaitGroup, sub subscribe.Subscribe) {
					defer w.Done()
					wk.Worker.Run(ctx, sub)
				}(&wg, subscriber)
			}
			wg.Wait()
		}

		logrus.Infof("All check finished , waiting for next check")
		// Delay for the next checking
		time.Sleep(10 * time.Second)
	}
}
