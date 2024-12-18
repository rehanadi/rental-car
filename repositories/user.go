package repositories

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"rental-car/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

var jwtSecret = []byte("secret")

type UserRepository interface {
	Register(user *models.RegisterRequest) (*models.RegisterResponse, int, error)
	Login(user *models.LoginRequest) (interface{}, int, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *userRepository {
	return &userRepository{DB}
}

func (r *userRepository) SendRegistrationEmail(email, name string) error {
	// Email configuration - consider using environment variables
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "noreply@yourcompany.com")
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Welcome to Our Rental Car Service")

	// Personalized email body
	body := fmt.Sprintf(`
		Dear %s,

		Welcome to our Rental Car Service! We're excited to have you on board.

		Your account has been successfully created. You can now log in and start exploring our services.

		If you have any questions, please don't hesitate to contact our support team.

		Best regards,
		The Rental Car Team
	`, name)

	mailer.SetBody("text/plain", body)

	// SMTP configuration - use environment variables for sensitive info
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	dialer := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		port,
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
	)

	// Send the email
	if err := dialer.DialAndSend(mailer); err != nil {
		return fmt.Errorf("failed to send registration email: %v", err)
	}

	return nil
}

func (r *userRepository) Register(user *models.RegisterRequest) (*models.RegisterResponse, int, error) {
	// check if email already exists
	var existingUser models.User
	r.DB.Where("email = ?", user.Email).First(&existingUser)
	if existingUser.Email != "" {
		return nil, http.StatusBadRequest, errors.New("email already exists")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return nil, http.StatusInternalServerError, errors.New("internal server error")
	}

	// insert into users
	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	err = r.DB.Create(&newUser).Error
	if err != nil {
		fmt.Println(err)
		return nil, http.StatusInternalServerError, errors.New("internal server error")
	}

	// Send registration email
	go func() {
		if emailErr := r.SendRegistrationEmail(newUser.Email, newUser.Name); emailErr != nil {
			fmt.Printf("Failed to send registration email: %v\n", emailErr)
		}
	}()

	// return response
	response := models.RegisterResponse{
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	return &response, http.StatusCreated, nil
}

func (r *userRepository) Login(user *models.LoginRequest) (interface{}, int, error) {
	// check if email exists
	var existingUser models.User
	r.DB.Where("email = ?", user.Email).First(&existingUser)
	if existingUser.Email == "" {
		return nil, http.StatusUnauthorized, errors.New("email not found")
	}

	// check password
	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return nil, http.StatusUnauthorized, errors.New("invalid password")
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": existingUser.UserID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return tokenString, http.StatusOK, nil
}
