package repository

import "github.com/dungvan/mailstation/app/domain/model"

type ChartRepository interface {
	GetChartData() ([]*model.ChartData, error)
}
