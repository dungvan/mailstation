import React, { useEffect } from 'react';
import { useHistory } from 'react-router-dom';
import { notification } from 'antd';
import { UserManager } from 'oidc-client';

const Callback: React.FC = () => {
  const history = useHistory();

  useEffect(() => {
    const userManager = new UserManager({});
    userManager.signinRedirectCallback().then(() => {
      notification.success({
        message: 'Login Successful',
        description: 'You have been successfully logged in.',
      });
      history.push('/dashboard');
    }).catch((error) => {
      notification.error({
        message: 'Login Error',
        description: `An error occurred: ${error.message}`,
      });
    });
  }, [history]);

  return (
    <div>
      <h1>Logging in...</h1>
    </div>
  );
};

export default Callback;
