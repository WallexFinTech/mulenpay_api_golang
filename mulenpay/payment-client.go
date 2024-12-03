package mulenpaygosdk

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"
)

type MulenpayClient struct {
	baseApi   string
	apiKey    string
	secretKey string
	timeout   time.Duration
}

func NewMPClient(baseApi, apiKey, secretKey string, timeout time.Duration) *MulenpayClient {
	return &MulenpayClient{
		baseApi:   baseApi,
		apiKey:    apiKey,
		timeout:   timeout,
		secretKey: secretKey}
}

func (c *MulenpayClient) CreatePayment(data *PaymentReq) (*PaymentRes, *ErrorResponse) {
	data.Sign = createSign(data, c.secretKey)
	jsonData, _ := json.Marshal(data)
	url := c.baseApi + "/payments"
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	res, err := c.sendRequest(req)
	if err != nil {
		status, _ := strconv.Atoi(res.Status)
		return nil, &ErrorResponse{Status: status, Error: err.Error()}
	}
	defer res.Body.Close()
	var output PaymentRes
	err = json.NewDecoder(res.Body).Decode(&output)
	if err != nil {
		status, _ := strconv.Atoi(res.Status)
		return nil, &ErrorResponse{Status: status, Error: err.Error()}
	}

	return &output, nil
}

func (c *MulenpayClient) GetAllPayments() (*AllPaymentsData, *ErrorResponse) {
	url := c.baseApi + "/payments"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	res, err := c.sendRequest(req)

	if err != nil {
		status, _ := strconv.Atoi(res.Status)
		return nil, &ErrorResponse{Status: status, Error: err.Error()}
	}
	defer res.Body.Close()
	var output AllPaymentsData
	err = json.NewDecoder(res.Body).Decode(&output)
	if err != nil {
		status, _ := strconv.Atoi(res.Status)
		return nil, &ErrorResponse{Status: status, Error: err.Error()}
	}

	return &output, nil

}

func (c *MulenpayClient) GetPayment(id string) (*PaymentData, *ErrorResponse) {
	url := c.baseApi + "/payments/" + id
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	res, err := c.sendRequest(req)
	if err != nil {
		status, _ := strconv.Atoi(res.Status)
		return nil, &ErrorResponse{Status: status, Error: err.Error()}
	}
	defer res.Body.Close()
	var output PaymentData
	err = json.NewDecoder(res.Body).Decode(&output)
	if err != nil {
		status, _ := strconv.Atoi(res.Status)
		return nil, &ErrorResponse{Status: status, Error: err.Error()}
	}

	return &output, nil
}
func (c *MulenpayClient) GetReceipt(id string) (*ReceiptData, *ErrorResponse) {
	url := c.baseApi + "/payments/" + id + "/receipt"

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	res, err := c.sendRequest(req)
	if err != nil {
		status, _ := strconv.Atoi(res.Status)
		return nil, &ErrorResponse{Status: status, Error: err.Error()}
	}
	defer res.Body.Close()
	var output ReceiptData
	err = json.NewDecoder(res.Body).Decode(&output)
	if err != nil {
		status, _ := strconv.Atoi(res.Status)
		return nil, &ErrorResponse{Status: status, Error: err.Error()}
	}

	return &output, nil
}

func (c *MulenpayClient) ConfirmPayment(id string) (*SuccessResponse, *ErrorResponse) {
	url := c.baseApi + "/payments/" + id + "/hold"
	req, _ := http.NewRequest(http.MethodPut, url, nil)
	res, err := c.sendRequest(req)
	if err != nil {
		status, _ := strconv.Atoi(res.Status)
		return nil, &ErrorResponse{Status: status, Error: err.Error()}
	}
	defer res.Body.Close()
	var output SuccessResponse
	err = json.NewDecoder(res.Body).Decode(&output)
	if err != nil {
		status, _ := strconv.Atoi(res.Status)
		return nil, &ErrorResponse{Status: status, Error: err.Error()}
	}

	return &output, nil
}
func (c *MulenpayClient) CancelPayment(id string) (*SuccessResponse, *ErrorResponse) {
	url := c.baseApi + "/payments/" + id + "/hold"
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	res, err := c.sendRequest(req)
	if err != nil {
		status, _ := strconv.Atoi(res.Status)
		return nil, &ErrorResponse{Status: status, Error: err.Error()}
	}
	defer res.Body.Close()
	var output SuccessResponse
	err = json.NewDecoder(res.Body).Decode(&output)
	if err != nil {
		status, _ := strconv.Atoi(res.Status)
		return nil, &ErrorResponse{Status: status, Error: err.Error()}
	}

	return &output, nil
}
func (c *MulenpayClient) RefundPayment(id string) (*SuccessResponse, *ErrorResponse) {
	url := c.baseApi + "/payments/" + id + "/refund"
	req, _ := http.NewRequest(http.MethodPut, url, nil)
	res, err := c.sendRequest(req)
	if err != nil {
		status, _ := strconv.Atoi(res.Status)
		return nil, &ErrorResponse{Status: status, Error: err.Error()}
	}
	defer res.Body.Close()
	var output SuccessResponse
	err = json.NewDecoder(res.Body).Decode(&output)
	if err != nil {
		status, _ := strconv.Atoi(res.Status)
		return nil, &ErrorResponse{Status: status, Error: err.Error()}
	}

	return &output, nil
}

func (c *MulenpayClient) GetAllSubscription() (*AllSubscribtionData, *ErrorResponse) {
	url := c.baseApi + "/subscribes"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	res, err := c.sendRequest(req)
	if err != nil {
		status, _ := strconv.Atoi(res.Status)
		return nil, &ErrorResponse{Status: status, Error: err.Error()}
	}
	defer res.Body.Close()
	var output AllSubscribtionData
	err = json.NewDecoder(res.Body).Decode(&output)
	if err != nil {
		status, _ := strconv.Atoi(res.Status)
		return nil, &ErrorResponse{Status: status, Error: err.Error()}
	}

	return &output, nil
}

func (c *MulenpayClient) CancelSubscription(id string) (*SuccessResponse, *ErrorResponse) {
	url := c.baseApi + "/subscribes/" + id
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	res, err := c.sendRequest(req)
	if err != nil {
		status, _ := strconv.Atoi(res.Status)
		return nil, &ErrorResponse{Status: status, Error: err.Error()}
	}
	defer res.Body.Close()
	var output SuccessResponse
	err = json.NewDecoder(res.Body).Decode(&output)
	if err != nil {
		status, _ := strconv.Atoi(res.Status)
		return nil, &ErrorResponse{Status: status, Error: err.Error()}
	}

	return &output, nil
}

func (c *MulenpayClient) sendRequest(req *http.Request) (*http.Response, error) {
	c.setReqHeaders(req)
	client := http.Client{
		Timeout: c.timeout,
	}
	return client.Do(req)
}

func (c *MulenpayClient) setReqHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
}

func createSign(data *PaymentReq, secretKey string) string {
	str := data.Currency + fmt.Sprintf("%v", data.Amount) + string(data.ShopId) + secretKey
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
