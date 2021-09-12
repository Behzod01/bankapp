package card

import "bank/pkg/bank/types"

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var paymentCards []types.PaymentSource
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			paymentCards = append(paymentCards, types.PaymentSource{Type: "card", Number: string(card.PAN), Balance: card.Balance})
		}
	}
	return paymentCards
}
