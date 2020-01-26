import client from '../client';

export default function fetchRules() {
  return client.get('/rules')
               .then(response => response.data)
               .catch(error => {
                 throw error;
               });
}
