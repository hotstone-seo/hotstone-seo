import React, {PureComponent} from 'react';
import {Provider} from 'react-redux';
import {object} from 'prop-types';
import { ErrorBoundary, ErrorView } from 'tix-react-ui';
import {renderRoutes} from 'react-router-config';
import {BrowserRouter} from 'react-router-dom';

import routes from './routes';

class App extends PureComponent {
  static propTypes = {
    store: object.isRequired
  };

  render() {
    const { store } = this.props;

    return (
      <ErrorBoundary render={(error, info) => <ErrorView error={error} info={info} />}>
        <Provider store={store}>
          <BrowserRouter>
            {renderRoutes(routes)}
          </BrowserRouter>
        </Provider>
      </ErrorBoundary>
    );
  }
}

export default App;