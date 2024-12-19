package controllers

import (
	"net/http"
	"rental-car/models"
	"rental-car/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	PaymentRepository repositories.PaymentRepository
}

func NewPaymentController(pr repositories.PaymentRepository) *PaymentController {
	return &PaymentController{pr}
}

func (pc *PaymentController) TopUpDeposit(c echo.Context) error {
	var payment models.TopUpDepositRequest
	c.Bind(&payment)

	userId := c.Get("user_id").(int)
	payment.UserId = userId

	if payment.UserId == 0 || payment.Amount == 0 || payment.PaymentMethod == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "user_id, amount, and payment_method are required"})
	}

	response, statusCode, err := pc.PaymentRepository.TopUpDeposit(&payment)
	if err != nil {
		return c.JSON(statusCode, map[string]string{"message": err.Error()})
	}

	return c.JSON(statusCode, map[string]any{"message": "please open redirect url to complete top up payment", "payment": response})
}

func (pc *PaymentController) VerifyPayment(c echo.Context) error {
	paymentId, err := strconv.Atoi(c.Param("id"))

	if err != nil || paymentId == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "payment id must be a number"})
	}

	status := c.QueryParam("status")

	if status == "" || (status != "success" && status != "failed") {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "status must be either success or failed"})
	}

	statusCode, err := pc.PaymentRepository.VerifyPayment(paymentId, status)
	if err != nil {
		return c.JSON(statusCode, map[string]string{"message": err.Error()})
	}

	if status == "failed" {
		return c.JSON(statusCode, map[string]string{"message": "fail top up deposit"})
	}

	return c.JSON(statusCode, map[string]string{"message": "success top up deposit"})
}
