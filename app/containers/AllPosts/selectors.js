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

const makeSelectGetAllPosts = () =>
  createSelector(
    selectAllPostsDomain,
    subState => subState.getAllPosts,
  );

const makeSelectLoading = () =>
  createSelector(
    selectAllPostsDomain,
    subState => subState.loading,
  );

const makeSelectError = () =>
  createSelector(
    selectAllPostsDomain,
    subState => subState.error,
  );

const makeSelectPostDialog = () =>
  createSelector(
    selectAllPostsDomain,
    subState => subState.postDialog,
  );

const makeSelectNewPost = () =>
  createSelector(
    selectAllPostsDomain,
    subState => subState.newPost,
  );

export default makeSelectAllPosts;
export {
  makeSelectNewPost,
  selectAllPostsDomain,
  makeSelectPostDialog,
  makeSelectGetAllPosts,
  makeSelectLoading,
  makeSelectError,
};
