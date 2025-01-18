import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import { ConfigProvider, App } from 'antd';
import Home from './components/Home';
import Dashboard from './components/Dashboard';
import Callback from './components/oidc/Callback';

const ReactApp: React.FC = () => {
  return (
    <ConfigProvider theme={{ token: { colorPrimary: '#1890ff' } }}>
      <App>
        <Router>
          <Switch>
            <Route path="/dashboard" component={Dashboard} />
            <Route path="/callback" component={Callback} />
            <Route path="/" component={Home} />
          </Switch>
        </Router>
      </App>
    </ConfigProvider>
  );
};

export default ReactApp;
