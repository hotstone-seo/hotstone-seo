const path = require('path');
const pkg = require('../package.json');

if (!process.env.NODE_ENV) {
  throw new Error('env NODE_ENV not defined');
}

let config = {
  env: process.env.NODE_ENV,
  host: process.env.HOST || 'localhost',
  port: process.env.PORT || 3000,
  secret: process.env.SECRET,
  version: pkg.version,
  Hostname: process.env.APPHOST || `http://${process.env.HOST}:${process.env.PORT}`,
  app: {
    title: 'Tiket',
    description: 'Tiket - Payment.',
    head: {
      defaultTitle: 'Tiket - Payment.',
      titleTemplate: 'Tiket - %s',
      meta: [
        {
          name: 'description',
          content: 'tiket.com - Payment.'
        },
        { name: 'theme-color', content: '#42b549' },
        { name: 'mobile-web-app-capable', content: 'yes' },
        { name: 'apple-mobile-web-app-title', content: 'tiket.com' },
        { name: 'apple-mobile-web-app-capable', content: 'yes' },
        { name: 'apple-mobile-web-app-status-bar-style', content: 'black-translucent' },
        { name: 'msvalidate.01', content: '3104E882307BB6900F56D266DC8320F6' },
        { name: 'msapplication-navbutton-color', content: '#42b549' },
        { name: 'format-detection', content: 'telephone=no' },
        { name: 'HandheldFriendly', content: 'true' },
        { name: 'robots', content: 'index,follow' },
        { property: 'og:site_name', content: 'tiket.com' },
        { property: 'og:image', content: 'https://tiket.com/images/logo-share.png' },
        { property: 'og:locale', content: 'en_US' },
        { property: 'og:title', content: 'Tiket - Payment.' },
        {
          property: 'og:description',
          content: 'tiket.com - Payment.'
        },
        { property: 'og:url', content: 'https://tiket.com/' }
      ],
      link: [
        { rel: 'canonical', href: 'https://tiket.com/' },
        {
          rel: 'search',
          href: 'https://tiket.com/opensearch.xml',
          title: 'tiket.com Search',
          type: 'application/opensearchdescription+xml'
        }
      ],
      scripts: []
    }
  },

  API_TIMEOUT: 3000,

  SessionDriver: process.env.SESSION_DRIVER || 'memory',
  SessionRedisConfig: {
    host: process.env.SESSION_REDIS_SERVER || '127.0.0.1',
    port: process.env.SESSION_REDIS_PORT || 6379,
    pass: process.env.SESSION_REDIS_SECRET || null,
    prefix: 'PHPREDIS_SESSION:'
  },

  SessionMySQLConfig: {
    host: process.env.SESSION_MYSQL_SERVER || '127.0.0.1',
    port: process.env.SESSION_MYSQL_PORT || 3306,
    user: process.env.SESSION_MYSQL_USER || 'root',
    password: process.env.SESSION_MYSQL_PASS || null,
    database: process.env.SESSION_MYSQL_DB || 'session',
    connectionLimit: 10,
    schema: {
      tableName: 'ci__sessions',
      columnNames: {
        session_id: 'session_id',
        expires: 'expires',
        data: 'user_data'
      }
    }
  },

    // ----------------------------------
  // Project Structure
  // ----------------------------------
  path_base: path.resolve(__dirname, '..'),
  dir_client: 'src',
  dir_build: 'build',
  dir_public: 'build/public/',
  dir_server: 'build',
  dir_test: 'tests',

  utils_paths: {
    base: () => path.resolve(config.path_base, arguments),
  }
};

config.logDir = process.env.LOGDIR || config.path_base;
config.logFileName = `${pkg.name}.log`;

config.compiler_public_path = `${config.Hostname}/`;

const base = (...args) => path.resolve(...[config.path_base, ...args]);

config.utils_paths = {
  base,
  client: base.bind(null, config.dir_client),
  build: base.bind(null, config.dir_build),
  public: base.bind(null, config.dir_public),
};

config.globals = {
  // Global Configuration Here
  // You can use it anywhere by accessing CONFIG.Hostname
  Hostname: config.Hostname
};

module.exports = config;