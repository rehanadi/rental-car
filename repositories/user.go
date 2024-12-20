package repositories

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"rental-car/helpers"
	"rental-car/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type UserRepository interface {
	Register(user *models.RegisterRequest) (*models.RegisterResponse, int, error)
	Login(user *models.LoginRequest) (interface{}, int, error)
	FindById(userId int) (*models.UserProfile, int, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *userRepository {
	return &userRepository{DB}
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
		if emailErr := helpers.SendRegistrationEmail(newUser.Email, newUser.Name); emailErr != nil {
			fmt.Printf("Failed to send registration email: %v\n", emailErr)
		}

		fmt.Println("Registration email sent")
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

func (r *userRepository) FindById(userId int) (*models.UserProfile, int, error) {
	// check if user exists
	var user models.User
	if err := r.DB.Where("user_id = ?", userId).
		First(&user).Error; err != nil {
		return nil, http.StatusNotFound, errors.New("user not found")
	}

	// return response
	response := models.UserProfile{
		Name:          user.Name,
		Email:         user.Email,
		DepositAmount: user.DepositAmount,
	}

	return &response, http.StatusOK, nil
}
