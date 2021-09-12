package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			PAN:     "5058 xxxx xxxx 9999",
			Active:  true,
			Balance: 600_00,
		},
		{
			PAN:     "5058 xxxx xxxx 6600",
			Active:  false,
			Balance: 10_00,
		},
		{
			PAN:     "5058 xxxx xxxx 1234",
			Active:  true,
			Balance: 100_00,
		},
		{
			PAN:     "5058 xxxx xxxx 8888",
			Active:  true,
			Balance: -5_00,
		},
	}
	paymetcard := PaymentSources(cards)
	fmt.Println(paymetcard)
	//Output:
	//[{card 5058 xxxx xxxx 9999 60000} {card 5058 xxxx xxxx 1234 10000}]

}