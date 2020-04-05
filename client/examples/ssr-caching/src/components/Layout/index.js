import React, { PureComponent } from 'react';
import { string, bool, object } from 'prop-types';
import { renderRoutes } from "react-router-config";

import '../../styles/index.scss';

class Layout extends PureComponent {
  static propTypes = {
    isOnline: bool,
    lang: string,
    route: object
  };

  render() {
    const { route, isOnline = true } = this.props;
    const ds = { filter: 'grayscale(100%)' };
    const finalDS = isOnline ? { height: '100%' } : { height: '100%', ...ds };

    return (
      <div>
        {renderRoutes(route.routes)}
      </div>
    );
  }
}

export default Layout;