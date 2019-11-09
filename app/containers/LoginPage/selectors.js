import { createSelector } from 'reselect';
import { initialState } from './reducer';

/**
 * Direct selector to the loginPage state domain
 */

const selectLoginPageDomain = state => state.loginPage || initialState;

/**
 * Other specific selectors
 */

/**
 * Default selector used by LoginPage
 */

const makeSelectLoginPage = () =>
  createSelector(
    selectLoginPageDomain,
    substate => substate,
  );

// const makeSelectLoading = () =>
//   createSelector(
//     selectLoginPageDomain,
//     subState => subState.loading,
//   );

// const makeSelectError = () =>
//   createSelector(
//     selectLoginPageDomain,
//     subState => subState.error,
//   );

// const makeSelectLoginDetails = () =>
//   createSelector(
//     selectLoginPageDomain,
//     subState => subState.loginDetails,
//   );

export default makeSelectLoginPage;
export {
  selectLoginPageDomain,
  // makeSelectLoading,
  // makeSelectError,
  // makeSelectLoginDetails,
};
