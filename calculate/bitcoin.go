package calculate

import (
	"github.com/gofiber/fiber"
)

func convertToSat(amt float64) int64 {

	//needs take in a BTC amount and convert it to satoshis
	var onehm float64 = 100000000
	var total int64

	total = int64(amt * onehm)

	return total
}

func returnSat(c *fiber.Ctx) {

	c.Send("Sats")
}
