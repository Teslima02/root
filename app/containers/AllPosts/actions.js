/*
 *
 * AllPosts actions
 *
 */

import {
  OPEN_NEW_POST_DIALOG,
  CLOSE_NEW_POST_DIALOG,
  OPEN_EDIT_POST_DIALOG,
  CLOSE_EDIT_POST_DIALOG,
  SAVE_NEW_POST,
  GET_ALL_POSTS,
  GET_ALL_POSTS_SUCCESS,
  GET_ALL_POSTS_ERROR,
} from './constants';

export function openNewPostDialog() {
  return {
    type: OPEN_NEW_POST_DIALOG,
  };
}

export function closeNewPostDialog() {
  return {
    type: CLOSE_NEW_POST_DIALOG,
  };
}

export function openEditPostDialog() {
  return {
    type: OPEN_EDIT_POST_DIALOG,
  };
}

export function closeEditPostDialog() {
  return {
    type: CLOSE_EDIT_POST_DIALOG,
  };
}

export function allPosts() {
  return {
    type: GET_ALL_POSTS,
  };
}

export function allPostsSuccess(data) {
  return {
    type: GET_ALL_POSTS_SUCCESS,
    payload: data,
  };
}

export function allPostsError(data) {
  return {
    type: GET_ALL_POSTS_ERROR,
    payload: data,
  };
}

export function saveNewPost(data) {
  console.log(data, 'new post');
  return dispatch => {

    Promise.all([
      dispatch({
        type: SAVE_NEW_POST,
        payload: data,
      }),
    ]).then(() => dispatch(allPosts()));
  };
}
