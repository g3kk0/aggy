package main

import (
	"fmt"
	"os"
)

func main() {

	//gdaxSecret := os.Getenv("GDAX_SECRET")
	//gdaxKey := os.Getenv("GDAX_KEY")
	//gdaxPassphrase := os.Getenv("GDAX_PASSPHRASE")
	binanceKey := os.Getenv("BINANCE_KEY")
	binanceSecret := os.Getenv("BINANCE_SECRET")
	cmcKey := os.Getenv("COINMARKETCAP_KEY")
	quoteCurrency := "GBP"

	//ge := NewExchange("gdax", gdaxKey, gdaxSecret, gdaxPassphrase)
	be := NewExchange("binance", binanceKey, binanceSecret, "")

	//geValue, err := ge.Value(cmcKey, quoteCurrency)
	//if err != nil {
	//	panic(err)
	//}

	beValue, err := be.Value(cmcKey, quoteCurrency)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("geValue = %+v\n", geValue)
	fmt.Printf("beValue = %+v\n", beValue)

	fmt.Println("done")
}
