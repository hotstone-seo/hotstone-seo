import { MOBILE_DEVICE, MOBILE_UA } from '../core/constants'

export default function deviceMiddleware(req, res, next) {
  const deviceType = req.query.device_type;
  const deviceUA = req.headers['user-agent'];

  res.locals.data = {
    app: {
      flash: {
        show: false,
        type: '',
        text: ''
      },
      popup: {
        show: false,
        header: '',
        footer: '',
        content: ''
      },
      account: {
        loading: false,
        loaded: false,
        data: {}
      },
      context: {
        query: req.query,
        params: req.params,
        lang: req.headers['lang'] || req.cookies.lang || req.query.lang,
        isMobile: MOBILE_DEVICE.includes(deviceType) || new RegExp(MOBILE_UA.join('|')).test(deviceUA)
      }
    }
  };

  return next()
}
