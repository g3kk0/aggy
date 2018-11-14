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
		switch {
		case b.Asset == "EON" || b.Asset == "EOP":
			continue
		case b.Asset == "IOTA":
			b.Asset = "MIOTA"
		}

		f, err := strconv.ParseFloat(b.Free, 64)
		if err != nil {
			return accounts, err
		}

		if f != 0 {
			accounts[b.Asset] = f
		}
	}

	return accounts, nil
}
