package gdax

import (
	"log"
	"strconv"
)

func (c *Client) GetBalances() map[string]float64 {
	accounts, err := c.Conn.GetAccounts()
	if err != nil {
		log.Println(err)
	}

	balances := make(map[string]float64)

	for _, a := range accounts {
		f, err := strconv.ParseFloat(a.Balance, 64)
		if err != nil {
			log.Println(err)
		}

		balances[a.Currency] = f
	}

	return balances
}
