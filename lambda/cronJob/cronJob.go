package cronJob

import (
	"context"
	"fmt"
	"lambda/database"
	"lambda/types"
	"time"

	yf "github.com/shoenig/yahoo-finance"
)

type CronJobHandler struct {
	dbStore database.DynamoDBClient
}

func NewCronJobHandler(dbStore database.DynamoDBClient) CronJobHandler {
	return CronJobHandler{
		dbStore: dbStore,
	}
}

func (cronJob CronJobHandler) GetStocksPrice(ctx context.Context) error {
	tickers := []string{"GOOGL", "AMZN", "AAPL", "META", "MSFT", "NFLX"}
	stocksPrice := map[string]float64{}

	client := yf.New(nil)

	for _, ticker := range tickers {
		chart, err := client.Lookup(ticker)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		stocksPrice[ticker] = chart.Price()
	}

	err := cronJob.dbStore.AddStocksPrice(types.StocksPrice{
		Timestamp: time.Now().Format(time.RFC3339),
		Prices:    stocksPrice,
	})
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}
