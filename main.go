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
	build out functional MVP slice
	COMPLETED 1-10-21

Phase 2:
	Expand MVP to allow user input, prepare for API integration for live price updates


Phase 3:
	build out basic Front end, add APIs for top Crypto markets to
	pull pricing data for calculations. Create database for future Analytic possibilities.

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

//var btc float64
var btcsym = "₿"

//var adasym = "₳"
//var ada uint32
var testmarketprice float64
var testsats float64
var testcost float64
var testamt1 float64
var testamt2 float64
var testint int64

func main() {
	//btc = 100000000
	//ada = 1000000
	/*
			Adding User input will in my opinion need to let the use enter numbers without needing the decimal.

		Why?

			Because I'm crazy is why!! We are focusing on Satoshis. This is the most common unit of
			 Bitcoin the majority of users will interact with, So to help with this they will enter satoshis as whole
		     numbers and I'll deal with conversion with the decimal in the background.. . . I'll find other ways to make this harder with
			 feature creep.

		1. A function that allows you to enter one satoshi amount and another satoshie amount that would be
		used in various arithmetic operations like finding the difference or multiplying them by a particular function.
	*/
	testamt1 = 0.02911480
	testamt2 = 0.03135440
	testmarketprice = 41000.24
	testsats = .000249980
	testcost = testsats * testmarketprice

	testint = int64(testamt2)
	fmt.Println("Enter How many Satoshis you have as a whole number\n")
	fmt.Printf("Stop use this first. %v", testint)
	fmt.Printf(btcsym+"%v satoshis is equalivannt to ", strconv.FormatFloat(testsats, 'f', 9, 64))
	fmt.Printf("$%v USD with ₿itcoin Market price ", strconv.FormatFloat(testcost, 'f', 2, 64))
	fmt.Printf("at $%v USD\n", strconv.FormatFloat(testmarketprice, 'f', 2, 32))

}
