import client from './client';

function fetchTags({ locale, rule_id }) {
  return client.get('/tags', { params: { locale, rule_id } })
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function getTag(id) {
  return client.get(`/tags/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function createTag(tag) {
  return client.post('/tags', tag)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function updateTag(tag) {
  return client.put('/tags', tag)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function deleteTag(id) {
  return client.delete(`/tags/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function addMetaTag(rule_id, locale, name, content) {
  return client.post('/center/addMetaTag', {
    rule_id, locale, name, content,
  })
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function addTitleTag(rule_id, locale, title) {
  return client.post('/center/addTitleTag', { rule_id, locale, title })
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function addCanonicalTag(rule_id, locale, canonical) {
  return client.post('/center/addCanonicalTag', { rule_id, locale, canonical })
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function addScriptTag(rule_id, locale, type, source) {
  return client.post('/center/addScriptTag', {
    rule_id, locale, type, source,
  })
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export {
  fetchTags, getTag, createTag, updateTag, deleteTag,
  addMetaTag, addTitleTag, addCanonicalTag, addScriptTag,
};

const TagAPI = {
  fetch: fetchTags,
  get: getTag,
  create: createTag,
  update: updateTag,
  delete: deleteTag,
  addMeta: addMetaTag,
  addTitie: addTitleTag,
  addCanonical: addCanonicalTag,
  addScript: addScriptTag,
};

export default TagAPI;
