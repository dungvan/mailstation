import { DashboardServiceClient } from './dashboard_grpc_web_pb';
import { ChartData, ChartDataRequest, ChartDataResponse } from './dashboard_pb';

const client = new DashboardServiceClient('http://localhost:50051');

export const fetchDashboardData = (): Promise<ChartData[]> => {
  return new Promise((resolve, reject) => {
    const request = new ChartDataRequest();
    client.getChartData(request, {}, (error, response: ChartDataResponse) => {
      if (error) {
        reject(error);
      } else {
        resolve(response.getDataList());
      }
    });
  });
};
