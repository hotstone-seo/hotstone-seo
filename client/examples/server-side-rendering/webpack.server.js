const path = require('path');
const merge = require('webpack-merge');
const webpackNodeExternals = require('webpack-node-externals');
const baseConfig = require('./webpack.base.js');

const config = {
  target: 'node',
  mode: 'production',
  entry: './src/server/index.js',
  externals: [webpackNodeExternals()],
  output: {
    filename: 'server.js',
    path: __dirname,
  }
}

module.exports = merge(baseConfig, config);
