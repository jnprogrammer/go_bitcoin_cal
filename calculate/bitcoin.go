package calculate

func ConvertToSat(btc float64) int64 {

	//needs take in a BTC amount and convert it to satoshis
	var onehm float64 = 100000000
	var total int64

	total = int64(btc * onehm)

	return total
}
