import React from 'react';
import axios from 'axios';

/**
 * A client that is used as the main building blocks for interacting with HotStone provider.
 * @typedef {Object} HotStoneClient
 */

/**
 * Create an instance of HotStoneClient.
 * 
 * @param {string} host - URL of the HotStone provider
 * @return {HotStoneClient} 
 */
function HotStone(host) {
  const apiCaller = axios.create({ baseURL: host });
  const client = {
    match: function(path) {
      const context = {
        async _getRule() {
          try {
            const { data } = await apiCaller.post('/provider/matchRule', { path });
            return data;
          } catch (error) {
            return undefined;
          }
        },
        async rule() {
          if (this.rule === undefined) {
            this.rule = await this._getRule();
          }
          return this.rule;
        },
        async tags() {
          try {
            const { data } = await apiCaller.get('/provider/tags', { params: { ruleID: this.rule().id } });
            return data;
          } catch (error) {
            return undefined;
          }
        },
        renderTags(data) {
          const tags = data.filter((element) => element.type === 'tags');
          return this._render(tags);
        },
        _render(data) {
          return data.map(({ type, props, children }) => (
            React.createElement(type, props, children)
          ));
        },
        // TODO add function for validation rule pattern against existing values
      };
      return context;
    }
  }
  return client;
}

export default HotStone;
