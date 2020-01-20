import React, { useState } from 'react';
import { Helmet } from 'react-helmet';
import { withRouter } from 'react-router-dom';
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

class HotStoneClient {
  constructor(hostURL) {
    this.apiCaller = axios.create({ baseURL: hostURL });
  }

  async match(path) {
    let rule = {};
    try {
      const { data } = await this.apiCaller.post('/provider/matchRule', { path });
      rule = data;
    } catch (e) {
      console.error('Failed to retrieve rule:', e.message);
    }
    return rule;
  }

  async tags(rule, locale, contentData={}) {
    let tags = [];
    try {
      const { data } = await this.apiCaller.post('/provider/tags', {
        rule_id: rule.rule_id,
        locale: locale,
        data: contentData
      });
      tags = data;
    } catch(e) {
      console.error('Failed to retrieve tags:', e.message);
    }
    return tags;
  }
}

// NOTE: What this comppnent should do:
// Manage meta tag which responds to path changes.
//
// There's a hook provided by react-router, useLocation, might be a good starting point
// but learn about React Hooks first!
//
// TODO: Find a way to detect path change
class HotStoneWrapper extends React.Component {
  constructor(props) {
    super(props);

    const { tags, client } = props;
    this.state = { tags };
    this.client = client;
    this.fetchTags = this.fetchTags.bind(this);
  }

  shouldComponentUpdate(nextProps, nextState) {
    if (this.props.location === nextProps.location) {
      return false;
    }
    return true;
  }
  
  componentDidUpdate(prevProps, prevState, snapshot) {
    if (this.props.location !== prevProps.location) {
      this.fetchTags(this.props.location);
    }
  }

  async fetchTags(path) {
    const rule = await this.client.match(path);
    const tags = await this.client.tags(rule);
    this.setState({ tags });
  }

  render() {
    const { tags } = this.state;
    const tagElements = tags.map(({ id, type, attributes, value }) => {
      attributes.key = id;
      return React.createElement(type, attributes, value);
    });
    return ( <Helmet>{tagElements}</Helmet> ); 
  }
}

const HotStone = withRouter(HotStoneWrapper);

export { HotStone, HotStoneClient };
