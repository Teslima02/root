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

import { useInjectSaga } from 'utils/injectSaga';
import { useInjectReducer } from 'utils/injectReducer';
import {
  Grid,
  Paper,
  TextField,
  makeStyles,
  FormControlLabel,
  Icon,
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
} from '@material-ui/core';
import makeSelectAllPosts, { makeSelectOpenNewPostDialog, makeSelectCloseNewPostDialog, makeSelectPostDialog } from '../selectors';
import reducer from '../reducer';
import saga from '../saga';
import { closeNewPostDialog } from '../actions';

const useStyles = makeStyles(theme => ({
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
    width: 200,
  },
  dense: {
    marginTop: 19,
  },
  menu: {
    width: 200,
  },
}));

export function AllPostsDialog({ postDialog, closeNewPostDialog }) {
  const classes = useStyles();
  useInjectReducer({ key: 'allPostsDialog', reducer });
  useInjectSaga({ key: 'allPostsDialog', saga });

  // useEffect(() => {
  //   console.log(postDialog, 'effect postDialog');
  // });

  // const closeComposeDialog = () => {
  //   postDialog.type === 'edit' ? closeNewPostDialog : closeNewPostDialog;
  // };

  return (
    <div>
      <Dialog
        {...postDialog.props}
        onClose={closeNewPostDialog}
        aria-labelledby="form-dialog-title"
      >
        <DialogTitle id="form-dialog-title">Subscribe</DialogTitle>
        <DialogContent>
          <DialogContentText>
            To subscribe to this website, please enter your email address here.
            We will send updates occasionally.
          </DialogContentText>
          <TextField
            autoFocus
            margin="dense"
            id="name"
            label="Email Address"
            type="email"
            fullWidth
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={closeNewPostDialog} color="primary">
            Cancel
          </Button>
          <Button onClick={closeNewPostDialog} color="primary">
            Subscribe
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}

AllPostsDialog.propTypes = {
  dispatch: PropTypes.func.isRequired,
  closeNewPostDialog: PropTypes.func,
};

const mapStateToProps = createStructuredSelector({
  allPosts: makeSelectAllPosts(),
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
)(AllPostsDialog);
