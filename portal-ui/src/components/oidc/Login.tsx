import React from 'react';
import { Button, Form, Input, Typography, Divider, notification } from 'antd';
import { useHistory } from 'react-router-dom';
import { keycloakIdpHints, keycloakIdpConfig } from '../../config/KeycloakIdpConfig';

const { Title } = Typography;

const Login: React.FC = () => {
  const history = useHistory();

  const handleUserPassLogin = async (values: { username: string; password: string }) => {
    const { username, password } = values;
    const params = new URLSearchParams();
    params.append('client_id', keycloakIdpConfig.clientId);
    params.append('client_secret', keycloakIdpConfig.clientSecret);
    params.append('grant_type', 'password');
    params.append('username', username);
    params.append('password', password);

    try {
      const response = await fetch(`${keycloakIdpConfig.issuer}/protocol/openid-connect/token`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: params.toString(),
      });

      if (!response.ok) {
        throw new Error('Failed to login');
      }

      const data = await response.json();
      console.log('Token:', data.access_token);
      // Handle the token (e.g., store it in local storage)
      notification.success({
        message: 'Login Successful',
        description: 'You have successfully logged in. Redirecting to your dashboard...',
      });
      history.push('/dashboard'); // Redirect to dashboard
    } catch (error) {
      notification.error({
        message: 'Login Failed',
        description: 'Unable to login. Please check your username and password and try again.',
      });
    }
  };

  const handleProviderHintLogin = (providerId: string) => {
    window.location.href = `${keycloakIdpConfig.issuer}/protocol/openid-connect/auth?client_id=${keycloakIdpConfig.clientId}&response_type=${keycloakIdpConfig.responseType || 'code'}&redirect_uri=${keycloakIdpConfig.redirectUri}&scope=${keycloakIdpConfig.scopes.join(' ')}&kc_idp_hint=${providerId}`;
  };

  const handleRegister = () => {
    window.location.href = `${keycloakIdpConfig.issuer}/protocol/openid-connect/registrations?client_id=${keycloakIdpConfig.clientId}&response_type=${keycloakIdpConfig.responseType || 'code'}&redirect_uri=${keycloakIdpConfig.redirectUri}&scope=${keycloakIdpConfig.scopes.join(' ')}`;
  };

  const handleForgotPassword = () => {
    window.location.href = `${keycloakIdpConfig.issuer}/protocol/openid-connect/auth?client_id=${keycloakIdpConfig.clientId}&response_type=${keycloakIdpConfig.responseType || 'code'}&redirect_uri=${keycloakIdpConfig.redirectUri}&scope=${keycloakIdpConfig.scopes.join(' ')}&kc_action=RESET_CREDENTIAL`;
  };

  return (
    <div style={{ minWidth: '350px', margin: '0 auto', textAlign: 'center' }}>
      <Title level={2}>Login</Title>
      <Form
        name="login"
        initialValues={{ remember: true }}
        onFinish={handleUserPassLogin}
      >
        <Form.Item
          name="username"
          rules={[{ required: true, message: 'Please input your username!' }]}
        >
          <Input placeholder="Username" />
        </Form.Item>
        <Form.Item
          name="password"
          rules={[{ required: true, message: 'Please input your password!' }]}
        >
          <Input.Password placeholder="Password" />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit" block>
            Login
          </Button>
        </Form.Item>
      </Form>
      <Button type="link" onClick={handleForgotPassword}>
        Forgot Password?
      </Button>
      <Button type="link" onClick={handleRegister}>
        Register
      </Button>
      <Divider>Or login with</Divider>
      {keycloakIdpHints.map(provider => (
        <Button
          key={provider.providerId}
          type="default"
          onClick={() => handleProviderHintLogin(provider.providerId)}
          block
          style={{ marginBottom: '10px' }}
        >
          {provider.displayName}
        </Button>
      ))}
    </div>
  );
};

export default Login;