import React, { useState } from 'react';
import { Layout, Row, Col, Typography, Menu, Affix } from 'antd';
import Login from './oidc/Login';
import '../styles/Home.css'; // Import the CSS file
import { MenuProps } from 'antd';

const { Header, Content } = Layout;
const { Title, Text } = Typography;

const Home: React.FC = () => {
  const [selectedSection, setSelectedSection] = useState<string>(''); // Initialize with an empty string

  const handleMenuClick: MenuProps['onClick'] = (e) => {
    setSelectedSection(e.key);
  };

  const menuItems = [
    {
      key: 'contact',
      label: 'Contact',
    },
    {
      key: 'about',
      label: 'About',
    },
  ];

  return (
    <Layout style={{ minHeight: '100vh', overflowX: 'hidden' }}>
      <Affix offsetTop={0}>
        <Header className="header custom-header">
          <Menu
            mode="horizontal"
            onClick={handleMenuClick}
            selectedKeys={[selectedSection]}
            items={menuItems}
            className="menu"
          />
        </Header>
      </Affix>
      <Content className="main-content">
        <Row gutter={16} className="content-row">
          <Col
            xs={24}
            md={18}
            lg={19}
            className="left-column"
          >
            <div className="overlay" />
            {selectedSection === 'about' && (
              <div id="about" className="section-content">
                <Title level={2} className="section-title">About</Title>
                <Text className="section-text">
                  Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
                </Text>
              </div>
            )}
            {selectedSection === 'contact' && (
              <div id="contact" className="section-content">
                <Title level={2} className="section-title">Contact</Title>
                <Text className="section-text">
                  Email: contact@example.com
                  <br />
                  Phone: (123) 456-7890
                </Text>
              </div>
            )}
            {!selectedSection && (
              <div className="welcome-content">
                <Title level={1} className="welcome-title">Welcome to Our Portal</Title>
                <Text className="welcome-text">
                  Please select a menu item to learn more about us.
                </Text>
              </div>
            )}
          </Col>
          <Col
            xs={24}
            md={6}
            lg={5}
            className="right-column"
          >
            <Login />
          </Col>
        </Row>
      </Content>
    </Layout>
  );
};

export default Home;
