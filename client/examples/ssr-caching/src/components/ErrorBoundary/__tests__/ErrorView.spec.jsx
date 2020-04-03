import React from 'react';
import ErrorView from '../ErrorView.jsx';

test('Should render ErrorView correctly', () => {
  const tree = shallow(<ErrorView error={new Error('test')} info={{}}/>);
  
  expect(tree).toMatchSnapshot();
});
