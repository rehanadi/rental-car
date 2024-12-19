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
