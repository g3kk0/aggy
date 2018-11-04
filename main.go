package main

import (
	"fmt"
	"os"
)

func main() {

	gdaxSecret := os.Getenv("GDAX_SECRET")
	gdaxKey := os.Getenv("GDAX_KEY")
	gdaxPassphrase := os.Getenv("GDAX_PASSPHRASE")
	binanceKey := os.Getenv("BINANCE_KEY")
	binanceSecret := os.Getenv("BINANCE_SECRET")
	quoteCurrency := "GBP"

	ge := NewExchange("gdax", gdaxKey, gdaxSecret, gdaxPassphrase)
	be := NewExchange("binance", binanceKey, binanceSecret, "")

	geValue, err := ge.Value(quoteCurrency)
	if err != nil {
		panic(err)
	}

	beValue, err := be.Value(quoteCurrency)
	if err != nil {
		panic(err)
	}

	fmt.Printf("geValue = %+v\n", geValue)
	fmt.Printf("beValue = %+v\n", beValue)

	// gc := gdax.NewClient(gdaxKey, gdaxSecret, gdaxPassphrase)
	// bc := binance.NewClient(binanceKey, binanceSecret)

	// gdaxTransfers, err := gc.GetTransfers()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Printf("gdaxTransfers = %+v\n", gdaxTransfers)

	// gdaxBalances, err := gc.GetBalances()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Printf("gdaxBalances = %+v\n", gdaxBalances)

	// binanceBalances, err := bc.GetBalances()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Printf("binanceBalances = %+v\n", binanceBalances)

	// GetQuotes(gdaxBalances)

	//	http.HandleFunc("/", handler)
	//	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println("done")
}
