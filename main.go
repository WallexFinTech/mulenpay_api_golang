package main

import (
	"fmt"
	mulenpaygosdk "mulenpaygosdk/mulenpay"
	"time"
)

func main() {
	mpClient := mulenpaygosdk.NewMPClient(
		"https://dev.mulenpay.ru/api/v2",
		"MZrvqMuviYVmAIit6o7N1fPNwrfaNV0VYX12YOVX8ed2c6a7",
		"2a6cb201769c27858a806d393eb0759a334334529834b40776315029888d6e0d",
		30*time.Second,
	)
	// data := &mulenpaygosdk.PaymentReq{

	// 	Currency:    "3",
	// 	Amount:      "1000.50",
	// 	Uuid:        "invoice_123",
	// 	ShopId:      21,
	// 	Description: "Покупка булочек",
	// }

	res, err := mpClient.GetAllSubscription()
	fmt.Printf("%+v %+v %+v ", res, err, res.Items)
	res1, err1 := mpClient.GetAllPayments()
	fmt.Printf("%+v %+v %+v ", res1, err1, res1.Payments)
}
