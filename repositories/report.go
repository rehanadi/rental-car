package repositories

import (
	"net/http"
	"rental-car/models"

	"gorm.io/gorm"
)

type ReportRepository interface {
	GetReportRentalDetail(userId int) (*[]models.ReportRentalDetail, int, error)
	GetReportRentalSummary(userId int) (*models.ReportRentalSummary, int, error)
}

type reportRepository struct {
	DB *gorm.DB
}

func NewReportRepository(DB *gorm.DB) *reportRepository {
	return &reportRepository{DB}
}

func (r *reportRepository) GetReportRentalDetail(userId int) (*[]models.ReportRentalDetail, int, error) {
	var reportRentalDetails []models.ReportRentalDetail
	if err := r.DB.Table("rentals").
		Select("rentals.rental_id, rentals.car_id, cars.name car_name, rentals.rental_cost, rentals.rental_days, rentals.subtotal_cost, rentals.tax_cost, rentals.total_cost, rentals.status, rentals.rented_at, rentals.expired_at, rentals.returned_at").
		Joins("JOIN cars ON rentals.car_id = cars.car_id").
		Where("rentals.user_id = ?", userId).
		Find(&reportRentalDetails).Error; err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &reportRentalDetails, http.StatusOK, nil
}

func (r *reportRepository) GetReportRentalSummary(userId int) (*models.ReportRentalSummary, int, error) {
	var reportRentalSummary models.ReportRentalSummary
	if err := r.DB.Table("rentals").
		Select("COUNT(DISTINCT rentals.car_id) as total_car, COUNT(rentals.rental_id) as total_rental, SUM(rentals.total_cost) as total_cost").
		Where("rentals.user_id = ?", userId).
		Scan(&reportRentalSummary).Error; err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &reportRentalSummary, http.StatusOK, nil
}
