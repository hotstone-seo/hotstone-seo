import 'isomorphic-fetch';
import 'winston-logrotate';
import bodyParser from 'body-parser';
import cookieParser from 'cookie-parser';
import connectRedis from 'connect-redis';
import connectMysql from 'express-mysql-session';
import connectMemory from 'memorystore';
import express from 'express';
import path from 'path';
import winston from 'winston';
import session from 'express-session';
import config from 'config';

import initMiddleware from './middlewares/initMiddleware';
import authMiddleware from './middlewares/authMiddleware';
import deviceMiddleware from './middlewares/deviceMiddleware';
import AppServer from './AppServer';

const additionalWhiteList = process.env.WHITE_LIST;
const __PROD__ = process.env.NODE_ENV === 'production';
const sessionStore = {
  memory: () => {
    const MemoryStore = connectMemory(session);

    return new MemoryStore();
  },
  mysql: () => {
    const MysqlStore = connectMysql(session);

    return new MysqlStore(config.SessionMySQLConfig);
  },
  redis: () => {
    const RedisStore = connectRedis(session);

    return new RedisStore(config.SessionRedisConfig);
  }
};

const sessionOptions = {
  memory: null,
  mysql: config.SessionMySQLConfig,
  redis: config.SessionRedisConfig
};

// Configure Default Logger
winston.configure({
  transports: [
    new winston.transports.Console({
      handleExceptions: true,
      humanReadableUnhandledException: true,
      json: process.env.PRETTY_LOG === 'true',
    }),
  ],
  exitOnError: false,
});

if (__PROD__) {
  winston.remove(winston.transports.Console);
  winston.add(winston.transports.Rotate, {
    file: `${config.logDir}/${config.logFileName}`,
    colorize: false,
    timestamp: true,
    size: process.env.LOG_SIZE || '100m',
    keep: process.env.LOG_KEEP || 5,
    compress: false,
    json: true,
  });
} else {
  winston.add(winston.transports.File, {
    filename: `${config.logDir}${config.logFileName}`,
    handleExceptions: true,
    json: true,
  });
}

winston.info(`Starting server using [${config.SessionDriver}] as session store.`);

const server = express();

process.on('uncaughtException', error => {
  winston.error(error);
});
process.on('unhandledRejection', winston.error);

server.use(cookieParser(config.AppSecret));
server.use(bodyParser.urlencoded({ extended: true }));
server.use(bodyParser.json());
server.use(session({
  store: sessionStore[config.SessionDriver](),
  secret: config.secret,
  resave: false,
  saveUninitialized: false,
  name: 'PHPSESSID'
}));
server.use(initMiddleware);
server.use(deviceMiddleware);

if (__PROD__) {
  server.set('trust proxy', 1);
  server.disable('x-powered-by');
  // sessionConfig.cookie.secure = true
}

let whitelist = [
  'localhost',
  '127\\.0\\.0\\.1',
];

if (additionalWhiteList && additionalWhiteList !== '') {
  whitelist.concat(additionalWhiteList.split(','));
}

const corsOptions = {
  origin: (origin, callback) => {
    if (new RegExp(whitelist.join('|')).test(origin) || origin === undefined) {
      callback(null, true);
    } else {
      callback(new Error('Not allowed by CORS'));
    }
  },
  credentials: true,
};

server.use(express.static(path.resolve(__dirname, 'public')));

//ENABLE CORS
//app.use(cors(corsOptions));
server.use(AppServer);

server.use((err, req, res) => {
  winston.error(err);

  return res.send(!__PROD__ && err ? err : 'Ooopss.. We could not process request');
});

process.on('SIGINT', () => {
  winston.info('Received SIGINT exiting');
  process.exit();
});

//
// Hot Module Replacement
// -----------------------------------------------------------------------------
if (module.hot) {
  server.hot = module.hot;
  module.hot.accept('./AppServer', () => {
    server.use(require('./AppServer').default);
  });
  module.hot.accept();
} else {
  server.listen(config.port, err => {
    if (err) {
      winston.error(err);
    } else {
      winston.info(`Server is up and listening at ${config.host}:${config.port} env:${process.env.NODE_ENV}`);
    }

    if (process.env.CIRCLECI) {
      // Circle CI Build Success
      process.exit(0);
    }
  });
}

export default server;
