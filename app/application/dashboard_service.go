package application

import (
	"context"

	"github.com/dungvan/mailstation/app/domain/service"
	pb "github.com/dungvan/mailstation/common/pb"
)

type DashboardService struct {
	pb.UnimplementedDashboardServiceServer
	domainService *service.DashboardService
}

func NewDashboardService(domainService *service.DashboardService) *DashboardService {
	return &DashboardService{domainService: domainService}
}

func (s *DashboardService) GetChartData(ctx context.Context, req *pb.ChartDataRequest) (*pb.ChartDataResponse, error) {
	// Use the domain service to get the chart data
	data, err := s.domainService.GetChartData()
	if err != nil {
		return nil, err
	}

	var chartData []*pb.ChartData
	for _, d := range data {
		chartData = append(chartData, &pb.ChartData{
			Name: d.Name,
			Uv:   d.Uv,
			Pv:   d.Pv,
		})
	}

	return &pb.ChartDataResponse{Data: chartData}, nil
}
