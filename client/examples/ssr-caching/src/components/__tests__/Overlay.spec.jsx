import React from 'react';
import Overlay from '../Overlay.jsx';

test('Should render Overlay correctly', () => {
  const tree = render(
    <Overlay>
      <div className="test">Overlay</div>
    </Overlay>
  );

  expect(tree).toMatchSnapshot();
});
