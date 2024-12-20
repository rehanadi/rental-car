package repositories

import (
	"rental-car/models"

	"github.com/stretchr/testify/mock"
)

type CarRepoMock struct {
	Mock mock.Mock
}

func (m *CarRepoMock) FindById(id int) *models.Car {
	res := m.Mock.Called(id)

	if res.Get(0) == nil {
		return nil
	}

	car := res.Get(0).(models.Car)
	return &car
}
