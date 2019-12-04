import axios from 'axios';

function HotStone(host) {
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
        }
        // TODO: Create function to render tags into React element
      };
      
      return context;
    }
  }
}

export default HotStone;
