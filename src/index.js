import React from 'react';
import axios from 'axios';

export default function HotStone(host) {
  const client = axios.create({ baseURL: `${host}/tags` });
  return {
    tags(path) {
      const context = {
        async _realize() {
          try {
            const { data } = await client.get(path);
            return data;
          } catch (error) {
            console.log('error: ', error);
            return undefined;
          }
        },
        async get() {
          if (this.cache === undefined) {
            this.cache = await this._realize();
          }
          return this.cache;
        },
        async renderElement() {
          const tags = await this.get();
          return tags.map(({ type, props, children }) => (
            React.createElement(type, props, children)
          ));
        },
        // TODO add function for validation rule pattern against existing values
      };
      
      return context;
    }
  }
}
