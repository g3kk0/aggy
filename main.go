package main

import (
	"fmt"
	"log"
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

	gdaxTransfers, err := gc.GetTransfers()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("gdaxTransfers = %+v\n", gdaxTransfers)

	gdaxBalances, err := gc.GetBalances()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("gdaxBalances = %+v\n", gdaxBalances)

	binanceBalances, err := bc.GetBalances()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("binanceBalances = %+v\n", binanceBalances)

	//	http.HandleFunc("/", handler)
	//	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println("done")
}
