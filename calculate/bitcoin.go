package calculate

func convertToSat(btc float64) int64 {

	//needs take in a BTC amount and convert it to satoshis
	var onehm float64 = 100000000
	var total int64

	total = int64(btc * onehm)

	return total
}

func convertToLace(ada float64) int64 {

	//needs take in a ADA amount and convert it to lovelaces
	var onemil float64 = 1000000
	var total int64

	total = int64(ada * onemil)

	return total
}
