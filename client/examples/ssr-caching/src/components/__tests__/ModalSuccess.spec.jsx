import React from 'react';
import ModalSuccess from '../ModalSuccess.jsx';

test('Should render ModalSuccess correctly', () => {
  const tree = enzyme.mount(<ModalSuccess />);

  console.log(tree.debug());
  expect(tree.find('div.modal-success').length).toEqual(1);
  expect(tree).toMatchSnapshot();
  tree.unmount();
  expect(tree.find('div.modal-success').length).toEqual(0);
});
