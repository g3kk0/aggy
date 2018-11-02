package gdax

import (
	"strconv"
)

func (c *Client) GetBalances() (map[string]float64, error) {
	balances := make(map[string]float64)

	accounts, err := c.Conn.GetAccounts()
	if err != nil {
		return balances, err
	}

	for _, a := range accounts {
		f, err := strconv.ParseFloat(a.Balance, 64)
		if err != nil {
			return balances, err
		}

		if f != 0 {
			balances[a.Currency] = f
		}
	}

	return balances, nil
}
