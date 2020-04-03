import ExpiredError from '../ExpiredError';
import ErrorImage from '../assets/payment-expired.svg';

const errorMessage = 'Test';

test('Should create ExpiredError correctly', () => {
  const error = new ExpiredError(errorMessage);
  const errorButton = i18n('error.expired.buttonText');

  expect(error.image).toEqual(ErrorImage);
  expect(error.message).toEqual(errorMessage);
  expect(error.buttonText).toEqual(errorButton);
  expect(error.link).toEqual(CONFIG.LINK.myOrder);
});
