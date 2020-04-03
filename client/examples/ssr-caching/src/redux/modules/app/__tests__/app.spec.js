import * as actions from '../actions';
import * as types from '../types';
import * as account from '../../../../models/account';

describe('app actions', () => {
  it('should create an action to setFlash', () => {
    const type = 'Finish docs';
    const text = 'Flash test';
    const expectedAction = {
      type: types.SET_FLASH,
      payload: { show: true, type, text }
    };

    expect(actions.setFlash(type, text)).toEqual(expectedAction);
  });

  it('should create an action to resetFlash', () => {
    const expectedAction = {
      type: types.RESET_FLASH
    };

    expect(actions.resetFlash()).toEqual(expectedAction);
  });

  it('should create an action to setPopup', () => {
    const payload = {
      show: true,
      header: 'test',
      footer: 'test',
      content: 'test'
    };
    const expectedAction = {
      type: types.SET_POPUP,
      payload
    };

    expect(actions.setPopup(payload)).toEqual(expectedAction);
  });

  it('should create an action to setPopup with empty data', () => {
    const payload = {
      show: true,
      header: '',
      footer: '',
      content: ''
    };
    const expectedAction = {
      type: types.SET_POPUP,
      payload
    };

    expect(
      actions.setPopup({
        show: true
      })
    ).toEqual(expectedAction);
  });

  it('should create an action to resetPopup', () => {
    const expectedAction = {
      type: types.RESET_POPUP
    };

    expect(actions.resetPopup()).toEqual(expectedAction);
  });

  it('should create an action to getAccount', async () => {
    account.getAccount();

    //noinspection JSAnnotator
    account.getAccount = jest.fn();
    account.getAccount
      .mockReturnValueOnce(Promise.resolve({ data: {} }))
      .mockReturnValueOnce(Promise.resolve({}))
      .mockReturnValueOnce(Promise.resolve(null))
      .mockReturnValueOnce(Promise.reject('test'));

    const expectedAction = {
      types: [types.LOAD_ACCOUNT, types.LOAD_ACCOUNT_SUCCESS, types.LOAD_ACCOUNT_FAIL],
      promise: () =>
        account.getAccount().then(res => {
          return res ? res.data || {} : {};
        })
    };

    const expectedActionResult = await expectedAction.promise();
    const result = await actions.getAccountAction().promise();

    expect(result).toMatchObject(expectedActionResult);

    // complete test branch
    account.getAccount();
    account.getAccount();
    account.getAccount();
    account.getAccount();
  });
});
