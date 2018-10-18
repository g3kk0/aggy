package main

import (
	"fmt"
	"os"

	gdax "github.com/preichenberger/go-gdax"
)

func GetDeposits() {

	secret := os.Getenv("GDAX_SECRET")
	key := os.Getenv("GDAX_KEY")
	passphrase := os.Getenv("GDAX_PASSPHRASE")

	client := gdax.NewClient(secret, key, passphrase)

	var ledgers []gdax.LedgerEntry

	accounts, err := client.GetAccounts()
	if err != nil {
		println(err.Error())
	}

	// create map of all transfers e.g. EUR: {{},{}}

	for _, a := range accounts {
		cursor := client.ListAccountLedger(a.Id)
		for cursor.HasMore {
			if err := cursor.NextPage(&ledgers); err != nil {
				println(err.Error())
			}

			for _, e := range ledgers {
				if e.Type == "transfer" && e.Details.ProductId == "" {
					fmt.Printf("wallet = %+v\n", a.Currency)
					fmt.Printf("e = %+v\n\n", e)
				}
			}
		}
	}

}
