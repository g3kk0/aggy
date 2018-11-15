package gdax

import (
	"strconv"
)

func (c *Client) GetAccounts() (map[string]float64, error) {
	accounts := make(map[string]float64)
	var err error

	accs, err := c.Conn.GetAccounts()
	if err != nil {
		return accounts, err
	}

	for _, a := range accs {
		amount, err := strconv.ParseFloat(a.Available, 64)
		if err != nil {
			return accounts, err
		}

		if amount != 0 {
			accounts[a.Currency] = amount
		}
	}

	return accounts, nil
}
