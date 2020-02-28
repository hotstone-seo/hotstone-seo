const path = require('path');

module.exports = {
  extends: ['react-app', 'airbnb'],
  rules: {
    'react/jsx-filename-extension': [1, { 'extensions': ['.js', '.jsx'] }],
    'no-underscore-dangle':  ['error', { 'allow': ['_sort', '_offset', '_limit', '_next_key', '_next'] }],
  },
  parser: "react-scripts/node_modules/babel-eslint",
  settings: {
    'import/resolver': {
      node: {
        paths: [path.resolve(__dirname, 'src')],
      }
    }
  }
}
