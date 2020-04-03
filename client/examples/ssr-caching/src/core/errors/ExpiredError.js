import ErrorImage from './assets/payment-expired.svg';

class ExpiredError extends Error {
  constructor(message = i18n('error.expired.errorMessage'), ...params) {
    super(message, ...params);

    if (Error.captureStackTrace) {
      Error.captureStackTrace(this, ExpiredError);
    }

    this.image = ErrorImage;
    this.message = message;
    this.buttonText = i18n('error.expired.buttonText');
    this.link = CONFIG.LINK.myOrder;
  }
}

export default ExpiredError;
