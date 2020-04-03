/**
 * Configuration for WallabyJS test runner (COMMERCIAL, NOT FREE)
 * www.wallabyjs.com
 **/

module.exports = function(wallaby) {
  return {
    files: ['src/**/*.js?(x)', '!**/*.spec.js?(x)', 'tests/setup.js'],
    tests: ['src/**/*.spec.js?(x)'],

    env: {
      type: 'node',
      runner: 'node'
    },

    compilers: {
      '**/*.js?(x)': wallaby.compilers.babel()
    },

    testFramework: 'jest',
    debug: true,
    setup: wallaby => {
      const jestConfig = require('./package.json').jest;

      wallaby.testFramework.configure(jestConfig);
    }
  };
};
