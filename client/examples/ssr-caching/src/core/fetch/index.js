import URL from 'url';
import qs from 'querystring';

import fetchServer from './fetch.server';
import fetchClient from './fetch.client';

let winston;
let fetch = fetchServer;
const isClient = !!process.env.BROWSER;

if (isClient) {
  fetch = fetchClient;
} else {
  winston = require('winston');
}

export const normalizeUrl = uri => {
  const url = URL.parse(uri);

  if (typeof url === 'object') {
    if (typeof url.format === 'function') {
      return url.format();
    }

    return '';
  }

  return url;
};

export const formatGetURL = (url, content) => {
  return `${url}${content ? `?${qs.stringify(content)}` : ''}`;
};

const request = (context) => (url, options) => {
  let start;

  if (!isClient) {
    start = process.hrtime();
  }

  const URI = normalizeUrl(url);
  const isGetMethod = options.method === 'GET';
  const isHeadMethod = options.method === 'HEAD';

  options.headers = options.headers || {};
  if (!isClient) {
    options.headers.cookie = context.cookie;
  }

  const body = isGetMethod || isHeadMethod ? null : qs.stringify(options.data);

  const finalURL =
    isGetMethod || options.queryURL ? formatGetURL(URI, options.queryURL || options.data) : URI;

  const finalOptions = {
    timeout: 30 * 1000,
    body,
    redirect: 'manual',
    credentials: 'include',
    ...options,
    headers: {
      TIXAPI: 1,
      source: 'desktop',
      'User-Agent': context.userAgent || 'tix-payment',
      'Content-Type': 'application/x-www-form-urlencoded',
      ...(isClient ? {} : {
        Origin: CONFIG.Hostname,
      }),
      ...options.headers,
    },
  };

  return fetch(finalURL, finalOptions)
    .then(res => {
      if (!isClient) {
        const end = process.hrtime(start) ;

        winston.info({
          executionTime: `${end[0]}s ${end[1]/1000000}ms`,
          url,
          options: finalOptions,
          ok: res.ok,
          status: res.status,
          statusText: res.statusText
        });
      }

      if (res.ok) {
        return options.rawResponse ? res : res.json()
          .catch(e => {
            winston.error(e);

            return res;
          });
      } else {
        return new Error(`[${res.status}] ${res.statusText}`);
      }
    }).catch(e => {
      if (!isClient) {
        winston.error(e);
      }

      throw e;
    });
};

export default request;
