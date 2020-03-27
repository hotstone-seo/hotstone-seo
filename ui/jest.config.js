module.exports = {
  moduleFileExtensions: [
    'web.js', 'js', 'web.ts', 'ts', 'web.tsx', 'tsx', 'json', 'web.jsx', 'jsx', 'node',
  ],
  collectCoverageFrom: [
    'src/views/**/*.{js,jsx}',
    'src/components/**/*.{js,jsx}',
  ],
  moduleNameMapper: {
    '\\.(jpg|jpeg|png|gif|eot|otf|webp|svg|ttf|woff|woff2|mp4|webm|wav|mp3|m4a|aac|oga)$': '<rootDir>/tests/mocks/fileMock.js',
    '\\.(s?css)$': '<rootDir>/tests/mocks/styleMock.js',
  },
  verbose: true,
  testURL: 'http://localhost',
};
