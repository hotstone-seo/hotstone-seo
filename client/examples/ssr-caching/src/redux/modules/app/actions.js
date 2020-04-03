import * as types from './types';
import { getAccount } from '../../../models/account';

export const setFlash = (type, text) => ({
  type: types.SET_FLASH,
  payload: {
    show: true,
    type,
    text
  }
});

export const resetFlash = () => ({
  type: types.RESET_FLASH
});

export const setPopup = ({ header = '', footer = '', content = '' }) => ({
  type: types.SET_POPUP,
  payload: {
    show: true,
    header,
    footer,
    content
  }
});

export const resetPopup = () => ({
  type: types.RESET_POPUP
});

export const getAccountAction = () => {
  return {
    types: [types.LOAD_ACCOUNT, types.LOAD_ACCOUNT_SUCCESS, types.LOAD_ACCOUNT_FAIL],
    promise: (client) =>
      getAccount(client).then(res => {
        return res || {};
      })
  };
};
