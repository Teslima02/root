/*
 *
 * LoginPage reducer
 *
 */
import produce from 'immer';
import { LOGIN, LOGIN_SUCCESS, LOGIN_ERROR } from './constants';

export const initialState = {
  loading: false,
  error: false,
  user: {},
  loginDetails: {},
};

/* eslint-disable default-case, no-param-reassign */
const loginPageReducer = (state = initialState, action) =>
  produce(state, (/* draft */) => {
    switch (action.type) {
      // case LOGIN: {
      //   return {
      //     loading: true,
      //     error: false,
      //     loginDetails: action.payload,
      //   };
      // }
      // case LOGIN_SUCCESS: {
      //   return {
      //     loading: false,
      //     error: false,
      //     user: action.payload,
      //   };
      // }
      // case LOGIN_ERROR: {
      //   return {
      //     loading: false,
      //     error: true,
      //   };
      // }
    }
  });

export default loginPageReducer;
