/*
 *
 * AllPosts reducer
 *
 */
import produce from 'immer';
import {
  OPEN_NEW_POST_DIALOG,
  CLOSE_NEW_POST_DIALOG,
  OPEN_EDIT_POST_DIALOG,
  CLOSE_EDIT_POST_DIALOG,
  GET_ALL_POSTS,
  GET_ALL_POSTS_SUCCESS,
  GET_ALL_POSTS_ERROR,
} from './constants';

export const initialState = {
  getAllPosts: [],
  loading: false,
  error: false,
  postDialog: {
    type: 'new',
    props: {
      open: false,
    },
    data: null,
  },
};

/* eslint-disable default-case, no-param-reassign */
const allPostsReducer = (state = initialState, action) =>
  produce(state, draft => {
    switch (action.type) {
      case GET_ALL_POSTS: {
        return {
          ...state,
          loading: true,
          error: false,
          getAllPosts: [],
        };
      }
      case GET_ALL_POSTS_SUCCESS: {
        return {
          ...state,
          loading: false,
          error: false,
          getAllPosts: action.payload,
          // draft: [(draft.getAllPosts = action.payload)],
        };
      }
      case GET_ALL_POSTS_ERROR: {
        return {
          ...state,
          loading: false,
          error: false,
        };
      }
      case OPEN_NEW_POST_DIALOG: {
        return {
          ...state,
          postDialog: {
            type: 'new',
            props: {
              open: true,
            },
            data: null,
          },
        };
      }
      case CLOSE_NEW_POST_DIALOG: {
        return {
          ...state,
          postDialog: {
            type: 'new',
            props: {
              open: false,
            },
            data: null,
          },
        };
      }
    }
  });

export default allPostsReducer;
