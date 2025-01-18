package service

import (
	"github.com/dungvan/mailstation/app/domain/model"
	"github.com/dungvan/mailstation/app/domain/repository"
)

type DashboardService struct {
	chartRepo repository.ChartRepository
}

func NewDashboardService(chartRepo repository.ChartRepository) *DashboardService {
	return &DashboardService{chartRepo: chartRepo}
}

func (s *DashboardService) GetChartData() ([]*model.ChartData, error) {
	// Implement your business logic here
	data, err := s.chartRepo.GetChartData()
	if err != nil {
		return nil, err
	}

	// Example of additional business logic
	for _, d := range data {
		d.Uv = d.Uv * 2 // Example transformation
		d.Pv = d.Pv * 3 // Example transformation
	}

	return data, nil
}
