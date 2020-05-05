import React, { useState } from "react";
import { Helmet } from "react-helmet";
import { withRouter } from "react-router-dom";
import queryString from "query-string"
import cachingFetch from "make-fetch-happen"


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
  for (let [key, value] of Object.entries(tag)) {
    tag[key] = value.replace(/{(\w+)}/g, function(match, capture) {
      return data[capture] || "";
    });
  }

  return tag;
}

class HotStoneClient {
  constructor(hostURL, opts = {}) {
    this.baseURL = hostURL
    this.fetch = cachingFetch.defaults(opts)
  }

  async match(path) {
    let rule = {};
    try {
      const param = {_path: path}
      const resp = await this.fetch(`${this.baseURL}/p/match?${queryString.stringify(param)}`);
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
      const resp = await this.fetch(`${this.baseURL}/p/fetch-tags?${queryString.stringify(param)}`);
      tags = await resp.json(); 
    } catch (e) {
      tags = [];
      console.error("Failed to retrieve tags:", e.message);
    }
    return tags;
  }
}

const HotStoneContext = React.createContext([]);

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
    if (this.props.location.pathname === nextProps.location.pathname) {
      return false;
    }
    return true;
  }

  componentDidUpdate(prevProps, prevState, snapshot) {
    if (this.props.location.pathname !== prevProps.location.pathname) {
      this.fetchTags(this.props.location.pathname);
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
    return (
      <div>
        <Helmet>{tagElements}</Helmet>
        <HotStoneContext.Provider value={tags}>
          {this.props.children}
        </HotStoneContext.Provider>
      </div>
    );
  }
}

const HotStone = withRouter(HotStoneWrapper);

export { HotStone, HotStoneClient, HotStoneContext };
