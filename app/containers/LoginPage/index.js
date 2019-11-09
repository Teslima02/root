/**
 *
 * LoginPage
 *
 */

import React, { memo, useState } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Redirect } from 'react-router-dom';
import { Helmet } from 'react-helmet';
import { createStructuredSelector } from 'reselect';
import { compose } from 'redux';

import { useInjectSaga } from 'utils/injectSaga';
import { useInjectReducer } from 'utils/injectReducer';
import makeSelectLoginPage from './selectors';
import reducer from './reducer';
import saga from './saga';
import LoginForm from './components/LoginForm';

import { useAuth } from '../context/AppContext';
import { makeSelectUserToken } from '../App/selectors';

export function LoginPage({ tokens }) {
  useInjectReducer({ key: 'loginPage', reducer });
  useInjectSaga({ key: 'loginPage', saga });

  // const [isLoggedIn, setLoggedIn] = useState(false);

  const { setAuthTokens } = useAuth();

  if (tokens) {
    setAuthTokens(tokens);
    return <Redirect to="/dashboard" />;
  }

  return (
    <div>
      <Helmet>
        <title>LoginPage</title>
        <meta name="description" content="Description of LoginPage" />
      </Helmet>
      <LoginForm />
    </div>
  );
}

LoginPage.propTypes = {
  tokens: PropTypes.oneOfType([PropTypes.string, PropTypes.bool]),
};

const mapStateToProps = createStructuredSelector({
  loginPage: makeSelectLoginPage(),
  tokens: makeSelectUserToken(),
});

function mapDispatchToProps(dispatch) {
  // console.log(dispatch, 'dispatch');
  return {};
}

const withConnect = connect(
  mapStateToProps,
  mapDispatchToProps,
);

export default compose(
  withConnect,
  memo,
)(LoginPage);
