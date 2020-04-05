import * as types from './types';

const initialState = {
  flash: {
    show: false,
    type: '',
    text: ''
  },
  popup: {
    show: false,
    header: '',
    footer: '',
    content: ''
  },
  account: {
    loading: false,
    loaded: false,
    data: {}
  },
  context: {
    query: {},
    params: {},
    lang: 'id',
    isWebView: false
  }
};

const ACTION_HANDLERS = {
  [types.SET_FLASH]: (state, action) => {
    return {
      ...state,
      flash: action.payload
    };
  },
  [types.RESET_FLASH]: state => {
    return {
      ...state,
      flash: {
        ...state.flash,
        show: false
      }
    };
  },
  [types.SET_POPUP]: (state, action) => {
    return {
      ...state,
      popup: action.payload
    };
  },
  [types.RESET_POPUP]: state => {
    return {
      ...state,
      popup: {
        ...state.popup,
        show: false
      }
    };
  },
  [types.LOAD_ACCOUNT]: state => {
    return {
      ...state,
      account: {
        ...state.account,
        loading: true,
        loaded: false
      }
    };
  },
  [types.LOAD_ACCOUNT_SUCCESS]: (state, action) => {
    const data = action.result || {};

    return {
      ...state,
      account: {
        data,
        loading: false,
        loaded: true
      }
    };
  },
  [types.LOAD_ACCOUNT_FAIL]: (state, action) => {
    return {
      ...state,
      account: {
        data: {},
        error: action.error,
        loading: false,
        loaded: false
      }
    };
  }
};

export default (state = initialState, action) => {
  const handler = ACTION_HANDLERS[action.type];

  return handler ? handler(state, action) : state;
};
