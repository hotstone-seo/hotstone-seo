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

class HotStone {
  constructor(hostURL) {
    this.apiCaller = axios.create({ baseURL: hostURL });
  }

  async match(path) {
    const { data } = await this.apiCaller.post('/provider/matchRule', { path });
    return data;
  }

  async tags(rule, locale, contentData={}) {
    const { data } = await this.apiCaller.post('/provider/tags', {
      rule_id: rule.rule_id,
      locale: locale,
      data: contentData
    });
    return data;
  }
}

export default HotStone;
