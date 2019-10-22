/**
 *
 * AllPosts
 *
 */

import React, { memo } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { createStructuredSelector } from 'reselect';
import { compose } from 'redux';

import { useInjectSaga } from 'utils/injectSaga';
import { useInjectReducer } from 'utils/injectReducer';
import { IconButton, Tooltip, Icon } from '@material-ui/core';
import { withStyles } from '@material-ui/styles';
import { Add, CloudUpload } from '@material-ui/icons';
import makeSelectAllPosts, {
  makeSelectOpenNewPostDialog,
  makeSelectPostDialog,
} from '../selectors';
import reducer from '../reducer';
import saga from '../saga';
import { openNewPostDialog, closeNewPostDialog } from '../actions';

const defaultToolbarStyles = {
  iconButton: {},
};

export function AddButton({
  classes,
  openNewPostDialog,
  closeNewPostDialog,
  postDialog,
}) {
  // console.log(postDialog, 'postDialog')
  // console.log(openNewPostDialog, 'openNewPostDialog')
  useInjectReducer({ key: 'allPosts', reducer });
  useInjectSaga({ key: 'allPosts', saga });

  return (
    <React.Fragment>
      <Tooltip title="Add New Post">
        <IconButton className={classes.iconButton} onClick={openNewPostDialog}>
          <Add className={classes.deleteIcon} />
        </IconButton>
      </Tooltip>

      {/* <AllPostsDialog
        postDialog={postDialog}
        closeNewPostDialog={closeNewPostDialog}
      /> */}
    </React.Fragment>
  );
}

AddButton.prototypes = {
  classes: PropTypes.object.isRequired,
  openNewPostDialog: PropTypes.func,
  closeNewPostDialog: PropTypes.func,
};

const mapStateToProps = createStructuredSelector({
  allPosts: makeSelectAllPosts(),
  postDialog: makeSelectPostDialog(),
  // openNewPostDialog: makeSelectOpenNewPostDialog(),
});

function mapDispatchToProps(dispatch) {
  return {
    openNewPostDialog: () => dispatch(openNewPostDialog()),
    // closeNewPostDialog: () => dispatch(closeNewPostDialog()),
    dispatch,
  };
}

const withConnect = connect(
  mapStateToProps,
  mapDispatchToProps,
);

export default compose(
  withStyles(defaultToolbarStyles, { name: 'CustomToolbar' }),
  withConnect,
  memo,
)(AddButton);
