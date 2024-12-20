package controllers

import (
	"rental-car/repositories"

	"github.com/labstack/echo/v4"
)

type ReportController struct {
	ReportRepository repositories.ReportRepository
}

func NewReportController(rr repositories.ReportRepository) *ReportController {
	return &ReportController{rr}
}

func (rc *ReportController) GetReportRentalDetail(c echo.Context) error {
	userId := c.Get("user_id").(int)

	report, statusCode, err := rc.ReportRepository.GetReportRentalDetail(userId)
	if err != nil {
		return c.JSON(statusCode, map[string]string{"message": err.Error()})
	}

	return c.JSON(statusCode, map[string]any{"message": "success get report rental detail", "report": report})
}

func (rc *ReportController) GetReportRentalSummary(c echo.Context) error {
	userId := c.Get("user_id").(int)

	report, statusCode, err := rc.ReportRepository.GetReportRentalSummary(userId)
	if err != nil {
		return c.JSON(statusCode, map[string]string{"message": err.Error()})
	}

	return c.JSON(statusCode, map[string]any{"message": "success get report rental summary", "report": report})
}
