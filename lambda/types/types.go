package types

type Prices map[string]float64

type StocksPrice struct {
	Timestamp string `json:"timestamp"`
	Prices    Prices `json:"prices"`
}
