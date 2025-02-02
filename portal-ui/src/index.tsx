import { createRoot } from 'react-dom/client';
import { Provider } from 'react-redux';
import store from './store';
import ReactApp from './App';
import './styles/index.css';
const container = document.getElementById('root');
if (container) {
  const root = createRoot(container);
  root.render(
    <Provider store={store}>
      <ReactApp />
    </Provider>
  );
} else {
  console.error('Root element not found');
}
