import React from 'react';
import Button from '../Button.jsx';

test('Should render Button correctly', () => {
  const tree = render(<Button />);

  expect(tree).toMatchSnapshot();
});
