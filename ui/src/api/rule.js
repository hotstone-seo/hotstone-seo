import client from './client';

function fetchRules(cfg = {}) {
  return client
    .get('/rules', cfg)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function getRule(id) {
  return client
    .get(`/rules/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function createRule(rule) {
  return client
    .post('/rules', rule)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function updateRule(rule) {
  return client
    .put('/rules', rule)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function deleteRule(id) {
  return client
    .delete(`/rules/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export {
  fetchRules, getRule, createRule, updateRule, deleteRule,
};

const RuleAPI = {
  fetch: fetchRules,
  get: getRule,
  create: createRule,
  update: updateRule,
  delete: deleteRule
};

export default RuleAPI;
