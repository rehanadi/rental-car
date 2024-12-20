package repositories

import (
	"rental-car/models"

	"github.com/stretchr/testify/mock"
)

type CategoryRepoMock struct {
	Mock mock.Mock
}

func (m *CategoryRepoMock) FindAll() *[]models.Category {
	res := m.Mock.Called()

	if res.Get(0) == nil {
		return nil
	}

	cars := res.Get(0).([]models.Category)
	return &cars
}
