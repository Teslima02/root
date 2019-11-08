/**
 *
 * App
 *
 * This component is the skeleton around the actual pages, and should only
 * contain code that should be seen on all pages. (e.g. navigation bar)
 */

import React, { memo, useEffect } from 'react';
import { Helmet } from 'react-helmet';
// import styled from 'styled-components';
import { Switch, Route } from 'react-router-dom';
import { CssBaseline } from '@material-ui/core';

import { connect } from 'react-redux';
import { compose } from 'redux';
import { createStructuredSelector } from 'reselect';

import HomePage from '../HomePage/Loadable';
import FeaturePage from '../FeaturePage/Loadable';
import NotFoundPage from '../NotFoundPage/Loadable';
import AllPosts from '../AllPosts/Loadable';
import LoginPage from '../LoginPage/Loadable';

import Layout1 from '../../components/layouts/layout1/Layout1';
import Layout2 from '../../components/layouts/layout2/Layout2';
import { makeSelectUserStatus } from './selectors';
import { getUserStatusAction } from './actions';

import PrivateRoute from './PrivateRoute';
import { AppContext } from './AppContext';

const App = ({ userStatus }) => {
  // useEffect(() => {
  //   getUserStatusAction();
  // }, []);

  console.log(userStatus, 'userStatus');

  return (
    <div>
      <AppContext.Provider value={false}>
        <React.Fragment>
          <CssBaseline />
          <main>
            <div>
              <Helmet
                titleTemplate="%s - React.js Boilerplate"
                defaultTitle="React.js Boilerplate"
              >
                <meta
                  name="description"
                  content="A React.js Boilerplate application"
                />
              </Helmet>

              <Switch>
                <Route exact path="/" component={LoginPage} />
                <Route exact path="/login" component={LoginPage} />
                <Layout1>
                  <PrivateRoute path="/dashboard" component={HomePage} />
                  {/* <Route exact path="/dashboard" component={HomePage} /> */}
                  {/* <Route path="/posts" component={AllPosts} /> */}
                  {/* <Route path="/features" component={FeaturePage} /> */}
                </Layout1>
                <Route path="" component={NotFoundPage} />
              </Switch>
            </div>
          </main>
        </React.Fragment>
      </AppContext.Provider>
    </div>
  );
};

const mapStateToProps = createStructuredSelector({
  userStatus: makeSelectUserStatus(),
});

function mapDispatchToProps(dispatch) {
  return {
    // getUserStatusAction: () => dispatch(getUserStatus()),
  };
}

const withConnect = connect(
  mapStateToProps,
  mapDispatchToProps,
);

export default compose(
  withConnect,
  memo,
)(App);
