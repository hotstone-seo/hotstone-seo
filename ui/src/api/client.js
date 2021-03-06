import axios from 'axios';
import { message } from 'antd';

const client = axios.create({
  baseURL: '/api/',
});

client.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response) {
      /* const {
        data: { message },
      } = error.response;
      return Promise.reject(
        new Error(message || "Unexpected error occured in server")
      ); */
      if (error.response.status === 504) {
        message.error('Network error.Failed to connect API');
        throw error; // return Promise.reject(error || 'Network error.Failed to connect API');
      } else if (error.response.status === 401 || error.response.status === 400) {
        message.error('Session expired. Please login first...');
        window.location.href = '/login';
      } else if (error.response.status === 404) {
        window.location.href = '/page-404';
      }
    }
    // return Promise.reject(error);
    throw error;
  },
);

export default client;
