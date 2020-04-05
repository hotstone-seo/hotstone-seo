import Promise from 'bluebird'
import fetch, {Request, Headers, Response} from 'node-fetch'

fetch.Promise = Promise;
Response.Promise = Promise;

const localUrl = (url) => {
  if (url.startsWith('//')) {
    return `https:${url}`
  }

  if (url.startsWith('http')) {
    return url
  }

  return `${CONFIG.Hostname}`
};

const localFetch = (url, options) => fetch(localUrl(url), options);

export {localFetch as default, Request, Headers, Response}