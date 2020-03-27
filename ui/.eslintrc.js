const path = require('path');

module.exports = {
  plugins: ["cypress"],
  extends: ['react-app', 'airbnb', "plugin:cypress/recommended"],
  rules: {
    'react/jsx-filename-extension': [1, { 'extensions': ['.js', '.jsx'] }],
    'no-underscore-dangle':  ['error', { 'allow': ['_sort', '_offset', '_limit', '_next_key', '_next'] }],
  },
  // parser: "react-scripts/node_modules/babel-eslint",
  settings: {
    'import/resolver': {
      node: {
        paths: [path.resolve(__dirname, 'src')],
      }
    }
  }
}
