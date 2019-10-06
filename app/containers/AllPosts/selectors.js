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

export default makeSelectAllPosts;
export { selectAllPostsDomain };
