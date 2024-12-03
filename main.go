package main

import (
	"fmt"
	mp "mulenpaygosdk/mulenpay"
	"time"
)

func main() {
	mpClient := mp.NewMPClient(
		"https://dev.mulenpay.ru/api/v2",
		"MZrvqMuviYVmAIit6o7N1fPNwrfaNV0VYX12YOVX8ed2c6a7",
		"2a6cb201769c27858a806d393eb0759a334334529834b40776315029888d6e0d",
		30*time.Second,
	)
	data := &mp.PaymentReq{

		Currency:    "rub",
		Amount:      1000.50,
		Uuid:        "invoice_1232321121",
		ShopId:      21,
		Description: "Покупка булочек",
		Items: []*mp.PaymentItem{
			&mp.PaymentItem{
				Description:             "string",
				Quantity:                1,
				Price:                   1000.50,
				VatCode:                 0,
				PaymentSubject:          1,
				PaymentMode:             1,
				ProductCode:             "string",
				CountryOfOriginCode:     "string",
				CustomDeclarationNumber: "string",
				Excise:                  "string",
				MeasurementUnit:         0,
			},
		},
	}
	res0, err0 := mpClient.CreatePayment(data)
	fmt.Printf("%+v %+v\n\n", res0, err0)

	res, err := mpClient.GetAllSubscription()
	fmt.Printf("%+v %+v %+v \n\n", res, err, res.Items)

	res1, err1 := mpClient.GetAllPayments()
	fmt.Printf("%+v %+v %+v \n\n", res1, err1, res1.Payments)

	res2, err2 := mpClient.GetReceipt("94")
	fmt.Printf("%+v %+v %+v \n\n", res2, err2, res2.Items)
}
