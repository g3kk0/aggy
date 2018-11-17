package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/g3kk0/aggy/gdax"
)

type Response struct {
	Value    float64   `json:"value"`
	Currency string    `json:"currency"`
	FiatIn   float64   `json:"fiat_in"`
	Pnl      float64   `json:"pnl"`
	PnlPc    float64   `json:"pnl_pc"`
	Holdings []Account `json:"holdings"`
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

	// fix this client duplication
	gc := gdax.NewClient(gdaxKey, gdaxSecret, gdaxPassphrase)

	transfers, err := gc.GetTransfers()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	for k, v := range transfers {
		if k == quoteCurrency {
			resp.FiatIn = resp.FiatIn + v
		} else {
			for _, f := range fiatSymbols {
				if k == f {
					value, err := fiatValue(k, quoteCurrency, v)
					if err != nil {
						http.Error(w, err.Error(), 500)
					}

					resp.FiatIn = resp.FiatIn + value
				}
			}
		}
	}

	geValue, err := ge.Value(cmcKey, quoteCurrency)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	beValue, err := be.Value(cmcKey, quoteCurrency)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	resp.Holdings = append(geValue, beValue...)

	var totalValue float64
	for _, a := range resp.Holdings {
		totalValue = totalValue + a.Value
	}

	resp.Currency = quoteCurrency
	resp.Value = totalValue
	resp.Pnl = resp.Value - resp.FiatIn
	resp.PnlPc = (resp.Pnl / resp.FiatIn) * 100

	json, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

var fiatSymbols = []string{"EUR", "GBP", "USD"}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api", apiHandler)
	fmt.Println("starting web server")
	log.Fatal(http.ListenAndServe(":9999", nil))
}
