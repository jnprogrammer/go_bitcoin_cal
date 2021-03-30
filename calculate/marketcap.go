package calculate

func findMk(ada float64) int64 {

	//needs take in a ADA amount and convert it to lovelaces
	var onemil float64 = 1000000
	var total int64

	total = int64(ada * onemil)

	return total
}
