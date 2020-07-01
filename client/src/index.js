import queryString from "query-string"
import cachingFetch from "make-fetch-happen"

class HotStoneClient {
  constructor(hostURL, key, opts = {}) {
    this.bearer = `Bearer ${key}`
    this.baseURL = this.stripTrailingSlash(hostURL)
    this.fetch = cachingFetch.defaults(opts)
  }

  stripTrailingSlash(str) {
    return str.replace(/^(.+?)\/*?$/, "$1");
  }

  async match(path) {
    let rule = {};

    const param = { _path: path }
    const resp = await this.fetch(`${this.baseURL}/p/match?${queryString.stringify(param)}`, {
      headers: {
        Authorization: this.bearer
      }
    });
    if (!resp.ok && resp.status != 304) {
      return {}
    }

    try {
      rule = await resp.json();
      if (!rule.rule_id || rule.rule_id == 0) {
        return {}
      }
    } catch (e) {
      console.error("Failed to match/retrieve rule:", e.message);
      return {}
    }

    return rule;
  }

  async tags(rule, locale) {
    let tags = [];
    const { rule_id, path_param } = rule;

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
      return []
    }

    try {
      tags = await resp.json();
    } catch (e) {
      console.error("Failed to retrieve tags:", e.message);
      return []
    }

    return tags;
  }
}


export { HotStoneClient };
