import axios from 'axios';

// https://gist.github.com/paulsturgess/ebfae1d1ac1779f18487d3dee80d1258
 
function handleRequest(req) {
  return req.then(response => response.data)
            .catch(error => {
              throw error;
            });
}

class HotstoneAPI {
  constructor() {
    this.client = axios.create({
      baseURL: process.env.REACT_APP_API_URL
    });
  }

  getRules() {
    return handleRequest(this.client.get('/rules'));
  }

  getRule(id) {
    return handleRequest(this.client.get(`/rules/${id}`));
  }

  createRule(rule) {
    return handleRequest(this.client.post('/rules', rule));
  }

  updateRule(rule) {
    return handleRequest(this.client.put('/rules', rule));
  }

  deleteRule(id) {
    return handleRequest(this.client.delete(`/rules/${id}`));
  }

  getDataSources() {
    return handleRequest(this.client.get('/data_sources'));
  }

  getDataSource(id) {
    return handleRequest(this.client.get(`/data_sources/${id}`));
  }

  getLocales() {
    return handleRequest(this.client.get('/locales'));
  }

  getLocale(id) {
    return handleRequest(this.client.get(`/locales/${id}`));
  }

  createLocale(locale) {
    return handleRequest(this.client.post('/locales', locale));
  }

  updateLocale(locale) {
    return handleRequest(this.client.put('/locales', locale));
  }

  deleteLocale(id) {
    return handleRequest(this.client.delete(`/locales/${id}`));
  }

  getTags() {
    return handleRequest(this.client.get('/tags'));
  }

  getTag(id) {
    return handleRequest(this.client.get(`/tags/${id}`));
  }

  createTag(tag) {
    return handleRequest(this.client.post('/tags', tag));
  }

  updateTag(tag) {
    return handleRequest(this.client.put('/tags', tag));
  }

  deleteTag(id) {
    return handleRequest(this.client.delete(`/tags/${id}`));
  }

  postProviderMatchRule(path) {
    return this.axios.post(`provider/matchRule`, { path: path });
  }
}

export default new HotstoneAPI();
