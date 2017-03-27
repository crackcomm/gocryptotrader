package main

import (
	//"log"
	"net/url"
	"reflect"
	"strconv"
	"testing"
)

func TestBitfinexGetTicker(t *testing.T) {
	t.Parallel()
	bitfinex := Bitfinex{}

	response, err := bitfinex.GetTicker("BTCUSD", url.Values{})
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(response.Timestamp).String() != "string" { //Types alread contained in struct.
		t.Error("Bitfinex ticker.Timestamp value is not a string variable")
	}
	if reflect.TypeOf(response.Ask).String() != "float64" {
		t.Error("Bitfinex ticker.Ask value is not a float64 variable")
	}
	if reflect.TypeOf(response.Bid).String() != "float64" {
		t.Error("Bitfinex ticker.Bid value is not a float64 variable")
	}
	if reflect.TypeOf(response.High).String() != "float64" {
		t.Error("Bitfinex ticker.High value is not a float64 variable")
	}
	if reflect.TypeOf(response.Last).String() != "float64" {
		t.Error("Bitfinex ticker.Last value is not a float64 variable")
	}
	if reflect.TypeOf(response.Low).String() != "float64" {
		t.Error("Bitfinex ticker.Low value is not a float64 variable")
	}
	if reflect.TypeOf(response.Mid).String() != "float64" {
		t.Error("Bitfinex ticker.Mid value is not a float64 variable")
	}
	if reflect.TypeOf(response.Volume).String() != "float64" {
		t.Error("Bitfinex ticker.Volume value is not a float64 variable")
	}

	responseTimestamp, err := strconv.ParseFloat(response.Timestamp, 64)
	if err != nil {
		t.Error("ticker.Timestamp value cannot be converted to an float64")
	}
	if responseTimestamp <= 0 { //Revise
		t.Error("ticker.Timestamp value is negative or zero")
	}
	if response.Ask < 0 {
		t.Error("ticker.Ask value is negative")
	}
	if response.Bid < 0 {
		t.Error("ticker.Bid value is negative")
	}
	if response.High < 0 {
		t.Error("ticker.High value is negative")
	}
	if response.Last < 0 {
		t.Error("ticker.Last value is negative")
	}
	if response.Low < 0 {
		t.Error("ticker.Low value is negative")
	}
	if response.Mid < 0 {
		t.Error("ticker.Mid value is negative")
	}
	if response.Volume < 0 {
		t.Error("ticker.ask value is negative")
	}
}

func TestBitfinexGetStats(t *testing.T) {
	t.Parallel()
	BitfinexGetStatsTest := Bitfinex{}

	response, err := BitfinexGetStatsTest.GetStats("BTCUSD")
	if err != nil {
		t.Error("BitfinexGetStatsTest init error: ", err)
	}

	for _, responsivity := range response { //had to put in a slice for the response

		if reflect.TypeOf(responsivity.Period).String() != "int64" { //Maybe not needed??
			t.Error("Bitfinex Getstats.Period is not an int64")
		}
		if reflect.TypeOf(responsivity.Volume).String() != "float64" {
			t.Error("Bitfiniex Getstats.Volume is not a float64")
		}

		if responsivity.Period <= 0 {
			t.Error("response.Period value is negative or zero")
		}
		if responsivity.Volume < 0 {
			t.Error("response.Volume value is negative")
		}
	}
}

