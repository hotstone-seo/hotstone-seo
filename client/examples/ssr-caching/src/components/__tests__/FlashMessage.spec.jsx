import React from 'react';
import FlashMessage from '../FlashMessage.jsx';

test('Should render FlashMessage correctly', () => {
  const resetFlash = jest.fn();
  const flash = {
    show: true,
    type: 'error',
    text: 'test'
  };
  const tree = enzyme.mount(
    <FlashMessage.WrappedComponent flash={flash} resetFlash={resetFlash} />
  );

  expect(tree).toMatchSnapshot();

  tree.setProps({
    flash: {
      ...flash,
      text: 'test2'
    }
  });
  tree.setProps({
    flash: {
      ...flash,
      show: false
    }
  });
});
