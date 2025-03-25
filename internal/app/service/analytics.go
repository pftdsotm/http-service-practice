package service

import (
	"context"

	"github.com/pftdsotm/http-service-practice/internal/app/model"
	"github.com/pftdsotm/http-service-practice/internal/app/repository"
	"github.com/google/uuid"
)

type AnalyticsService struct {
	analyticsRepo *repository.AnalyticsRepository
}

func NewAnalyticsService(analyticsRepo *repository.AnalyticsRepository) *AnalyticsService {
	return &AnalyticsService{analyticsRepo: analyticsRepo}
}

func (s *AnalyticsService) RecordSale(ctx context.Context, warehouseID, productID uuid.UUID, soldQuantity int, totalSum float64) error {
	analytics := model.Analytics{
		WarehouseID:  warehouseID,
		ProductID:    productID,
		SoldQuantity: soldQuantity,
		TotalSum:     totalSum,
	}

	return s.analyticsRepo.RecordSale(ctx, analytics)
}

func (s *AnalyticsService) GetAnalyticsByWarehouse(ctx context.Context, warehouseID uuid.UUID) (model.AnalyticsData, error) {
	return s.analyticsRepo.GetAnalyticsByWarehouse(ctx, warehouseID)
}

func (s *AnalyticsService) GetTopWarehouses(ctx context.Context) ([]model.Analytics, error) {
	return s.analyticsRepo.GetTopWarehouses(ctx)
}