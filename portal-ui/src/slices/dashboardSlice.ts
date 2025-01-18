import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import { fetchDashboardData } from '../grpc/grpcClient';
import { ChartData } from '../grpc/dashboard_pb'; // Import ChartData

interface DashboardState {
  loading: boolean;
  data: ChartData[] | null;
  error: string | null;
}

const initialState: DashboardState = {
  loading: false,
  data: null,
  error: null,
};

export const fetchDashboardDataAsync = createAsyncThunk(
  'dashboard/fetchDashboardData',
  async () => {
    const response = await fetchDashboardData();
    return response;
  }
);

const dashboardSlice = createSlice({
  name: 'dashboard',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchDashboardDataAsync.pending, (state) => {
        state.loading = true;
      })
      .addCase(fetchDashboardDataAsync.fulfilled, (state, action) => {
        state.loading = false;
        state.data = action.payload;
      })
      .addCase(fetchDashboardDataAsync.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error.message || 'Failed to fetch data';
      });
  },
});

export default dashboardSlice.reducer;
