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
    .post("p/match", { path })
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function fetchTags(rule, locale, contentData) {
  const { rule_id, path_param } = rule;
  path_param["_locale"] = locale;
  path_param["_rule"] = rule_id;
  let queryParam = serialize(path_param);
  return client
    .get(`p/fetch-tags?${queryParam}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function serialize(obj) {
  let str = Object.keys(obj)
    .reduce(function (a, k) {
      a.push(k + "=" + encodeURIComponent(obj[k]));
      return a;
    }, [])
    .join("&");
  return str;
}

export { match, fetchTags };
