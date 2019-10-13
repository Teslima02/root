/**
 *
 * AllPosts
 *
 */

import React, { memo } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Helmet } from 'react-helmet';
import { FormattedMessage } from 'react-intl';
import { createStructuredSelector } from 'reselect';
import { compose } from 'redux';

import { useInjectSaga } from 'utils/injectSaga';
import { useInjectReducer } from 'utils/injectReducer';
import { Grid, Paper, TextField, makeStyles, Button } from '@material-ui/core';
import makeSelectAllPosts, { makeSelectPostDialog } from './selectors';
import reducer from './reducer';
import saga from './saga';
import messages from './messages';
import { AllPostsList } from './components/AllPostsList';
import { AllPostsDialog } from './components/AllPostsDialog';
import { closeNewPostDialog } from './actions';

const useStyles = makeStyles(theme => ({
  textField: {
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
  },
}));

export function AllPosts({ postDialog, closeNewPostDialog }) {
  const classes = useStyles();
  useInjectReducer({ key: 'allPosts', reducer });
  useInjectSaga({ key: 'allPosts', saga });

  return (
    <React.Fragment>
      <Helmet>
        <title>AllPosts</title>
        <meta name="description" content="Description of AllPosts" />
      </Helmet>

      <Grid container spacing={3}>
        <Grid item xs={12} sm={12} md={12}>
          <AllPostsList />
        </Grid>
      </Grid>

      <AllPostsDialog
        postDialog={postDialog}
        closeNewPostDialog={closeNewPostDialog}
      />
    </React.Fragment>
  );
}

AllPosts.propTypes = {
  // dispatch: PropTypes.func.isRequired,
  postDialog: PropTypes.object.isRequired,
  closeNewPostDialog: PropTypes.func.isRequired,
};

const mapStateToProps = createStructuredSelector({
  allPosts: makeSelectAllPosts(),
  postDialog: makeSelectPostDialog(),
});

function mapDispatchToProps(dispatch) {
  return {
    closeNewPostDialog: () => dispatch(closeNewPostDialog()),
    dispatch,
  };
}

const withConnect = connect(
  mapStateToProps,
  mapDispatchToProps,
);

export default compose(
  withConnect,
  memo,
)(AllPosts);
