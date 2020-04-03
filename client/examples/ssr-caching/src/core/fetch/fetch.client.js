import fetch, {Request, Headers, Response} from 'isomorphic-fetch'

const localUrl = (url) => {
  if (url.startsWith('//')) {
    return `https:${url}`
  }

  if (url.startsWith('http')) {
    return url
  }

  if (url === '/graphql') {
    return (typeof window !== 'undefined' ? url : `${CONFIG.Hostname}/graphql`)
  } else {
    return url
  }
};

const localFetch = (url, options) => fetch(localUrl(url), options);

export {localFetch as default, Request, Headers, Response}