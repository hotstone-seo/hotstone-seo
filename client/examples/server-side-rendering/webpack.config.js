var path = require('path');
var webpack = require('webpack');
var nodeExternals = require('webpack-node-externals');

const js = {
  test: /\.js$/,
  exclude: /node_modules/,
  use: {
    loader: "babel-loader",
    options: {
      presets: ["@babel/preset-env", "@babel/preset-react"],
    },
  },
};

var browserConfig = {
  entry: './src/browser/index.js',
  target: 'web',
  output: {
    path: path.resolve(__dirname, 'public'),
    filename: 'bundle.js',
    publicPath: '/'
  },
  module: {
    rules: [js],
  },
  plugins: [
    new webpack.DefinePlugin({
      __isBrowser__: 'true'
    })
  ]
}

var serverConfig = {
  entry: './src/server/index.js',
  target: 'node',
  externals: [nodeExternals()],
  output: {
    path: __dirname,
    filename: 'server.js',
    publicPath: '/'
  },
  module: {
    rules: [js],
  },
  plugins: [
    new webpack.DefinePlugin({
      __isBrowser__: 'false'
    })
  ],
  node: {
    __dirname: false,
  },
}

module.exports = [browserConfig, serverConfig]
