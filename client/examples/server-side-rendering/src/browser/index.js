import React from 'react';
import ReactDOM from 'react-dom';
import App from '../component/App';

ReactDOM.hydrate(<App data={window.__INITIAL_DATA__} />, document.getElementById('root'));
