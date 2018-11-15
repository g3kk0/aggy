package binance

import (
	"context"
	"strconv"
)

func (c *Client) GetAccounts() (map[string]float64, error) {
	accounts := make(map[string]float64)

	accs, err := c.Conn.NewGetAccountService().Do(context.Background())
	if err != nil {
		return accounts, err
	}

	for _, a := range accs.Balances {
		switch {
		case a.Asset == "EON" || a.Asset == "EOP":
			continue
		case a.Asset == "IOTA":
			a.Asset = "MIOTA"
		}

		amount, err := strconv.ParseFloat(a.Free, 64)
		if err != nil {
			return accounts, err
		}

		if amount != 0 {
			accounts[a.Asset] = amount
		}
	}

	return accounts, nil
}
