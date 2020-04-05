import React, { Component } from 'react';
import { object, func } from 'prop-types';
import { connect } from 'react-redux';
import { Helmet } from 'react-helmet';
import { withRouter } from 'react-router-dom';
import PageError from '../../core/errors/PageError';

import { getAccountAction, setPopup, resetPopup } from '../../redux/actions';

import criticalCSS from './components/critical';
import './styles.scss';

class Landing extends Component {
  static propTypes = {
    account: object.isRequired,
    setPopup: func,
    resetPopup: func,
    query: object
  };

  state = {
    showLoading: true
  };

  componentDidMount() {
    const { order_id, order_hash } = this.props.query;
    const { loaded, loading } = this.props.account;

    if (order_id && order_hash && !loaded && !loading) {
      this.props.getAccountAction(order_id, order_hash);
    }

    setTimeout(() => {
      this.setState({
        showLoading: false
      });
    }, 3000);
  }

  // server side fetching
  static fetchData(store, query, params) {
    const { order_id, order_hash } = query;

    if (order_id && order_hash) {
      return store.dispatch(getAccountAction(order_id, order_hash));
    }

    return Promise.resolve(null);
  }

  handleShowLoading() {
    this.setState({
      showLoading: true
    }, () => {
      setTimeout(() => {
        this.setState({
          showLoading: false
        });
      }, 3000);
    });
  }

  render() {
    const {
      account,
      query
    } = this.props;
    const { showLoading } = this.state;

    const {
      error_msgs
    } = account;

    // accessing query params
    const { order_id, order_hash } = query;

    if (error_msgs) {
      // Error page handling
      throw new PageError(error_msgs);
    }

    return (
      <div className="content">
        {/** Critical CSS **/}
        <Helmet>
          <style>
            {criticalCSS}
          </style>
        </Helmet>

        <div className="card">
          <div className="col-12">
            <h1>React TIX Skeleton</h1>
          </div>
          <div className="row">
            <div className="col-4">
              <ul>
                <li>
                  <a className="side-menu">Menu 1</a>
                </li>
                <li>
                  <a className="side-menu">Menu 2</a>
                </li>
              </ul>
            </div>
            <div className="col-8">
              <p>Content Here</p>
              { showLoading ? 'Loading...' : <a className="btn" onClick={this.handleShowLoading.bind(this)}>Show Loading</a>}
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default withRouter(connect(
  state => ({
    account: state.app.account,
    query: state.app.context.query
  }),
  { getAccountAction, setPopup, resetPopup }
)(Landing));
