// ---------------------------------------
// Test Environment Setup
const React = require('react');
const enzyme = require('enzyme');
const fetch = require('axios');
const { Provider } = require('react-redux');
const { BrowserRouter } = require('react-router-dom');
const Adapter = require('enzyme-adapter-react-16');
const i18n = require('../src/core/lang').default;
const { JSDOM } = require('jsdom');
const config = require('config');

const jsdom = new JSDOM(
  '<!doctype html><html><body><div id="app"><!--APP--></div><div id="modal-root" ><!--MODAL--></div></body></html>'
);
const { window } = jsdom;

function copyProps(src, target) {
  const props = Object.getOwnPropertyNames(src)
    .filter(prop => typeof target[prop] === 'undefined')
    .reduce(
      (result, prop) => ({
        ...result,
        [prop]: Object.getOwnPropertyDescriptor(src, prop)
      }),
      {}
    );

  Object.defineProperties(target, props);
}

global.CONFIG = config.globals;
global.i18n = i18n;
global.window = window;
global.location = window.location;
global.document = window.document;
global.navigator = {
  userAgent: 'node.js'
};
copyProps(window, global);

const configureStore = require('../src/redux/configureStore');

process.env.NODE_ENV = 'test';



global.mockStore = configureStore({
  app: {
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
      query: {
        order_id: '39982221',
        order_hash: '2a71d91259eefafd4bea3465050d4fa7ecb4186e'
      },
      params: {},
      isWebView: false
    }
  }
}, fetch);

enzyme.configure({ adapter: new Adapter() });
global.enzyme = enzyme;

global.render = component => {
  return enzyme.render(
    <Provider store={mockStore}>
      <BrowserRouter>{component}</BrowserRouter>
    </Provider>
  );
};

global.shallow = component => {
  return enzyme.shallow(
    <Provider store={mockStore}>
      <BrowserRouter>{component}</BrowserRouter>
    </Provider>
  );
};

global.mount = component => {
  return enzyme.mount(
    <Provider store={mockStore}>
      <BrowserRouter>{component}</BrowserRouter>
    </Provider>
  );
};

if (process.env.NODE_ENV === 'test') {
  window.matchMedia =
    window.matchMedia ||
    function() {
      return {
        matches: false,
        addListener: function() {},
        removeListener: function() {}
      };
    };
}
