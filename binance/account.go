package binance

import (
	"context"
	"strconv"
)

func (c *Client) GetAccounts() (map[string]float64, error) {
	accounts := make(map[string]float64)

	resp, err := c.Conn.NewGetAccountService().Do(context.Background())
	if err != nil {
		return accounts, err
	}

	for _, b := range resp.Balances {
		f, err := strconv.ParseFloat(b.Free, 64)
		if err != nil {
			return accounts, err
		}

		if f != 0 {
			accounts[b.Asset] = f
		}
	}

	// clean these symbols: EON,EOP,IOTA

	return accounts, nil
}
