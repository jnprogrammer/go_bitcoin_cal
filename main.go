package main

/*
Bitcoin can be divided into 100,000,000 Million Satoshis
This calculator is to help calculate trades from alt coins into BTC
and arbitrage for profit in Satoshis.

*/
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
	testmarketprice = 34022.24
	//ada = 1000000
	testsats = .2399499
	testcost = testsats * testmarketprice

	//fmt.Printf(" = .%v\n"+btcsym, btc)
	//fmt.Printf(  "%v is the price of ₿itcoin ",testcost)
	//strconv.FormatFloat(testcost,'f', -2, 64)
	fmt.Printf("$%v is the price of ₿itcoin ", strconv.FormatFloat(testcost, 'f', 2, 64))

	//fmt.Printf(adasym + " = %d\n", ada)

}
