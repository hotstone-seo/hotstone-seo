import client from '../client';

function fetchRules() {
  return client.get('/rules')
               .then(response => response.data)
               .catch(error => {
                 throw error;
               });
}

export { fetchRules };
