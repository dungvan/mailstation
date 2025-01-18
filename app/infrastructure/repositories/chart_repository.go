package repositories

import (
	"github.com/dungvan/mailstation/app/domain/model"
	"github.com/dungvan/mailstation/app/domain/repository"
)

type ChartRepository struct{}

func NewChartRepository() repository.ChartRepository {
	return &ChartRepository{}
}

func (r *ChartRepository) GetChartData() ([]*model.ChartData, error) {
	// Implement your data fetching logic here
	data := []*model.ChartData{
		{Name: "Example", Uv: 100, Pv: 200},
	}
	return data, nil
}
