import { createSelector } from 'reselect';
import { initialState } from './reducer';

/**
 * Direct selector to the allPosts state domain
 */

const selectAllPostsDomain = state => state.allPosts || initialState;

/**
 * Other specific selectors
 */

/**
 * Default selector used by AllPosts
 */

const makeSelectAllPosts = () =>
  createSelector(
    selectAllPostsDomain,
    substate => substate,
  );

const makeSelectPostDialog = () =>
  createSelector(
    selectAllPostsDomain,
    postDialog => postDialog.postDialog,
  );

// const makeSelectOpenNewPostDialog = () =>
// createSelector(
//   selectAllPostsDomain,
//   openNewState => openNewState.postDialog,
// );

// const makeSelectCloseNewPostDialog = () =>
//   createSelector(
//     selectAllPostsDomain,
//     closeNewState => closeNewState.postDialog,
//   );

export default makeSelectAllPosts;
export {
  selectAllPostsDomain,
  makeSelectPostDialog,
  // makeSelectOpenNewPostDialog,
  // makeSelectCloseNewPostDialog,
};
