import client from './client';

export function fetchMismatched(cfg = {}) {
  return client
    .get('metrics/mismatched', cfg)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function fetchListCountHitPerDay(cfg = {}) {
  return client
    .get('metrics/hit/range', cfg)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function fetchCountHit(cfg = {}) {
  return client
    .get('metrics/hit', cfg)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function fetchCountUniquePage(cfg = {}) {
  return client
    .get('metrics/unique-page', cfg)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function fetchClientKeyLastUsed(cfg = {}) {
  return client
    .get('metrics/client-key/last-used', cfg)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

const MetricAPI = {
  fetchMismatched,
  fetchCountHit,
  fetchListCountHitPerDay,
  fetchCountUniquePage,
  fetchClientKeyLastUsed,
};

export default MetricAPI;
