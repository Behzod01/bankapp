package card

import "bank/pkg/bank/types"

func Total(cards []types.Card) types.Money {
	sum := types.Money(0)
	for _, v := range cards {
		if v.Active && v.Balance > 0 {
			sum += v.Balance
		}
	}
	return sum
}