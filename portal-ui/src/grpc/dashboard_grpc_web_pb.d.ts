import * as grpcWeb from 'grpc-web';

import * as dashboard_pb from './dashboard_pb'; // proto import: "dashboard.proto"


export class DashboardServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  getChartData(
    request: dashboard_pb.ChartDataRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: dashboard_pb.ChartDataResponse) => void
  ): grpcWeb.ClientReadableStream<dashboard_pb.ChartDataResponse>;

}

export class DashboardServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  getChartData(
    request: dashboard_pb.ChartDataRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<dashboard_pb.ChartDataResponse>;

}

