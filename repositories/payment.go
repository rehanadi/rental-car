package repositories

import (
	"errors"
	"net/http"
	"os"
	"rental-car/helpers"
	"rental-car/models"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	FindAllPaymentMethods() ([]models.PaymentMethod, int, error)
	TopUpDeposit(payment *models.TopUpDepositRequest) (*models.TopUpDepositResponse, int, error)
	VerifyPayment(paymentId int, status string) (int, error)
}

type paymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) *paymentRepository {
	return &paymentRepository{DB}
}

func (r *paymentRepository) FindAllPaymentMethods() ([]models.PaymentMethod, int, error) {
	var paymentMethods []models.PaymentMethod
	if err := r.DB.Find(&paymentMethods).Error; err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return paymentMethods, http.StatusOK, nil
}

func (r *paymentRepository) TopUpDeposit(newPayment *models.TopUpDepositRequest) (*models.TopUpDepositResponse, int, error) {
	// check if user exists
	var user models.User
	if err := r.DB.Where("user_id = ?", newPayment.UserId).
		First(&user).Error; err != nil {
		return nil, http.StatusNotFound, errors.New("user not found")
	}

	// check if payment method exists
	var paymentMethod models.PaymentMethod
	if err := r.DB.Where("code = ?", newPayment.PaymentMethod).
		First(&paymentMethod).Error; err != nil {
		return nil, http.StatusNotFound, errors.New("payment method not found")
	}

	// insert into payments
	payment := models.Payment{
		UserId:        newPayment.UserId,
		Amount:        newPayment.Amount,
		PaymentMethod: newPayment.PaymentMethod,
		Status:        "pending",
		CreatedAt:     time.Now().Format(time.RFC3339),
		UpdatedAt:     time.Now().Format(time.RFC3339),
	}

	result := r.DB.Create(&payment)
	if result.Error != nil {
		return nil, http.StatusInternalServerError, result.Error
	}

	// create invoice
	baseURL := os.Getenv("BASE_URL")
	invoice := models.CreateInvoiceRequest{
		ExternalId:         strconv.Itoa(payment.PaymentId),
		Amount:             payment.Amount,
		Description:        "Top up deposit",
		InvoiceDuration:    86400,
		GivenNames:         user.Name,
		Email:              user.Email,
		Currency:           "IDR",
		PaymentMethod:      payment.PaymentMethod,
		SuccessRedirectURL: baseURL + "/payments/verify/" + strconv.Itoa(payment.PaymentId) + "?status=success",
		FailureRedirectURL: baseURL + "/payments/verify/" + strconv.Itoa(payment.PaymentId) + "?status=failed",
	}

	resInvoice, statusCode, err := helpers.CreateInvoice(invoice)
	if err != nil {
		return nil, statusCode, err
	}

	// return response
	response := models.TopUpDepositResponse{
		PaymentId:     payment.PaymentId,
		UserId:        payment.UserId,
		Amount:        payment.Amount,
		PaymentMethod: payment.PaymentMethod,
		Status:        payment.Status,
		CreatedAt:     payment.CreatedAt,
		RedirectURL:   resInvoice.InvoiceURL,
	}

	return &response, http.StatusCreated, nil
}

func (r *paymentRepository) VerifyPayment(paymentId int, status string) (int, error) {
	// check if payment status is pending
	var payment models.Payment
	if err := r.DB.Where("payment_id = ?", paymentId).First(&payment).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	if payment.Status != "pending" {
		return http.StatusBadRequest, errors.New("payment status is not pending")
	}

	// update payment status
	result := r.DB.Model(&models.Payment{}).Where("payment_id = ?", paymentId).Updates(map[string]interface{}{"status": status, "updated_at": gorm.Expr("NOW()")})
	if result.Error != nil {
		return http.StatusInternalServerError, result.Error
	}

	// add deposit to user balance if status is success
	if status == "success" {
		var payment models.Payment
		if err := r.DB.Where("payment_id = ?", paymentId).First(&payment).Error; err != nil {
			return http.StatusInternalServerError, err
		}

		var user models.User
		if err := r.DB.Where("user_id = ?", payment.UserId).First(&user).Error; err != nil {
			return http.StatusInternalServerError, err
		}

		user.DepositAmount += payment.Amount
		if err := r.DB.Save(&user).Error; err != nil {
			return http.StatusInternalServerError, err
		}
	}

	return http.StatusOK, nil
}
