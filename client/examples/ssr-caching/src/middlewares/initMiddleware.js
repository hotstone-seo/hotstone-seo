import newrelic from 'newrelic';
import config from 'config';

export default (req, res, next) => {
  newrelic.setTransactionName(req.path.substring(1));
  global.CONFIG = config.globals;
  res.locals = res.locals || {};

  return next();
};
