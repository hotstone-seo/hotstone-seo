import client from './client';

export async function fetchSettings(cfg = {}) {
  return client
    .get('/settings', cfg)
    .then((response) => response.data)
    .catch((error) => {
      throw error;
    });
}

export default { fetchSettings };
