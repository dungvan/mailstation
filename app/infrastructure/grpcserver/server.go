package grpcserver

import (
	"github.com/dungvan/mailstation/app/application"
	"github.com/dungvan/mailstation/app/domain/service"
	"github.com/dungvan/mailstation/app/infrastructure/repositories"
	pb "github.com/dungvan/mailstation/common/pb"

	"google.golang.org/grpc"
)

func RegisterServices(s *grpc.Server) {
	chartRepo := repositories.NewChartRepository()
	domainService := service.NewDashboardService(chartRepo)
	dashboardService := application.NewDashboardService(domainService)
	pb.RegisterDashboardServiceServer(s, dashboardService)
}
