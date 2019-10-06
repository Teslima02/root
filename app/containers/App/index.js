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

import FeaturePage from 'containers/FeaturePage/Loadable';
import NotFoundPage from 'containers/NotFoundPage/Loadable';
// import Header from 'components/Header';
// import Footer from 'components/Footer';

import { CssBaseline, Container } from '@material-ui/core';
import Layout1 from '../../components/layouts/layout1/Layout1';
// import Layout2 from '../../components/layouts/layout2/Layout2';
// import GlobalStyle from '../../global-styles';

export default function App() {
  return (
    <React.Fragment>
      <CssBaseline />
      <main>
        <Layout1>
          <Container maxWidth="md">
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
              {/* <Route exact path="/" component={HomePage} /> */}
              <Route path="/features" component={FeaturePage} />
              <Route path="" component={NotFoundPage} />
            </Switch>
          </Container>
        </Layout1>
      </main>
    </React.Fragment>
  );
}