func TestBitfinexGetLendbook(t *testing.T) { //BROKEN RETURNS
	t.Parallel()
	BitfinexGetLendbook := Bitfinex{}

	response, err := BitfinexGetLendbook.GetLendbook("BTC", url.Values{}) //Status code was not 200. Changed to BTC error here.
	if err != nil {
		t.Error("BitfinexGetLendbook init error: ", err)
	}

	for _, asks := range response.Asks {
		if reflect.TypeOf(asks.Amount).String() != "float64" {
			t.Error("Bitfinex GetLendbook.Asks.Amount is not a float64")
		}

		if reflect.TypeOf(asks.FlashReturnRate).String() != "string" { //Help here?
			t.Error("Bitfinex GetLendbook.Asks.FlashReturnRate is not a string")
		}
		if reflect.TypeOf(asks.Period).String() != "int" {
			t.Error("Bitfinex GetLendbook.Asks.Period is not an int")
		}
		if reflect.TypeOf(asks.Rate).String() != "float64" {
			t.Error("Bitfinex GetLendbook.Asks.Rate is not a string")
		}
		responseTimetamp, err := strconv.ParseFloat(asks.Timestamp, 64)
		if err != nil {
			t.Error("Could not convert Bitfinex GetLendbook.Asks.Timestamp into float64")
		}

		if asks.Amount <= 0 {
			t.Error("Bitfinex GetLendbook.Asks.Amount is negative or 0")
		}
		// if asks.FlashReturnRate <= ??????? {}
		if asks.Period <= 0 {
			t.Error("Bitfinex GetLendbook.Asks.Period is negative or 0")
		}
		if asks.Rate <= 0 {
			t.Error("Bitfinex GetLendbook.Asks.Rate is negative or 0")
		}
		if responseTimetamp <= 0 {
			t.Error("Bitfinex GetLendbook.Asks.Timestamp is negative or 0")
		}
	}

	for _, bids := range response.Bids {
		if reflect.TypeOf(bids.Amount).String() != "float64" {
			t.Error("Bitfinex GetLendbook.Asks.Amount is not a float64")
		}

		if reflect.TypeOf(bids.FlashReturnRate).String() != "string" { //Help here?
			t.Error("Bitfinex GetLendbook.Asks.FlashReturnRate is not a string")
		}
		if reflect.TypeOf(bids.Period).String() != "int" {
			t.Error("Bitfinex GetLendbook.Asks.Period is not an int")
		}
		if reflect.TypeOf(bids.Rate).String() != "float64" {
			t.Error("Bitfinex GetLendbook.Asks.Rate is not a string")
		}
		responseTimetamp, err := strconv.ParseFloat(bids.Timestamp, 64)
		if err != nil {
			t.Error("Could not convert Bitfinex GetLendbook.Asks.Timestamp into float64")
		}

		if bids.Amount <= 0 {
			t.Error("Bitfinex GetLendbook.Asks.Amount is negative or 0")
		}
		// if bids.FlashReturnRate <= ??????? {}
		if bids.Period <= 0 {
			t.Error("Bitfinex GetLendbook.Asks.Period is negative or 0")
		}
		if bids.Rate <= 0 {
			t.Error("Bitfinex GetLendbook.Asks.Rate is negative or 0")
		}
		if responseTimetamp <= 0 {
			t.Error("Bitfinex GetLendbook.Asks.Timestamp is negative or 0")
		}
	}
}

