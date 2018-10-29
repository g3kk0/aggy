package binance

import (
	"context"
	"strconv"
)

func (c *Client) GetBalances() (map[string]float64, error) {
	balances := make(map[string]float64)

	resp, err := c.Conn.NewGetAccountService().Do(context.Background())
	if err != nil {
		return balances, err
	}

	for _, b := range resp.Balances {
		f, err := strconv.ParseFloat(b.Free, 64)
		if err != nil {
			return balances, err
		}

		if f != 0 {
			balances[b.Asset] = f
		}
	}

	return balances, nil
}
