import queryString from "query-string"
import cachingFetch from "make-fetch-happen"

class HotStoneClient {
  constructor(hostURL, key, opts = {}) {
    this.bearer = `Bearer ${key}`
    this.baseURL = hostURL
    this.fetch = cachingFetch.defaults(opts)
  }

  async match(path) {
    let rule = {};
    try {
      const param = {_path: path}
      const resp = await this.fetch(`${this.baseURL}/p/match?${queryString.stringify(param)}`, {
        headers: {
          Authorization: this.bearer
        }
      });
      if (!resp.ok && resp.status != 304) {
        throw new Error("HTTP status code: " + resp.status + " Resp: " + await resp.text())
      }
      rule = await resp.json();
      if (rule.rule_id == 0) {
        throw new Error("No matched rule")
      }
    } catch (e) {
      rule = {}
      console.error("Failed to retrieve rule:", e.message);
    }
    return rule;
  }

  async tags(rule, locale) {
    let tags = [];
    const { rule_id, path_param } = rule;
    try {
      const param = {
        _rule: rule_id,
        _locale: locale,
        ...path_param
      }
      const resp = await this.fetch(`${this.baseURL}/p/fetch-tags?${queryString.stringify(param)}`, {
        headers: {
          Authorization: this.bearer
        }
      });
      if (!resp.ok && resp.status != 304) {
        throw new Error("HTTP status code: " + resp.status + " Resp: " + await resp.text())
      }
      tags = await resp.json(); 
    } catch (e) {
      tags = [];
      console.error("Failed to retrieve tags:", e.message);
    }
    return tags;
  }
}


export { HotStoneClient };