func TestBitfinexGetOrderbook(t *testing.T) {
	t.Parallel()
	BitfinexGetOrderbook := Bitfinex{}

	orderBook, err := BitfinexGetOrderbook.GetOrderbook("BTCUSD", url.Values{})
	if err != nil {
		t.Error("BitfinexGetOrderbook init error: ", err)
	}

	for _, asks := range orderBook.Asks {
		amount, err := strconv.ParseFloat(asks.Amount, 64)
		if err != nil {
			t.Error("Cannot convert Bitfinex Orderbook.Asks.Amount into a float64")
		}
		if amount < 0 {
			t.Error("Bitfinex Orderbook.Asks.Amount is negative")
		}
		price, err2 := strconv.ParseFloat(asks.Price, 64)
		if err2 != nil {
			t.Error("Cannot convert Bitfinex Orderbook.Asks.Price into a float64")
		}
		if price < 0 {
			t.Error("Bitfinex Orderbook.Asks.Price is negative")
		}
		timestamp, err3 := strconv.ParseFloat(asks.Timestamp, 64)
		if err3 != nil {
			t.Error("Cannot convert Bitfinex Orderbook.Asks.timestamp into a float64")
		}
		if timestamp <= 0 {
			t.Error("Bitfinex Orderbook.Asks.Amount is negative or 0")
		}
	}

	for _, bids := range orderBook.Bids {
		amount, err := strconv.ParseFloat(bids.Amount, 64)
		if err != nil {
			t.Error("Cannot convert Bitfinex Orderbook.bids.Amount into a float64")
		}
		if amount < 0 {
			t.Error("Bitfinex Orderbook.bids.Amount is negative")
		}
		price, err2 := strconv.ParseFloat(bids.Price, 64)
		if err2 != nil {
			t.Error("Cannot convert Bitfinex Orderbook.bids.Price into a float64")
		}
		if price < 0 {
			t.Error("Bitfinex Orderbook.bids.Price is negative")
		}
		timestamp, err3 := strconv.ParseFloat(bids.Timestamp, 64)
		if err3 != nil {
			t.Error("Cannot convert Bitfinex Orderbook.bids.timestamp into a float64")
		}
		if timestamp <= 0 {
			t.Error("Bitfinex Orderbook.bids.Amount is negative or 0")
		}
	}
}

func TestBitfinexGetTrades(t *testing.T) {
	t.Parallel()
	BitfinexGetTrades := Bitfinex{}

	trades, err := BitfinexGetTrades.GetTrades("BTCUSD", url.Values{})
	if err != nil {
		t.Error("BitfinexGetTrades init error: ", err)
	}

	for _, explicitTrades := range trades {
		amount, err := strconv.ParseFloat(explicitTrades.Amount, 64)
		if err != nil {
			t.Error("Cannot convert Bitfinex GetTrades.Amount into a float64")
		}
		if amount <= 0 {
			t.Error("Bitfinex GetTrades.Amount is negative or 0")
		}
		if explicitTrades.Exchange != "bitfinex" {
			t.Error("Bitfinex GetTrades.Exchange incorrect name")
		}
		price, err2 := strconv.ParseFloat(explicitTrades.Price, 64)
		if err2 != nil {
			t.Error("Cannot convert Bitfinex GetTrades.Price into a float64")
		}
		if price <= 0 {
			t.Error("Bitfinex GetTrades.Price is negative or 0")
		}
		if explicitTrades.Tid <= 0 {
			t.Error("Bitfinex GetTrades.Tid is negative or 0")
		}
		if explicitTrades.Timestamp <= 0 {
			t.Error("Bitfinex GetTrades.Timestamp is negative or 0")
		}
		// if explicitTrades.Type == "buy" || explicitTrades.Type == "sell" { //SUPER BRAIN FART!
		// 	t.Error("Bitfinex GetTrades.Type is wrong")
		// }
	}
}

func TestBitfinexGetLends(t *testing.T) {
	t.Parallel()
	BitfinexGetLends := Bitfinex{}

	lends, err := BitfinexGetLends.GetLends("BTC", url.Values{}) //status 200 on BTCUSD not BTC
	if err != nil {
		t.Error("BitfinexGetLends init error: ", err)
	}

	for _, explicitLends := range lends {
		if explicitLends.AmountLent <= 0 {
			t.Error("Bitfinex GetLends.AmountLent is negative or 0")
		}
		if explicitLends.AmountUsed <= 0 {
			t.Error("Bitfinex GetLends.AmountUsed is negative or 0")
		}
		if explicitLends.Rate <= 0 {
			t.Error("Bitfinex GetLends.Rate is negative or 0")
		}
		if explicitLends.Timestamp <= 0 {
			t.Error("Bitfinex GetLends.Timestamp is negative or 0")
		}
	}
}

func TestBitfinexGetSymbols(t *testing.T) {
	t.Parallel()
	BitfinexGetSymbols := Bitfinex{}

	symbols, err := BitfinexGetSymbols.GetSymbols()
	if err != nil {
		t.Error("BitfinexGetSymbols init error: ", err)
	}

	for _, explicitSymbol := range symbols {
		if reflect.TypeOf(explicitSymbol).String() != "string" {
			t.Error("Bitfinex GetSymbols is not a string")
		}
	}
}

