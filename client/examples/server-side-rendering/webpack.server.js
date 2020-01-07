const path = require('path');
const merge = require('webpack-merge');
const webpackNodeExternals = require('webpack-node-externals');
const baseConfig = require('./webpack.base.js');

const config = {
  target: 'node',
  mode: 'development',
  entry: './src/index.js',
  externals: [webpackNodeExternals()],
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'server.js'
  }
}

module.exports = merge(baseConfig, config);
