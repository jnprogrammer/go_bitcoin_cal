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
10,000,000 BTC = 1,000,000,000,000,000 One Quadrillion 16 digits

Maximum Limit
21,000,000 BTC = 21,000,000,000,000,000 Twenty One Quadrillion

Phase 1:
	build out functional MVP slice
	COMPLETED 1-10-21

Phase 2:
	Expand MVP to allow user input, prepare for API integration for live price updates
	I may ned to use all ints and deal with calculations manually to represent the multiple place values from 100M to 1 Quadrillion

	I may need to toggle if a user is starting with a BTC value amount or
	a Satoshi value amount.
	So I figured out that to convert any amount I need to multiply by 1,000,000,000
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
	"github.com/gofiber/fiber"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var btcsym = "₿"

//var adasym = "₳"

var userInput float64

func main() {
	//btc = 100000000
	//ada = 1000000
	/*
				Adding User input will in my opinion need to let the use enter numbers without needing the decimal.
			Why?
				Because I'm crazy is why!! ha. . .We are focusing on Satoshis. This is the most common unit of
				 Bitcoin the majority of users will interact with, So to help with this they will enter satoshis as whole
			     numbers and I'll deal with conversion with the decimal in the background.. . . I'll find other ways to make this harder with
				 feature creep.

			1. A function that allows you to enter one satoshi amount and another satoshi amount that would be
			used in various arithmetic operations like finding the difference or multiplying them by a particular function.

		I'll use int64s and track the placement manually as a start.
				testcost = testfloat * testmarketprice
	*/

	fmt.Println("Enter an amount of ₿itcoin")
	fmt.Scan(&userInput)

	p := message.NewPrinter(language.English)

	fmt.Printf(btcsym+" %v is equivalent to \n", userInput)
	p.Printf("丰%d satoshis\n", convertToSat(userInput))

}

func setupRoutes(app *fiber.App) {

}
func convertToSat(amt float64) int64 {

	//needs take in a BTC amount and convert it to satoshis
	var onehm float64 = 100000000
	var total int64

	total = int64(amt * onehm)

	return total
}

func convertToLace(amt float64) int64 {

	//needs take in a ADA amount and convert it to lovelaces
	var oneht float64 = 1000000
	var total int64

	total = int64(amt * oneht)

	return total
}
