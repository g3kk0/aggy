package gdax

import (
	"log"
	"strconv"

	gdax "github.com/preichenberger/go-gdax"
)

func (c *Client) GetTransfers() map[string]float64 {
	var ledgers []gdax.LedgerEntry

	accounts, err := c.Conn.GetAccounts()
	if err != nil {
		log.Println(err)
	}

	transfers := make(map[string]float64)

	for _, a := range accounts {
		cursor := c.Conn.ListAccountLedger(a.Id)
		for cursor.HasMore {
			if err := cursor.NextPage(&ledgers); err != nil {
				log.Println(err)
			}

			for _, e := range ledgers {
				if e.Type == "transfer" && e.Details.ProductId == "" {
					f, err := strconv.ParseFloat(e.Amount, 64)
					if err != nil {
						log.Println(err)
					}

					transfers[a.Currency] += f
				}
			}
		}
	}

	return transfers
}
