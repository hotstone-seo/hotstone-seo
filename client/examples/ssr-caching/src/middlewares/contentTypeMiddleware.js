export default function contentTypeMiddleware(req, res, next) {
  if (req.method === 'POST' && !req.is('application/json')) {
    res.statusCode = 415;

    return res.end();
  }

  return next();
}
