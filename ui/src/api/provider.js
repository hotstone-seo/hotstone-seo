import axios from 'axios';
import qs from 'qs';
import Cookies from 'js-cookie';
import jwt from 'jsonwebtoken';
import _ from 'lodash';

export function getSimulationKey() {
  const token = Cookies.get('token');
  if (_.isEmpty(token)) {
    return '';
  }
  return jwt.decode(token).simulation_key;
}

export function getSimulationKeyPrefix() {
  const key = getSimulationKey();
  return key.slice(0, 7);
}

export default class ProviderAPI {
  constructor() {
    const key = getSimulationKey();
    this.client = axios.create({
      headers: { Authorization: `Bearer ${key}` },
    });
    this.client.interceptors.response.use(
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
  }

  async match(path) {
    return this.client
      .get(`p/match?_path=${path}`)
      .then((response) => response.data)
      .catch((error) => {
        throw error;
      });
  }

  async fetchTags(rule, locale) {
    const { rule_id, path_param } = rule;
    if (rule_id > 0) {
      path_param._locale = locale;
      path_param._rule = rule_id;
    }

    const queryParam = qs.stringify(path_param);
    return this.client
      .get(`p/fetch-tags?${queryParam}`)
      .then((response) => response.data)
      .catch((error) => {
        throw error;
      });
  }
}
