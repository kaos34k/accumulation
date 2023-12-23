package test

import (
	"testing"
	"yourlambda/internal/domain"
	"yourlambda/internal/usecase"

	"github.com/stretchr/testify/assert"
)

func TestPointUsecase_GetPointByID(t *testing.T) {
	mockRepository := &MockPointRepository{}
	usecase := usecase.NewPointUsecase(mockRepository)

	// Implement your mock data and expectations
	mockRepository.On("GetPointByID", "1").Return(&domain.Point{ID: "1", Name: "Compra TC", }, nil)

	point, err := usecase.GetPointByID("1")

	assert.NoError(t, err)
	assert.NotNil(t, point)
	assert.Equal(t, "John Doe", point.Name)
}

func TestPointUsecase_CreatePoint(t *testing.T) {
	mockRepository := &MockPointRepository{}
	usecase := usecase.NewPointUsecase(mockRepository)

	// Implement your mock data and expectations
	mockRepository.On("CreatePoint", &domain.Point{ID: "1", Name: "Jane Doe"}).Return(nil)

	err := usecase.CreatePoint(&domain.Point{ID: "1", Name: "Jane Doe"})

	assert.NoError(t, err)
}
