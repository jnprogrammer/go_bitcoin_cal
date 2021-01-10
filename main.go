/*
1 Bitcoin can be divided into 100,000,000 Million Satoshis
This calculator is to help calculate trades from alt coins into BTC
and arbitrage for profit in Satoshis.


Phase 1:
	build out functional MVP

Phase 2:
	Refine MVP and expand to fulfill the focus of this calculators
	use for the front end.

Phase 3:
	Once basic front end is added, add APIs for top Crypto markets to
	pull pricing data for calculations

Phase 4:
	touch up Front-end and make available for public use.


MISC:
Test out Humanize library for commas in large numbers
https://github.com/dustin/go-humanize
*/

package main

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"strconv"
)

var btc float64
var btcsym = "₿"

//var adasym = "₳"
//var ada uint32
var testmarketprice float64
var testsats float64
var testcost float64

func main() {
	btc = 100000000
	testmarketprice = 34022.24
	//ada = 1000000
	testsats = 24.239949980
	testcost = testsats * testmarketprice

	//fmt.Printf(" = .%v\n"+btcsym, btc)
	//fmt.Printf(  "%v is the price of ₿itcoin ",testcost)
	//strconv.FormatFloat(testcost,'f', -2, 64)
	fmt.Printf("$%v is the Cost of ₿itcoin\n", strconv.FormatFloat(testcost, 'f', 2, 64))
	fmt.Printf("At $%v USD\n", testmarketprice)
	fmt.Printf("This is equalivannt to "+btcsym+"%v satoshis ", testsats)
	//fmt.Printf(adasym + " = %d\n", ada)

}