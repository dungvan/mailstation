import { combineReducers } from '@reduxjs/toolkit';
import dataReducer from './dataSlice';
import dashboardReducer from './dashboardSlice';

const rootReducer = combineReducers({
  data: dataReducer,
  dashboard: dashboardReducer,
  // ...other reducers
});

export type RootState = ReturnType<typeof rootReducer>;
export default rootReducer;
