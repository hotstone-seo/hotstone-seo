import axios from 'axios';

// https://gist.github.com/paulsturgess/ebfae1d1ac1779f18487d3dee80d1258

class HotstoneAPI {
  constructor() {
    this.client = axios.create({
      baseURL: process.env.REACT_APP_API_URL
    });
  }

  getRules() {
    return this.client.get('/rules')
               .then(response => response.data)
               .catch(error => {
                 throw error;
               });
  }

  createRule(rule) {
    return this.client.post('/rules', rule)
               .then(response => response.data)
               .catch(error => {
                 throw error;
               })
  }

  updateRule(rule) {
    return this.client.put('/rules', rule)
               .then(response => response.data)
               .catch(error => {
                 throw error;
               });
  }

  deleteRule(id) {
    return this.client.delete(`/rules/${id}`)
               .then(response => response.data)
               .catch(error => {
                 throw error;
               });
  }

  getDataSources() {
    return this.client.get('/data_sources')
               .then(response => response.data)
               .catch(error => {
                 throw error;
               });
  }

  getDataSource(id) {
    return this.client.get(`/data_sources/${id}`)
               .then(response => response.data)
               .catch(error => {
                 throw error;
               });
  }

  getLocales() {
    return this.client.get('/locales')
               .then(response => response.data)
               .catch(error => {
                 throw error;
               });
  }

  postProviderMatchRule(path) {
    return this.axios.post(`provider/matchRule`, { path: path });
  }
}

export default new HotstoneAPI();
