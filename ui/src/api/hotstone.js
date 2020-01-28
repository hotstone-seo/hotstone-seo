import axios from "axios";

// https://gist.github.com/paulsturgess/ebfae1d1ac1779f18487d3dee80d1258

class HotstoneAPI {
  constructor() {
    let customAxios = axios.create({
      baseURL: process.env.REACT_APP_API_URL
    });

    this.axios = customAxios;
  }

  getLocales() {
    return this.axios.get("locales");
  }

  postProviderMatchRule(path) {
    return this.axios.post(`provider/matchRule`, { path: path });
  }
}

export default new HotstoneAPI();
