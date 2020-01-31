import client from '../client';

function fetchDatasources() {
  return client.get('/data_sources')
               .then(response => response.data)
               .catch(error => {
                 throw error;
               });
}

export { fetchDatasources };
