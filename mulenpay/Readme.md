<h1>MulenPay API Client на языке Go</h1>
<h2>Описание</h2>
<p>Асинхронный Go-клиент для работы с API MulenPay. Подходит для управления платежами и подписками.</p>

<h2>импортировать<h2>
```bash
go get https://github.com/WallexFinTech/mulenpay_api_golang
```
<h2>Настройка и использование клиента </h2>
<h3>Для начала работы создайте client.</h3>
```bash
mpClient := mulenpaygosdk.NewMPClient(
		apiUrl,
		apiKey,
		sekretKey,
		requestTimeoutDuration,
	)
```

<h2>Пример использования.</h2>
<p>Создание платежа</p>
```bash
	 data := &mulenpaygosdk.PaymentReq{

 	Currency:    "3",
 	Amount:      "1000.50",
 	Uuid:        "invoice_123",
 	ShopId:      21,
 	Description: "Покупка булочек",
 }

res, err := mpClient.CreatePayment(data)
```
    <p>Где res результат и если ошибка то err </p>

<p>Получение списка платежей</p>
```bash
res, err := mpClient.GetAllSubscription()
```
<p>Получение платежа по ID</p>
```bash
res, err := mpClient.GetPayment(id)
```
Где id string

<p>Подтверждение платежа</p>
```bash
res, err := mpClient.ConfirmPayment(id)
```

<p>Получение чека по ID</p>
```bash
res, err := mpClient.GetReceipt(id)
```
Где id string

<p>Отмена платежа</p>
```bash
res, err := mpClient.CancelPayment(id)
```
<p>Возврат платежа</p>
```bash
res, err := mpClient.RefundPayment(id)
```
<h2>Работа с подписками<h2>
<p>Получение списка подписок</p>
```bash
	res, err := mpClient.GetAllPayments()
```
<p>Удаление подписки по ID</p>
```bash
res, err := mpClient.CancelSubscription(id)
```
<p>Требования</p>
<p>go 1.23</p>
