package mulenpaygosdk

type PaymentReq struct {
	Currency    string         `json:"currency"`
	Amount      float64        `json:"amount"`
	Uuid        string         `json:"uuid"`
	ShopId      int            `json:"shopId"`
	Description string         `json:"description"`
	Subscribe   string         `json:"subscribe"`
	HoldTime    string         `json:"holdTime"`
	Sign        string         `json:"sign"`
	Items       []*PaymentItem `json:"items"`
}

type PaymentItem struct {
	Description             string  `json:"description"`
	Quantity                float64 `json:"quantity"`
	Price                   float64 `json:"price"`
	VatCode                 int     `json:"vat_code"`
	PaymentSubject          int     `json:"payment_subject"`
	PaymentMode             int     `json:"payment_mode"`
	ProductCode             string  `json:"product_code"`
	CountryOfOriginCode     string  `json:"country_of_origin_code"`
	CustomDeclarationNumber string  `json:"customs_declaration_number"`
	Excise                  string  `json:"excise"`
	MeasurementUnit         int     `json:"measurement_unit"`
}
type Payment struct {
	Id          int            `json:"id"`
	Uuid        string         `json:"uuid"`
	Amount      string         `json:"amount"`
	Currency    string         `json:"currency"`
	Description string         `json:"description"`
	Status      int            `json:"status"`
	Items       []*PaymentItem `json:"items"`
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

type ReceiptData struct {
	Status string        `json:"status"`
	Items  []ReceiptItem `json:"items"`
}

type ReceiptItem struct {
	Id                      int      `json:"id"`
	DaemonCode              string   `json:"daemon_code"`
	DeviceCode              string   `json:"device_code"`
	Warnings                []string `json:"warnings"`
	Error                   []string `json:"error"`
	EcrRegistrationNumber   string   `json:"ecr_registration_number"`
	FiscalDocumentAttribute int      `json:"fiscal_document_attribute"`
	FiscalDocumentNumber    int      `json:"fiscal_document_number"`
	FiscalReceiptNumber     int      `json:"fiscal_receipt_number"`
	FnNumber                string   `json:"fn_number"`
	FnsSite                 string   `json:"fns_site"`
	ReceiptDatetime         string   `json:"receipt_datetime"`
	ShiftNumber             string   `json:"shift_number"`
	Total                   float64  `json:"total"`
	OfdInn                  string   `json:"ofd_inn"`
	OfdReceiptUrl           string   `json:"ofd_receipt_url"`
	Status                  string   `json:"status"`
	Uuid                    string   `json:"uuid"`
	CreatedAt               string   `json:"created_at"`
	UpdatedAt               string   `json:"updated_at"`
}
