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

export function getAllPosts() {
  return {
    type: GET_ALL_POSTS,
  };
}

export function saveNewPost(data) {
  console.log(data, 'new post');
  return dispatch => {
    // type: SAVE_NEW_POST,
    // payload: data,

    Promise.all([
      dispatch({
        type: SAVE_NEW_POST,
        payload: data,
      }),
    ]).then(() => dispatch(getAllPosts()));
  };
}
