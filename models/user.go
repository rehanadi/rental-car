package models

type User struct {
	UserID        int     `json:"user_id" gorm:"primaryKey"`
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	Password      string  `json:"password"`
	DepositAmount float64 `json:"deposit_amount"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserProfile struct {
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	DepositAmount float64 `json:"deposit_amount"`
}
