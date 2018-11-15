package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Value         float64   `json:"value"`
	QuoteCurrency string    `json:"quote_currency"`
	FiatIn        float64   `json:"fiat_in"`
	FiatOut       float64   `json:"fiat_out"`
	Holdings      []Account `json:"holdings"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "UI")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	var resp Response

	gdaxSecret := os.Getenv("GDAX_SECRET")
	gdaxKey := os.Getenv("GDAX_KEY")
	gdaxPassphrase := os.Getenv("GDAX_PASSPHRASE")
	binanceKey := os.Getenv("BINANCE_KEY")
	binanceSecret := os.Getenv("BINANCE_SECRET")
	cmcKey := os.Getenv("COINMARKETCAP_KEY")
	quoteCurrency := "GBP"

	ge := NewExchange("gdax", gdaxKey, gdaxSecret, gdaxPassphrase)
	be := NewExchange("binance", binanceKey, binanceSecret, "")

	geValue, err := ge.Value(cmcKey, quoteCurrency)
	if err != nil {
		panic(err)
	}

	beValue, err := be.Value(cmcKey, quoteCurrency)
	if err != nil {
		panic(err)
	}

	resp.Holdings = append(geValue, beValue...)

	var totalValue float64
	for _, a := range resp.Holdings {
		totalValue = totalValue + a.Value
	}

	resp.QuoteCurrency = quoteCurrency
	resp.Value = totalValue

	json, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/api", apiHandler)
	fmt.Println("starting web server")
	log.Fatal(http.ListenAndServe(":9999", nil))
}
