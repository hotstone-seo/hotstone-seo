import axios from 'axios';

const client = axios.create({
  baseURL: '/api/',
});

client.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.status !== 401) {
      return Promise.reject(error || 'Network error.Failed to connect API');
    }
    if (error.response) {
      const {
        data: { message },
      } = error.response;
      return Promise.reject(
        new Error(message || 'Unexpected error occured in server'),
      );
    }
    return Promise.reject(error);
  },
);

export default client;
