import React from 'react';
import ScrollToTop from '../ScrollToTop.jsx';

test('Should render ScrollToTop correctly', () => {
  const tree = enzyme.mount(
    <ScrollToTop.WrappedComponent location="/">
      <div style={{ height: 500 }} />
    </ScrollToTop.WrappedComponent>
  );

  tree.setProps({
    location: '/test'
  });
  tree.setProps({
    location: '/test2'
  });
  expect(tree).toMatchSnapshot();
});
