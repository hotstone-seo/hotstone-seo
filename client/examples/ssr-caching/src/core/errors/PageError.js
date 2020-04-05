class PageError extends Error {
  constructor(message = 'Unknown Error, please try again later.', ...params) {
    super(message, ...params);

    if (Error.captureStackTrace) {
      Error.captureStackTrace(this, PageError);
    }

    this.message = message;
  }
}

export default PageError;
