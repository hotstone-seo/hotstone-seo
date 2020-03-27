import client from './client';

function fetchMismatched(cfg = {}) {
  return client
    .get('metrics/mismatched', cfg)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function fetchListCountHitPerDay(cfg = {}) {
  return client
    .get('metrics/hit/range', cfg)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function fetchCountHit(cfg = {}) {
  return client
    .get('metrics/hit', cfg)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function fetchCountUniquePage(cfg = {}) {
  return client
    .get('metrics/unique-page', cfg)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export {
  fetchMismatched, fetchCountHit, fetchListCountHitPerDay, fetchCountUniquePage,
};

const MetricAPI = {
  fetchMismatched,
  fetchCountHit,
  fetchListCountHitPerDay,
  fetchCountUniquePage,
};

export default MetricAPI;
