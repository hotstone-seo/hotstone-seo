import React, { useState } from 'react';
import { Helmet } from 'react-helmet';
import { useLocation } from 'react-router-dom';
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

// NOTE: What this comppnent should do:
// Manage meta tag which responds to path changes.
//
// There's a hook provided by react-router, useLocation, might be a good starting point
// but learn about React Hooks first!
//
// TODO: Find a way to detect path change
class HotStone extends React.Component {
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
    try {
      const rule = await this.client.match(path);
      const tags = await this.client.tags(rule);
      this.setState({ tags });
    } catch(error) {
      console.log(error);
    }
  }

  render() {
    const { tags } = this.state;
    const tagElements = tags.map(({ type, attributes, value }) => (
      React.createElement(type, attributes, value)
    ));
    return ( <Helmet>{tagElements}</Helmet> ); 
  }
}

export { HotStone, HotStoneClient };
