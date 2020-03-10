import React from 'react';
import {
  render, wait,
} from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import mockAxios from 'jest-mock-axios';

import { shallow, configure, mount } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import ViewRules from './ViewRules';

configure({ adapter: new Adapter() });

const respMock = [{
  id: 1, name: 'Airport Rule', url: '/foo-ds', url_pattern: '/airport', updated_at: new Date().toISOString(), created_at: new Date().toISOString(),
}];

jest.mock('react-router-dom', () => ({
  useHistory: () => ({
    push: jest.fn(),
  }),
}));


test('View Rules unit test', () => {
  const props = {
    match: { url: 'tes.com' },
  };
  const tree = mount(<ViewRules {...props} />);
  expect(tree).toMatchSnapshot();
});

/*
describe('View Rules unit test', () => {
  test('Rule List', async () => {
    const url = '/rules';
    const {
      queryByText,
    } = render(<ViewRules match={{ url }} />);

    await wait(() => {


      expect(mockAxios.get).toHaveBeenNthCalledWith(2);
      // expect(mockAxios.get).toHaveBeenCalledWith('/rules');
      mockAxios.mockResponse({ data: respMock });

      expect(queryByText(/Rules/)).toBeInTheDocument();
      expect(queryByText(/Airport Rule/)).toBeInTheDocument();
    });
  });
});
*/
