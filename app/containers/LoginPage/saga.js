import { call, put, select, takeLatest } from 'redux-saga/effects';
import request from '../../utils/request';

import { BaseUrl } from '../../components/BaseUrl';
import { makeSelectLoginDetails } from './selectors';
import { loginSuccessAction, loginErrorAction } from './actions';
import { LOGIN } from './constants';

export function* login() {
  const loginDetails = yield makeSelectLoginDetails();

  console.log(loginDetails, 'loginDetails');

  const requestURL = `${BaseUrl}/login`;

  try {
    const loginResponse = yield call(request, requestURL, {
      method: 'POST',
      body: JSON.stringify(loginDetails),
    });

    yield put(loginSuccessAction(loginResponse));
  } catch (err) {
    yield put(loginErrorAction(err));
  }
}

// Individual exports for testing
export default function* loginPageSaga() {
  yield takeLatest(LOGIN, login);
}
