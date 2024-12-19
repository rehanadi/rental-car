package models

type Payment struct {
	PaymentId     int     `json:"payment_id" gorm:"primaryKey"`
	UserId        int     `json:"user_id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
	Status        string  `json:"status"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type TopUpDepositRequest struct {
	UserId        int     `json:"user_id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
}

type TopUpDepositResponse struct {
	PaymentId     int     `json:"payment_id"`
	UserId        int     `json:"user_id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
	Status        string  `json:"status"`
	CreatedAt     string  `json:"created_at"`
	RedirectURL   string  `json:"redirect_url"`
}

type CreateInvoiceRequest struct {
	ExternalId         string  `json:"external_id"`
	Amount             float64 `json:"amount"`
	Description        string  `json:"description"`
	InvoiceDuration    int     `json:"invoice_duration"`
	GivenNames         string  `json:"given_names"`
	Email              string  `json:"email"`
	Currency           string  `json:"currency"`
	PaymentMethod      string  `json:"payment_method"`
	SuccessRedirectURL string  `json:"success_redirect_url"`
	FailureRedirectURL string  `json:"failure_redirect_url"`
}

type CreateInvoiceResponse struct {
	Id         string  `json:"id"`
	Status     string  `json:"status"`
	Amount     float64 `json:"amount"`
	InvoiceURL string  `json:"invoice_url"`
}
