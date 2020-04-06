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

function addMetaTag(metaTag) {
  return client.post('/center/metaTag', metaTag)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function updateMetaTag(metaTag) {
  return client.put('/center/metaTag', metaTag)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function addTitleTag(titleTag) {
  return client.post('/center/titleTag', titleTag)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function updateTitleTag(titleTag) {
  return client.put('/center/titleTag', titleTag)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function addCanonicalTag(canonicalTag) {
  return client.post('/center/canonicalTag', canonicalTag)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function updateCanonicalTag(canonicalTag) {
  return client.put('/center/canonicalTag', canonicalTag)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function addScriptTag(scriptTag) {
  return client.post('/center/scriptTag', scriptTag)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

function updateScriptTag(scriptTag) {
  return client.put('/center/scriptTag', scriptTag)
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
  updateMeta: updateMetaTag,
  addTitie: addTitleTag,
  updateTitle: updateTitleTag,
  addCanonical: addCanonicalTag,
  updateCanonical: updateCanonicalTag,
  addScript: addScriptTag,
  updateScript: updateScriptTag,
};

export default TagAPI;
