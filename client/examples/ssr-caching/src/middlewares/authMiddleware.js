export default function authMiddleware(req, res, next) {
  if ((!req.session || !req.session.user) && req.path !== '/auth/login') {
    return res.redirect('/auth/login');
  }

  return next();
}
