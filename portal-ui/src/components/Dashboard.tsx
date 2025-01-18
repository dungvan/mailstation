import React, { useState, useEffect } from 'react';
import { Layout, Menu, Typography } from 'antd';
import {
  DashboardOutlined,
  UserOutlined,
  SettingOutlined,
} from '@ant-design/icons';
import { useDispatch, useSelector } from 'react-redux';
import { fetchDashboardDataAsync } from '../slices/dashboardSlice';
import { RootState } from '../slices';
import { MenuProps } from 'antd';

const { Sider, Content } = Layout;
const { Title, Text } = Typography;

const Dashboard: React.FC = () => {
  const [selectedSection, setSelectedSection] = useState<string>('dashboard');
  const dispatch = useDispatch();
  const { loading, data, error } = useSelector((state: RootState) => state.dashboard);

  useEffect(() => {
    dispatch(fetchDashboardDataAsync());
  }, [dispatch]);

  const handleMenuClick: MenuProps['onClick'] = (e) => {
    setSelectedSection(e.key);
  };

  const menuItems = [
    {
      key: 'dashboard',
      icon: <DashboardOutlined />,
      label: 'Dashboard',
    },
    {
      key: 'profile',
      icon: <UserOutlined />,
      label: 'Profile',
    },
    {
      key: 'settings',
      icon: <SettingOutlined />,
      label: 'Settings',
    },
  ];

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sider width={200} className="site-layout-background">
        <Menu
          mode="inline"
          selectedKeys={[selectedSection]}
          onClick={handleMenuClick}
          style={{ height: '100%', borderRight: 0 }}
          items={menuItems}
        />
      </Sider>
      <Layout style={{ paddingLeft: '8px' }}>
        <Content className="site-layout-background" style={{ padding: 24, margin: 0, minHeight: 280 }}>
          {selectedSection === 'dashboard' && (
            <div>
              <Title level={2}>Dashboard</Title>
              <Text>Welcome to your dashboard.</Text>
              {loading && <p>Loading...</p>}
              {error && <p>Error: {error}</p>}
              {data && (
                <pre>{JSON.stringify(data, null, 2)}</pre>
              )}
            </div>
          )}
          {selectedSection === 'profile' && (
            <div>
              <Title level={2}>Profile</Title>
              <Text>Manage your profile information here.</Text>
            </div>
          )}
          {selectedSection === 'settings' && (
            <div>
              <Title level={2}>Settings</Title>
              <Text>Adjust your settings here.</Text>
            </div>
          )}
        </Content>
      </Layout>
    </Layout>
  );
};

export default Dashboard;
