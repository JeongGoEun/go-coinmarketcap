// NOTE: before testing, apply your CMC API key to init() of `coinmarketcap_test.go`
package coinmarketcap

import (
	"testing"

	"github.com/hexoul/go-coinmarketcap/types"
)

func TestCryptoInfo(t *testing.T) {
	ret, err := GetInstance().CryptoInfo(&types.Options{
		Symbol: "BTC",
	})
	if err != nil {
		t.Fatal(err)
	}
	if ret.CryptoInfo["BTC"].Name != "Bitcoin" {
		t.FailNow()
	}
}

func TestCryptoMap(t *testing.T) {
	ret, err := GetInstance().CryptoMap(&types.Options{
		Symbol: "BTC",
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(ret.CryptoMap) == 0 || ret.CryptoMap[0].Name != "Bitcoin" {
		t.FailNow()
	}
}

func TestCryptoListingsLatest(t *testing.T) {
	listings, err := GetInstance().CryptoListingsLatest(&types.Options{
		Limit: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(listings.CryptoMarket) == 0 {
		t.FailNow()
	}
	if listings.CryptoMarket[0].Name != "Bitcoin" {
		t.FailNow()
	}
	if listings.CryptoMarket[0].Quote["USD"].Price <= 0 {
		t.FailNow()
	}
}

func TestCryptoMarketPairsLatest(t *testing.T) {
	info, err := GetInstance().CryptoMarketPairsLatest(&types.Options{
		Symbol: "BTC",
	})
	if err != nil {
		t.Fatal(err)
	}
	if info.Name != "Bitcoin" {
		t.FailNow()
	}
}

func TestCryptoMarketQuotesLatest(t *testing.T) {
	quotes, err := GetInstance().CryptoMarketQuotesLatest(&types.Options{
		Symbol: "BTC,ETH",
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(quotes.CryptoMarket) == 0 || len(quotes.CryptoMarket) != 2 {
		t.FailNow()
	}
	if quotes.CryptoMarket["BTC"].Name != "Bitcoin" || quotes.CryptoMarket["ETH"].Name != "Ethereum" {
		t.FailNow()
	}
}
