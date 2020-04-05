/**
 * THIS IS THE ENTRY POINT FOR THE CLIENT, JUST LIKE server.js IS THE ENTRY POINT FOR THE SERVER.
 */

import React from 'react';
import ReactDOM from 'react-dom';
import get from 'lodash/get';
import {AppContainer as HotEnabler} from 'react-hot-loader';
import i18n from './core/lang';
import App from './App';
import configureStore from './redux/configureStore';
import fetch from './core/fetch';

const lang = get(window.__INITIAL_STATE__, 'app.context.lang', false) || 'id';
const client = fetch({
  lang
});

const MOUNT_NODE = document.getElementById('app');
const store = configureStore(window.__INITIAL_STATE__, client);

const render = (TheApp) => {
  global.i18n = i18n(lang);

  ReactDOM.hydrate(
    <HotEnabler>
      <TheApp store={store}/>
    </HotEnabler>,
    MOUNT_NODE,
  );
};

render(App);

if (module.hot) {
  // Accept changes to this file for hot reloading.
  module.hot.accept('./client.js');
  module.hot.accept('./routes', () => {
    return true;
  });
  module.hot.accept('./App', () => {
    render(require('./App').default);
  });
}
