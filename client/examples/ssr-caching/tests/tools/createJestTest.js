function createJestTest(filename, originalFile) {
  const fileWithOutExt = filename.replace(/\.[^/.]+$/, '');
  const readableFilename = capitalizeFirstLetter(fileWithOutExt);

  return `import React from 'react';
import ${readableFilename} from '../${originalFile.replace('index.jsx', '')}';

test('Should render ${readableFilename} correctly', () => {
  const tree = mount(<${readableFilename} />);
  
  expect(tree).toMatchSnapshot();
});
`;
}

const capitalizeFirstLetter = string => string.charAt(0).toUpperCase() + string.slice(1);

module.exports = { createJestTest };
