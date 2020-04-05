import React, { Component } from 'react';
import { func, node, object } from 'prop-types';
import { canUseDOM } from 'fbjs/lib/ExecutionEnvironment'

import './styles.scss';

const isDev = process.env.NODE_ENV === 'development';
/* istanbul ignore next */
const defaultActionHandler = () => {
  if(canUseDOM) {
    window.location.reload()
  }
};

class ErrorView extends Component {
  static propTypes = {
    buttonActionHandler: func,
    buttonText: node,
    history: object,
    error: object,
    info: object
  };

  /* istanbul ignore next */
  render() {
    const { buttonActionHandler, buttonText, error, info } = this.props;
    const errorImage =
      error.image ||
      'https://res.cloudinary.com/debraf3cg/image/upload/v1519309968/temp/app-error_3x.png';
    const finalButtonActionHandler = error.link
      ? () => {
          if(canUseDOM) {
            location.href = error.link;
          }
        }
      : buttonActionHandler;

    return (
      <div className="error-boundary-wrapper">
        <div className="error-boundary-box">
          <img className="error-boundary-image" src={errorImage} />
          <p
            className="error-boundary-text"
            dangerouslySetInnerHTML={{ __html: i18n(error.message) }}
          />
          <a className="btn" onClick={finalButtonActionHandler || defaultActionHandler}>
            {error.buttonText || buttonText || i18n('errorBoundary.retry')}
          </a>
          {isDev && (
            /* eslint-disable */ <div className="error-boundary-stack"
            ><pre>{info.componentStack}</pre></div>
          ) /* eslint-enable */}
        </div>
      </div>
    );
  }
}

export default ErrorView;
