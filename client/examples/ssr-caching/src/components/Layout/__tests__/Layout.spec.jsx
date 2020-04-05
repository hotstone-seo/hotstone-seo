import React from 'react';
import Layout from '../index.js';

test('Should render Layout correctly', () => {
  const tree = mount(<Layout route={{ routes: [] }}/>);
  
  expect(tree).toMatchSnapshot();
});
