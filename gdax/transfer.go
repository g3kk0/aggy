package gdax

import (
	"strconv"

	gdax "github.com/preichenberger/go-gdax"
)

func (c *Client) GetTransfers() (map[string]float64, error) {
	transfers := make(map[string]float64)

	accounts, err := c.Conn.GetAccounts()
	if err != nil {
		return transfers, err
	}

	var ledgers []gdax.LedgerEntry

	for _, a := range accounts {
		cursor := c.Conn.ListAccountLedger(a.Id)
		for cursor.HasMore {
			if err := cursor.NextPage(&ledgers); err != nil {
				return transfers, err
			}

			for _, l := range ledgers {
				if l.Type == "transfer" && l.Details.ProductId == "" {
					amount, err := strconv.ParseFloat(l.Amount, 64)
					if err != nil {
						return transfers, err
					}

					transfers[a.Currency] += amount
				}
			}
		}
	}

	return transfers, nil
}
