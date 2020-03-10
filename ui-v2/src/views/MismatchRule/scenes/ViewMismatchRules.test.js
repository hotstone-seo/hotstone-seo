import React from 'react';
import {
  render, wait,
} from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import mockAxios from 'jest-mock-axios';

import { configure, mount } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import ViewMismatchRules from './ViewMismatchRules';

configure({ adapter: new Adapter() });

const respMock = [{
  url: 1, first_seen: '03 Mar 2020 - 08:00(7 days ago)', last_seen: '05 Mar 2020 - 08:00(5 days ago)', count: 2,
}];

jest.mock('react-router-dom', () => ({
  useHistory: () => ({
    push: jest.fn(),
  }),
}));

test('View Rules unit test', () => {
  global.window.matchMedia = function () { return { matches: false, addListener() {}, removeListener() {} }; };
  const props = {
    match: { url: 'tes.com' },
  };
  const tree = mount(<ViewMismatchRules {...props} />);
  expect(tree).toMatchSnapshot();
});
