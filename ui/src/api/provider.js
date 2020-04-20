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
  return client
    .get(`p/rules/${rule.rule_id}/tags?locale=${locale}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export { match, fetchTags };
