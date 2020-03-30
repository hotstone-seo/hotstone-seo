import client from './client';

export function fetchAuditTrails(cfg = {}) {
  return client
    .get('/audit-trail', cfg)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

const AuditTrailAPI = {
  fetch: fetchAuditTrails,
};

export default AuditTrailAPI;
