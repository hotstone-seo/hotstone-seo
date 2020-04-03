import React from 'react';
import winston from 'winston';
import get from 'lodash/get';
import {renderToNodeStream} from 'react-dom/server';
import { Provider } from 'react-redux';
import { ErrorBoundary, ErrorView } from 'tix-react-ui';
import {matchRoutes, renderRoutes} from 'react-router-config';
import {StaticRouter} from 'react-router-dom';
import upperHTML from './core/html/start';
import lowerHTML from './core/html/end';
import i18n from './core/lang';
import { promiseState } from './core/utils';
import configureStore from './redux/configureStore';
import fetch from './core/fetch';

import routes from './routes';

export default (req, res, next) => {
  const lang = get(res, 'locals.data.app.context.lang', '');

  if (!req.session) {
    return next(new Error('Session not started'));
  }

  const client = fetch({
    lang,
    cookie: req.headers.cookie
  });

  const store = configureStore(res.locals.data, client);

  const hydrateOnClient = () => {
    const htmlStates = {
      ...res.locals.store.getState(),
      ...res.locals.data
    };

    return res.type('html').send(`${upperHTML()}${lowerHTML(htmlStates)}`);
  };

  const writeLowerHtml = (preloadedState) => {
    res.write(lowerHTML(preloadedState));

    return res.end();
  };

  const branch = matchRoutes(routes[0].routes, req.path);

  const promises = branch.map(({route}) => {
    let fetchData = route.component && route.component.fetchData;

    return fetchData instanceof Function ? fetchData(store, req.query, req.params) : Promise.resolve(null)
  });

  res.write(upperHTML());

  return Promise.all(promises).then(() => {
    global.i18n = i18n(lang);
    let context = {};
    const app = (
      <ErrorBoundary render={(error, info) => <ErrorView error={error} info={info} />}>
        <Provider store={store}>
          <StaticRouter location={req.url} context={context}>
            {renderRoutes(routes)}
          </StaticRouter>
        </Provider>
      </ErrorBoundary>
    );

    const preloadedState = store.getState();

    const html = renderToNodeStream(app);

    if(context.status === 404) {
      res.status(404);
    }

    if (context.status === 302) {
      return res.redirect(302, context.url);
    }

    html.pipe(res, {end: false});
    html.on('error', (e) => {
      winston.error(e);

      return writeLowerHtml(preloadedState);
    });
    html.on('end', () => {
      return writeLowerHtml(preloadedState);
    });
  }).catch(e => {
    winston.error(e);

    promises.forEach(p => {
      return promiseState(p) === 'pending' && p.cancel();
    });

    return hydrateOnClient();
  });
}