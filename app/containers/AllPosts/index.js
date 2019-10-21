/**
 *
 * AllPosts
 *
 */

import React, { memo, useEffect } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Helmet } from 'react-helmet';
import { createStructuredSelector } from 'reselect';
import { compose } from 'redux';

import { Grid, makeStyles } from '@material-ui/core';
import { useInjectSaga } from '../../utils/injectSaga';
import { useInjectReducer } from '../../utils/injectReducer';
import {
  makeSelectPostDialog,
  makeSelectGetAllPosts,
  makeSelectLoading,
  makeSelectError,
} from './selectors';
import reducer from './reducer';
import saga from './saga';
import { AllPostsList } from './components/AllPostsList';
import { AllPostsDialog } from './components/AllPostsDialog';
import { closeNewPostDialog, allPosts, saveNewPost } from './actions';

const useStyles = makeStyles(theme => ({
  textField: {
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
  },
}));

export function AllPosts({
  postDialog,
  closeNewPostDialog,
  dispatchAllPostsAction,
  getAllPosts,
  loading,
  error,
  dispatchNewPostAction,
}) {
  const classes = useStyles();
  useInjectReducer({ key: 'allPosts', reducer });
  useInjectSaga({ key: 'allPosts', saga });

  useEffect(() => {
    dispatchAllPostsAction();
  }, []);

  console.log(getAllPosts, 'getAllPosts')

  return (
    <React.Fragment>
      <Helmet>
        <title>AllPosts</title>
        <meta name="description" content="Description of AllPosts" />
      </Helmet>

      <Grid container spacing={3}>
        <Grid item xs={12} sm={12} md={12}>
          <AllPostsList
            loading={loading}
            error={error}
            getAllPosts={getAllPosts}
          />

          <AllPostsDialog
            postDialog={postDialog}
            closeNewPostDialog={closeNewPostDialog}
            dispatchNewPostAction={dispatchNewPostAction}
          />
        </Grid>
      </Grid>
    </React.Fragment>
  );
}

AllPosts.propTypes = {
  postDialog: PropTypes.object.isRequired,
  closeNewPostDialog: PropTypes.func.isRequired,
  getAllPosts: PropTypes.array.isRequired,
  dispatchAllPostsAction: PropTypes.func,
  loading: PropTypes.bool,
  error: PropTypes.oneOfType([PropTypes.object, PropTypes.bool]),
  dispatchNewPostAction: PropTypes.func,
};

const mapStateToProps = createStructuredSelector({
  postDialog: makeSelectPostDialog(),
  getAllPosts: makeSelectGetAllPosts(),
  loading: makeSelectLoading(),
  error: makeSelectError(),
});

function mapDispatchToProps(dispatch) {
  return {
    dispatchNewPostAction: evt => dispatch(saveNewPost(evt)),
    dispatchAllPostsAction: () => dispatch(allPosts()),
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
