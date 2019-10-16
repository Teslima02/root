import { take, call, put, select, takeLatest } from 'redux-saga/effects';
import { GET_ALL_POSTS } from './constants';
import request from '../../utils/request';
import { allPostsSuccess, allPostsError } from './actions';

// Individual exports for testing
export function* getAllPosts() {
  console.log('come here');
  const requestURL = 'http://127.0.0.1:8081/articles';

  try {
    const allPostsRequ = yield call(request, requestURL);

    // console.log(allPostsRequ, 'allPostsRequ success');

    yield put(allPostsSuccess(allPostsRequ));
  } catch (err) {
    yield put(allPostsError(err));
  }
}

export default function* posts() {
  yield takeLatest(GET_ALL_POSTS, getAllPosts);
}
