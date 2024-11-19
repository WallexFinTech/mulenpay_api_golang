# MulenPay API Client на языке Go

## Описание
Асинхронный Go-клиент для работы с API MulenPay. Подходит для управления платежами и подписками.

## Установка
```bash
go get github.com/WallexFinTech/mulenpay_api_golang
```

## Настройка и использование клиента
Для начала работы создайте клиента:
```bash
mpClient := mulenpaygosdk.NewMPClient(
	apiUrl,
	apiKey,
	sekretKey,
	requestTimeoutDuration,
)
```
## Пример использования
Создание платежа
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
Где res — результат, а если возникла ошибка, то она будет в err.

## Получение списка платежей
```bash
res, err := mpClient.GetAllPayments()
```

Получение платежа по ID
```bash
res, err := mpClient.GetPayment(id)
```
Где id — строка.

Подтверждение платежа
```bash
res, err := mpClient.ConfirmPayment(id)
```

Отмена платежа
```bash
res, err := mpClient.CancelPayment(id)
```

Возврат платежа
```bash
res, err := mpClient.RefundPayment(id)
```

## Работа с подписками
Получение списка подписок
```bash
res, err := mpClient.GetAllSubscriptions()
```

Удаление подписки по ID
```bash
res, err := mpClient.CancelSubscription(id)
```

## Требования
Go 1.23+




