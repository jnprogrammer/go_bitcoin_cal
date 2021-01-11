/*
1 Bitcoin can be divided into 100,000,000 Million Satoshis
This calculator is to help calculate trades from alt coins into BTC
and arbitrage for profit in Satoshis.

1 BTC = 100,000,000 One Hundred Million
10 BTC = 1,000,000,000 One Billion
100 BTC = 10,000,000,000 Ten Billion
1000 BTC = 100,000,000,000 One Hundred Billion
10,000 BTC = 1,000,000,000,000 One Trillion
100,000 BTC = 10,000,000,000,000 Ten Trillion
1,000,000 BTC = 100,000,000,000,000 One Hundred Trillion
10,000,000 BTC = 1,000,000,000,000,000 One Quadrillion

Maximum Limit
21,000,000 BTC = 21,000,000,000,000,000 Twenty One Quadrillion

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

naa
*/

package main

import (
	"fmt"
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
	testmarketprice = 41000.24
	//ada = 1000000
	testsats = .000249980
	testcost = testsats * testmarketprice

	//fmt.Printf(" = .%v\n"+btcsym, btc)
	//fmt.Printf(  "%v is the price of ₿itcoin ",testcost)
	//strconv.FormatFloat(testcost,'f', -2, 64)
	fmt.Printf(btcsym+"%v satoshis is equalivannt to ", strconv.FormatFloat(testsats, 'f', 9, 64))

	fmt.Printf("$%v USD with ₿itcoin Market price ", strconv.FormatFloat(testcost, 'f', 2, 64))

	fmt.Printf("at $%v USD\n", strconv.FormatFloat(testmarketprice, 'f', 2, 32))
	//fmt.Printf(adasym + " = %d\n", ada)

}
