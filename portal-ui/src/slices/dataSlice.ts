import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { fetchDashboardData } from '../grpc/grpcClient';
import { RootState } from '../store';
import { ChartData } from '../grpc/dashboard_pb'; // Import ChartData

interface DataState {
  data: ChartData[];
  status: 'idle' | 'loading' | 'succeeded' | 'failed';
  error: string | null;
}

const initialState: DataState = {
  data: [],
  status: 'idle',
  error: null,
};

interface FetchDataArgType {
  userId: string;
}

export const fetchData = createAsyncThunk<ChartData[], FetchDataArgType>(
  'data/fetchData',
  async () => {
    const response = await fetchDashboardData();
    return response;
  }
);

const dataSlice = createSlice({
  name: 'data',
  initialState,
  reducers: {
    updateData(state, action: PayloadAction<ChartData[]>) {
      state.data = action.payload;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchData.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(fetchData.fulfilled, (state, action) => {
        state.status = 'succeeded';
        state.data = action.payload;
      })
      .addCase(fetchData.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.error.message || 'Failed to fetch data';
      });
  },
});

export const selectData = (state: RootState) => state.data.data;
export const selectDataStatus = (state: RootState) => state.data.status;
export const selectDataError = (state: RootState) => state.data.error;

export default dataSlice.reducer;
