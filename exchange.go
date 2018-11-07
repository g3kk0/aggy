package main

import (
	"errors"
	"fmt"

	"github.com/g3kk0/aggy/gdax"
)

type Exchange struct {
	Type       string
	Key        string
	Secret     string
	Passphrase string
}

type Response struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
	// define type for this!
	//	Exchanges []
	Holdings []Asset `json:"holdings"`
	//Holdings []gdax.Asset
}

type Asset struct {
	Exchange      string  `json:"exchange"`
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
	Amount        float64 `json:"amount"`
	QuoteCurrency string  `json:"quote_currency"`
	Value         float64 `json:"value"`
}

func NewExchange(t, key, secret, passphrase string) *Exchange {
	return &Exchange{
		Type:       t,
		Key:        key,
		Secret:     secret,
		Passphrase: passphrase,
	}
}

func (e *Exchange) Value(quoteCurrency string) (Response, error) {
	var r Response
	var err error

	fiatSymbols := []string{"EUR", "GBP", "USD"}
	cryptos := map[string]string{}
	//cryptos := map[string]map[string]string{}

	switch e.Type {
	case "gdax":
		gc := gdax.NewClient(e.Key, e.Secret, e.Passphrase)
		assets, err := gc.Assets(quoteCurrency)
		if err != nil {
			return r, err
		}

		// do this better (avoid the copy)
		for _, a := range assets {
			asset := Asset{
				Exchange: "gdax",
				Symbol:   a.Symbol,
				Amount:   a.Amount,
			}

			r.Holdings = append(r.Holdings, asset)
		}
	case "binance":
		//bc := binance.NewClient(e.Key, e.Secret)
		fmt.Println("checking binance")
	default:
		return r, errors.New("unknown exchange type")
	}

	// get a list of cryptos
	for _, h := range r.Holdings {

		var fiatSymbol bool
		for _, e := range fiatSymbols {
			if h.Symbol == e {
				fiatSymbol = true
				break
			}
		}

		if fiatSymbol {
			fmt.Printf("fiat = %+v\n", h.Symbol)
		} else {
			fmt.Printf("crypto = %+v\n", h.Symbol)
			cryptos[h.Symbol] = ""
			//cryptoSymbols = append(cryptoSymbols, h.Symbol)

		}

	}

	fmt.Printf("cryptos = %+v\n", cryptos)

	return r, err
}