func TestBitfinexGetSymbolsDetails(t *testing.T) {
	t.Parallel()
	BitfinexGetSymbolsDetails := Bitfinex{}

	symboldetails, err := BitfinexGetSymbolsDetails.GetSymbolsDetails()
	if err != nil {
		t.Error("BitfinexGetSymbolsDetails init error: ", err)
	}

	for _, explicitDetails := range symboldetails {
		if explicitDetails.Expiration != "NA" {
			expiration, err := strconv.ParseFloat(explicitDetails.Expiration, 64)
			if err != nil {
				t.Error("Cannot convert Bitfinex GetSymbolsDetails.Expiration into a float64")
			}
			if expiration < 0 {
				t.Error("Bitfinex GetSymbolsDetails.Expiration is negative")
			}
		}
		if explicitDetails.InitialMargin <= 0 {
			t.Error("Bitfinex GetSymbolsDetails.InitialMargin is negative or 0")
		}
		if explicitDetails.MaximumOrderSize <= 0 {
			t.Error("Bitfinex GetSymbolsDetails.MaximumOrderSize is negative or 0")
		}
		if explicitDetails.MinimumMargin <= 0 {
			t.Error("Bitfinex GetSymbolsDetails.MinimumMargin is negative or 0")
		}
		if explicitDetails.MinimumOrderSize <= 0 {
			t.Error("Bitfinex GetSymbolsDetails.MinimumOrderSize is negative or 0")
		}
		if len(explicitDetails.Pair) != 6 {
			t.Error("Bitfinex GetSymbolsDetails.Pair incorrect length")
		}
		if explicitDetails.PricePrecision <= 0 {
			t.Error("Bitfinex GetSymbolsDetails.PricePrecision is negative or 0") //Could go less could go higher than 5 but 5 is ideal
		}
	}
}

func TestBitfinexGetAccountInfo(t *testing.T) {
	t.Parallel()
	BitfinexGetAccountInfo := Bitfinex{}

	accountInfo, err := BitfinexGetAccountInfo.GetAccountInfo()
	if err != nil {
		t.Error("BitfinexGetAccountInfo init error: ", err)
	}

	for _, explicitAI := range accountInfo {
		makerFees, err := strconv.ParseFloat(explicitAI.MakerFees, 64)
		if err != nil {
			t.Error("Cannot convert Bitfinex GetAccountInfo.MakerFees into float64")
		}
		if makerFees < 0 {
			t.Error("Bitfinex GetAccountInfo.MakerFees is negative")
		}

		takerFees, err := strconv.ParseFloat(explicitAI.TakerFees, 64)
		if err != nil {
			t.Error("Cannot convert Bitfinex GetAccountInfo.TakerFees into float64")
		}
		if takerFees < 0 {
			t.Error("Bitfinex GetAccountInfo.TakerFees is negative")
		}

		for _, fees := range explicitAI.Fees {
			MakerFees, err := strconv.PasrseFloat(fees.MakerFees, 64)
			if err != nil {
				t.Error("Cannot convert Bitfinex GetAccountInfo.Fees.MakerFees into float64")
			}
			if MakerFees < 0 {
				t.Error("Bitfinex GetAccountInfo.Fees.MakerFees is negative")
			}
			if len(fees.Pairs) != 6 {
				t.Error("Bitfinex GetAccountInfo.Fees.Pairs incorrect length")
			}
			TakerFees, err := strconv.ParseFloat(fees.TakerFees, 64)
			if err != nil {
				t.Error("Cannot convert Bitfinex GetAccountInfo.Fees.TakerFees into float64")
			}
			if TakerFees < 0 {
				t.Error("Bitfinex GetAccountInfo.Fees.TakerFees is negative")
			}
		}
	}
}
