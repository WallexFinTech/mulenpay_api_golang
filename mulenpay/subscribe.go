package mulenpaygosdk

type Subscribe struct {
	Id          int    `json:"id"`
	Description string `json:"	description"`
	Status      int    `json:"	status"`
	Currency    string `json:"	currency"`
	Amount      string `json:"	amount"`
	StartDate   string `json:"	start_date"`
	NextPayDate string `json:"	next_pay_date"`
	Interval    string `json:"	interval"`
}

type AllSubscribtion struct {
	Items []*Subscribe `json:"items"`
}

type AllSubscribtionData struct {
	Success bool           `json:"success"`
	Items   *SubscribeList `json:"items"`
}

type SubscribeList struct {
	CurrentPage  int              `json:"current_page"`
	FirstPageUrl string           `json:"first_page_url"`
	From         int              `json:"from"`
	LastPage     int              `json:"last_page"`
	LastPageUrl  string           `json:"last_page_url"`
	Links        []*Link          `json:"links"`
	NextPageUrl  string           `json:"next_page_url"`
	Path         string           `json:"path"`
	PerPage      int              `json:"per_page"`
	PrevPageUrl  string           `json:"prev_page_url"`
	To           int              `json:"to"`
	Total        int              `json:"total"`
	Data         []*SubscribeItem `json:"data"`
}

type SubscribeItem struct {
	Id                  int    `json:"id"`
	Token               string `json:"token"`
	ShopId              int    `json:"shop_id"`
	Description         string `json:"description"`
	Amount              string `json:"amount"`
	Currency            string `json:"currency"`
	RequireConfirmation bool   `json:"require_confirmation"`
	StartDate           string `json:"start_date"`
	NextPayDate         string `json:"next_pay_date"`
	Interval            string `json:"interval"`
	Period              int    `json:"period"`
	MaxPeriods          int    `json:"max_periods"`
	Status              int    `json:"status"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
	Client              string `json:"client"`
	StartPaymentId      int    `json:"start_payment_id"`
}
