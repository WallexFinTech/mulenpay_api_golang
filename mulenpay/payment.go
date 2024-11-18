package mulenpaygosdk

type PaymentReq struct {
	Currency    string `json:"currency"`
	Amount      string `json:"amount"`
	Uuid        string `json:"uuid"`
	ShopId      int    `json:"shopId"`
	Description string `json:"description"`
	Subscribe   string `json:"subscribe"`
	HoldTime    string `json:"holdTime"`
	Sign        string `json:"sign"`
}

type Payment struct {
	Id          int    `json:"id"`
	Uuid        string `json:"uuid"`
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type AllPayments struct {
	Items []*Payment `json:"items"`
}

type PaymentRes struct {
	Success    bool   `json:"success"`
	PaymentUrl string `json:"paymentUrl"`
	Id         int    `json:"id"`
}

type PaymentData struct {
	Success bool     `json:"success"`
	Payment *Payment `json:"payment"`
}
type SuccessResponse struct {
	Success bool `json:"success"`
}
type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

type AllPaymentsData struct {
	Success  bool         `json:"success"`
	Payments *PaymentList `json:"payments"`
}

type PaymentList struct {
	CurrentPage  int        `json:"current_page"`
	FirstPageUrl string     `json:"first_page_url"`
	From         int        `json:"from"`
	LastPage     int        `json:"last_page"`
	LastPageUrl  string     `json:"last_page_url"`
	Links        []*Link    `json:"links"`
	NextPageUrl  string     `json:"next_page_url"`
	Path         string     `json:"path"`
	PerPage      int        `json:"per_page"`
	PrevPageUrl  string     `json:"prev_page_url"`
	To           int        `json:"to"`
	Total        int        `json:"total"`
	Data         []*Payment `json:"data"`
}
type Link struct {
	Url    string `json:"url"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
}
