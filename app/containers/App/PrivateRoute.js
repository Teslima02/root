import React from 'react';
import { Route, Redirect } from 'react-router-dom';
import { useAuth } from './AppContext';

function PrivateRoute({ component: Component, ...rest }) {
  const isAuthenticated = useAuth();
  // const { setAuthTokens } = useAuth();
  // console.log(setAuthTokens, 'setAuthTokens')
  return (
    <Route
      {...rest}
      render={props =>
        isAuthenticated ? <Component {...props} /> : <Redirect to="/" />
      }
    />
  );
}

export default PrivateRoute;
