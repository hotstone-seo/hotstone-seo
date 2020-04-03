import React, { PureComponent } from 'react';
import { bool, func, node, string } from 'prop-types';
import ErrorView from './ErrorView';

class ErrorBoundary extends PureComponent {
  static propTypes = {
    children: node,
    debug: bool,
    errorMessage: string,
    render: func,
  };

  static defaultProps = {
    debug: false,
    errorMessage: '',
    render: null,
  };

  state = {
    hasError: false,
  };

  /* istanbul ignore next */
  componentDidCatch(error, info) {
    if (this.props.debug) {
      console.groupCollapsed(`Error occured!`);
      console.log('Error:', error);
      console.log('Info:', info);
      console.groupEnd();
    }

    this.setState({ hasError: true, error, info });
  }

  defaultRender = /* istanbul ignore next */ () => {
    const message = this.props.errorMessage || 'Sorry, something went wrong.';

    return <div>{message}</div>;
  };

  render() {
    const { error, info } = this.state;
    const { children, render } = this.props;
    const renderError = render || /* istanbul ignore next */ this.defaultRender;

    /* istanbul ignore else */
    return this.state.hasError ? /* istanbul ignore next */ renderError(error, info) : children;
  }
}

export { ErrorView };
export default ErrorBoundary;
