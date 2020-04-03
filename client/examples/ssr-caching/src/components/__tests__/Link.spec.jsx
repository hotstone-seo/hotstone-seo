import React from 'react';
import Link from '../Link.jsx';

test('Should render Link correctly', () => {
  const tree = shallow(<Link to="/">Link</Link>);

  expect(tree).toMatchSnapshot();
});
