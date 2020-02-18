import client from "./client";

function fetchMismatched(cfg = {}) {
  return client
    .get("metrics/mismatched", cfg)
    .then(response => response.data)
    .catch(error => {
      throw error;
    });
}

export { fetchMismatched };

const MetricAPI = {
  fetchMismatched: fetchMismatched
};

export default MetricAPI;
