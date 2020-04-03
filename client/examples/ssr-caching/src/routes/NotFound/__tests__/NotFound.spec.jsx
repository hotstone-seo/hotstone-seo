import React from 'react';
import NotFound from '../';

test('Should render NotFound correctly', () => {
  const tree = mount(<NotFound />);
  
  expect(tree).toMatchSnapshot();
});
