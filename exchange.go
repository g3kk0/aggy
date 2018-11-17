package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/g3kk0/aggy/binance"
	"github.com/g3kk0/aggy/gdax"
	coinmarketcap "github.com/g3kk0/go-coinmarketcap"
	forex "github.com/g3kk0/go-forex"
)

type Exchange struct {
	Type       string
	Key        string
	Secret     string
	Passphrase string
}

type Account struct {
	Exchange      string  `json:"exchange"`
	Symbol        string  `json:"symbol"`
	Price         float64 `json:"price"`
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

func (e *Exchange) Value(cmcKey, quoteCurrency string) ([]Account, error) {
	var accounts []Account
	var err error

	switch e.Type {
	case "gdax":
		gc := gdax.NewClient(e.Key, e.Secret, e.Passphrase)

		accs, err := gc.GetAccounts()
		if err != nil {
			return accounts, err
		}

		// do this better (avoid the copy & duplication)
		for k, v := range accs {
			account := Account{
				Exchange: "gdax",
				Symbol:   k,
				Amount:   v,
			}

			accounts = append(accounts, account)
		}
	case "binance":
		bc := binance.NewClient(e.Key, e.Secret)

		accs, err := bc.GetAccounts()
		if err != nil {
			return accounts, err
		}

		for k, v := range accs {
			account := Account{
				Exchange: "binance",
				Symbol:   k,
				Amount:   v,
			}

			accounts = append(accounts, account)
		}
	default:
		return accounts, errors.New("unknown exchange type")
	}

	// loop over all assets and get values
	var cryptos []string

	for i, a := range accounts {
		var fiatSymbol bool
		for _, e := range fiatSymbols {
			if a.Symbol == e {
				fiatSymbol = true
				break
			}
		}

		if a.Symbol == quoteCurrency {
			a.Value = a.Amount
		} else if fiatSymbol {
			value, err := fiatValue(a.Symbol, quoteCurrency, a.Amount)
			if err != nil {
				return accounts, err
			}

			accounts[i].QuoteCurrency = quoteCurrency
			accounts[i].Value = value
		} else {
			cryptos = append(cryptos, a.Symbol)
		}
	}

	err = cryptoValue(accounts, cryptos, cmcKey, quoteCurrency)
	if err != nil {
		return accounts, err
	}

	return accounts, err
}

func cryptoValue(accounts []Account, cryptos []string, cmcKey, quoteCurrency string) error {
	var err error

	cmc := coinmarketcap.NewClient(cmcKey)

	params := map[string]string{"symbol": strings.Join(cryptos, ","), "convert": quoteCurrency}

	quotes, err := cmc.QuotesLatest(params)
	if err != nil {
		return err
	}

	for _, c := range cryptos {
		for i, a := range accounts {
			if c == a.Symbol {
				price := quotes.Data[a.Symbol].Quote[quoteCurrency].Price
				value := (price / 100000000) * (a.Amount * 100000000)

				accounts[i].Price = price
				accounts[i].QuoteCurrency = quoteCurrency
				accounts[i].Value = value
			}
		}
	}

	return err
}

func fiatValue(from, to string, amount float64) (float64, error) {
	fc := forex.NewClient()

	amountStr := fmt.Sprintf("%f", amount)

	params := map[string]string{"from": from, "to": to, "amount": amountStr}

	c, err := fc.Convert(params)
	if err != nil {
		return 0, err
	}

	return c.Result, nil
}
