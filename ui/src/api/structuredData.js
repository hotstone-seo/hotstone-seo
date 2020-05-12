import client from './client';

export function fetchStructuredDatas({ rule_id }) {
  return client.get('/structured-data', { params: { rule_id } })
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function getStructuredData(id) {
  return client.get(`/structured-data/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function createStructuredData(values) {
  return client.post('/structured-data', values)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function updateStructuredData(values) {
  return client.put(`/structured-data/${values.id}`, values)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function deleteStructuredData(id) {
  return client.delete(`/structured-data/${id}`)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function addFAQPage(values) {
  return client.post('/center/faq-page', values)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function updateFAQPage(values) {
  return client.put('/center/faq-page', values)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function addBreadcrumbList(values) {
  return client.post('/center/breadcrumb-list', values)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function updateBreadcrumbList(values) {
  return client.put('/center/breadcrumb-list', values)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function addLocalBusiness(values) {
  return client.post('/center/local-business', values)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export function updateLocalBusiness(values) {
  return client.put('/center/local-business', values)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}
