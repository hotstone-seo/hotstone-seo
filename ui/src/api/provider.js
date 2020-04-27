import axios from "axios";

const client = axios.create({});

client.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.status !== 401) {
      return Promise.reject(error || "Network error.Failed to connect API");
    }
    if (error.response) {
      const {
        data: { message },
      } = error.response;
      return Promise.reject(
        new Error(message || "Unexpected error occured in server")
      );
    }
    return Promise.reject(error);
  }
);

function match(path) {
  return client
    .get(`p/match?_path=${path}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function fetchTags(rule, locale, contentData) {
  const { rule_id, path_param } = rule;
  if ( rule_id > 0) {
    path_param['_locale'] = locale;
    path_param['_rule'] = rule_id;
  }

  let qs = require('qs');

  let queryParam = qs.stringify(path_param);
  return client
    .get(`p/fetch-tags?${queryParam}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export { match, fetchTags };
