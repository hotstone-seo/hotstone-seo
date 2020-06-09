import client from './client';

export function fetchRules(cfg = {}) {
  return client
    .get('/rules', cfg)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function getRule(id) {
  return client
    .get(`/rules/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function createRule(rule) {
  return client
    .post('/rules', rule)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function updateRule(rule) {
  return client
    .put('/rules', rule)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function patchRule(id, fields) {
  return client
    .patch(`/rules/${id}`, fields)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function deleteRule(id) {
  return client
    .delete(`/rules/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

const RuleAPI = {
  fetch: fetchRules,
  get: getRule,
  create: createRule,
  update: updateRule,
  patch: patchRule,
  delete: deleteRule,
};

export default RuleAPI;
