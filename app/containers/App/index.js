/**
 *
 * App
 *
 * This component is the skeleton around the actual pages, and should only
 * contain code that should be seen on all pages. (e.g. navigation bar)
 */

import React from 'react';
import { Helmet } from 'react-helmet';
// import styled from 'styled-components';
import { Switch, Route } from 'react-router-dom';
import { CssBaseline } from '@material-ui/core';

import HomePage from '../HomePage/Loadable';
import FeaturePage from '../FeaturePage/Loadable';
import NotFoundPage from '../NotFoundPage/Loadable';
import AllPosts from '../AllPosts/Loadable';
import LoginPage from '../LoginPage/Loadable';

import Layout1 from '../../components/layouts/layout1/Layout1';
import Layout2 from '../../components/layouts/layout2/Layout2';;

export default function App() {
  return (
    <React.Fragment>
      <CssBaseline />
      <main>
        <Layout1>
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
            <Route exact path="/dashboard" component={HomePage} />
            <Route path="/posts" component={AllPosts} />
            <Route path="/features" component={FeaturePage} />
            <Route path="" component={NotFoundPage} />
          </Switch>
        </Layout1>
      </main>
    </React.Fragment>
  );
}
