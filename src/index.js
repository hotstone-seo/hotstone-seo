import React from 'react';
import axios from 'axios';

/**
 * Returns a tag object after its string values have been interpolated with data
 *
 * @param {Object} tag - Tag object which contains values which may be a string template
 * @param {Object} data - The mapping for interpolating tag's string placeholders
 *
 * TODO: Provide implementation by iterating through all of tag's value
 */
function interpolate(tag, data) {
  // NOTE: This implementation is O(n2), please provide better implementation in the future
  for(let [key, value] of Object.entries(tag)) {
    tag[key] = value.replace(/{(\w+)}/g, function(match, capture) {
      return data[capture] || '';
    });
  }

  return tag;
}

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
        async _matchRule() {
          try {
            const { data } = await apiCaller.post('/provider/matchRule', { path });
            return data;
          } catch (error) {
            return undefined;
          }
        },
        async rule() {
          if (this.rule === undefined) {
            this.rule = await this._matchRule();
          }
          return this.rule;
        },
        async retrieveData() {
          try {
            const rule = await this.rule();
            const { data } = await apiCaller.post('/provider/retrieveData', { path, rule })
            return data;
          } catch(error) {
            return undefined;
          }
        },
        async tags() {
          try {
            const rule = await this.rule();
            const { data } = await apiCaller.get('/provider/tags', { params: { ruleID: rule.id } });
            return data;
          } catch (error) {
            return undefined;
          }
        },
        async articles() {
          try {
            const rule = await this.rule();
            const { data } = await apiCaller.get('provider/articles', { params: { ruleID: rule.id } });
            return data;
          } catch (error) {
            return undefined;
          }
        },
        render(template, data) {
          const tags = template.map((tag) => interpolate(tag, data))
          return tags.map(({ type, props, children }) => (
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
