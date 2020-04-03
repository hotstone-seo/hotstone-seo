import ExpiredError from '../PageError';

const errorMessage = 'Test';

test('Should create PageError correctly', () => {
  const error = new ExpiredError(errorMessage);

  expect(error.message).toEqual(errorMessage);
});
