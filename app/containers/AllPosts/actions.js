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
} from './constants';

export function openNewPostDialog() {
  return {
    type: OPEN_NEW_POST_DIALOG,
  };
}

export function closeNewPostDialog() {
  console.log('close welcome here');
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
