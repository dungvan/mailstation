import * as jspb from 'google-protobuf'



export class ChartDataRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChartDataRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ChartDataRequest): ChartDataRequest.AsObject;
  static serializeBinaryToWriter(message: ChartDataRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChartDataRequest;
  static deserializeBinaryFromReader(message: ChartDataRequest, reader: jspb.BinaryReader): ChartDataRequest;
}

export namespace ChartDataRequest {
  export type AsObject = {
  }
}

export class ChartDataResponse extends jspb.Message {
  getDataList(): Array<ChartData>;
  setDataList(value: Array<ChartData>): ChartDataResponse;
  clearDataList(): ChartDataResponse;
  addData(value?: ChartData, index?: number): ChartData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChartDataResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ChartDataResponse): ChartDataResponse.AsObject;
  static serializeBinaryToWriter(message: ChartDataResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChartDataResponse;
  static deserializeBinaryFromReader(message: ChartDataResponse, reader: jspb.BinaryReader): ChartDataResponse;
}

export namespace ChartDataResponse {
  export type AsObject = {
    dataList: Array<ChartData.AsObject>,
  }
}

export class ChartData extends jspb.Message {
  getName(): string;
  setName(value: string): ChartData;

  getUv(): number;
  setUv(value: number): ChartData;

  getPv(): number;
  setPv(value: number): ChartData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChartData.AsObject;
  static toObject(includeInstance: boolean, msg: ChartData): ChartData.AsObject;
  static serializeBinaryToWriter(message: ChartData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChartData;
  static deserializeBinaryFromReader(message: ChartData, reader: jspb.BinaryReader): ChartData;
}

export namespace ChartData {
  export type AsObject = {
    name: string,
    uv: number,
    pv: number,
  }
}

