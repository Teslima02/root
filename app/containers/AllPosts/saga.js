import {
  take,
  call,
  put,
  select,
  takeLatest,
  actionChannel,
} from 'redux-saga/effects';
import { GET_ALL_POSTS, SAVE_NEW_POST, UPDATE_POST } from './constants';
import request from '../../utils/request';
import {
  allPosts,
  allPostsSuccess,
  allPostsError,
  saveNewPostSuccess,
  saveNewPostError,
  updatePostSuccess,
  updatePostError,
} from './actions';
import { makeSelectNewPost, makeSelectUpdatePost } from './selectors';

// Individual exports for testing
export function* getAllPosts() {
  const requestURL = 'http://127.0.0.1:8081/articles';

  try {
    const allPostsRequ = yield call(request, requestURL);

    yield put(allPostsSuccess(allPostsRequ));
  } catch (err) {
    yield put(allPostsError(err));
  }
}

export function* saveNewPost() {
  const newPostData = yield select(makeSelectNewPost());

  const requestURL = 'http://127.0.0.1:8081/article';

  try {
    const newPostsRequ = yield call(request, requestURL, {
      method: 'POST',
      body: JSON.stringify(newPostData),
    });

    yield put(saveNewPostSuccess(newPostsRequ));
  } catch (err) {
    yield put(saveNewPostError(err));
  }
}

export function* updatePost() {
  const updatePostData = yield select(makeSelectUpdatePost());

  console.log(updatePostData, 'updatePostData');

  const requestURL = `http://127.0.0.1:8081/article/update/${
    updatePostData.id
  }`;

  console.log(requestURL, 'requestURL')

  try {
    const updatePostsRequ = yield call(request, requestURL, {
      method: 'PUT',
      body: JSON.stringify(updatePostData),
    });

    yield put(updatePostSuccess(updatePostsRequ));
  } catch (err) {
    yield put(updatePostError(err));
  }
}

export default function* posts() {
  yield takeLatest(GET_ALL_POSTS, getAllPosts);
  yield takeLatest(SAVE_NEW_POST, saveNewPost);
  yield takeLatest(UPDATE_POST, updatePost);
}
