package main

import (
	"fmt"
	"os"

	"github.com/g3kk0/aggy/binance"
	"github.com/g3kk0/aggy/gdax"
)

func main() {
	gdaxSecret := os.Getenv("GDAX_SECRET")
	gdaxKey := os.Getenv("GDAX_KEY")
	gdaxPassphrase := os.Getenv("GDAX_PASSPHRASE")
	binanceKey := os.Getenv("BINANCE_KEY")
	binanceSecret := os.Getenv("BINANCE_SECRET")

	gc := gdax.NewClient(gdaxKey, gdaxSecret, gdaxPassphrase)
	bc := binance.NewClient(binanceKey, binanceSecret)

	gdaxTransfers := gc.GetTransfers()
	fmt.Printf("gdaxTransfers = %+v\n", gdaxTransfers)

	gdaxBalances := gc.GetBalances()
	fmt.Printf("gdaxBalances = %+v\n", gdaxBalances)

	binanceBalances := bc.GetBalances()
	fmt.Printf("binanceBalances = %+v\n", binanceBalances)

	//	http.HandleFunc("/", handler)
	//	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println("done")
}
