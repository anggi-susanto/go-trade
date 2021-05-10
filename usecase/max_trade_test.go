package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxTrade_CalculateProfit(t *testing.T) {
	t.Parallel()

	t.Run("#1.", func(t *testing.T) {
		cekArr := []int{5, 4, 3, 2, 1}
		diff, _ := CalculateProfit(cekArr)
		exp := &Calculation{
			MaxProfit: 0,
			BuyHour:   1,
			BuyPrice:  5,
			SellHour:  1,
			SellPrice: 5,
		}
		assert.Equal(t, exp, diff)
	})
	t.Run("#2.", func(t *testing.T) {
		cekArr := []int{3, 2, 1, 5, 6, 2}
		diff, _ := CalculateProfit(cekArr)
		exp := &Calculation{
			MaxProfit: 5,
			BuyHour:   3,
			BuyPrice:  1,
			SellHour:  5,
			SellPrice: 6,
		}
		assert.Equal(t, exp, diff)
	})
}
