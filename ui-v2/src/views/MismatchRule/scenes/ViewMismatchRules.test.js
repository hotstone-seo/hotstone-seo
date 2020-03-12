import React from 'react';
import { createMemoryHistory } from 'history';
import { Router } from 'react-router-dom';
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
  url: 1, first_seen: new Date(), last_seen: new Date(), count: 2,
}];

jest.mock('react-router-dom', () => ({
  useHistory: () => ({
    push: jest.fn(),
  }),
}));

/*
test('View Rules unit test', () => {
  global.window.matchMedia = function () { return { matches: false, addListener() {}, removeListener() {} }; };
  const props = {
    match: { url: 'tes.com' },
  };
  const tree = mount(<ViewMismatchRules {...props} />);
  expect(tree).toMatchSnapshot();
});
*/
describe('View mismatched rule', () => {
  test('first load', async () => {
    global.window.matchMedia = function () { return { matches: false, addListener() {}, removeListener() {} }; };
    const url = '/metrics/mismatched';
    const {
      queryByText, container,
    } = render(<ViewMismatchRules match={{ url }} />, { initialEntries: [url] });

    await wait(() => {
      const queryParam = { params: { _limit: 5, _next_key: '-count', _sort: '-last_seen' } };
      expect(mockAxios.get).toHaveBeenCalledWith('metrics/mismatched', queryParam);
      mockAxios.mockResponse({ data: respMock });
    });
  });
});
