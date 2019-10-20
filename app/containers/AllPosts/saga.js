import { take, call, put, select, takeLatest, actionChannel } from 'redux-saga/effects';
import { GET_ALL_POSTS, SAVE_NEW_POST } from './constants';
import request from '../../utils/request';
import {
  allPosts,
  allPostsSuccess,
  allPostsError,
  saveNewPostSuccess,
  saveNewPostError,
} from './actions';
import { makeSelectNewPost } from './selectors';

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
  const requestURL = 'http://127.0.0.1:8081/article';

  const newPostData = yield select(makeSelectNewPost());

  try {
    const allPostsRequ = yield call(request, requestURL, {
      method: 'POST',
      body: newPostData,
      // body: JSON.stringify(newPostData),
    });

    console.log(allPostsRequ, 'allPostsRequ');

    yield put(saveNewPostSuccess(allPostsRequ));
    yield actionChannel(allPosts);
  } catch (err) {
    yield put(saveNewPostError(err));
  }
}

export default function* posts() {
  yield takeLatest(GET_ALL_POSTS, getAllPosts);
  yield takeLatest(SAVE_NEW_POST, saveNewPost);
}
