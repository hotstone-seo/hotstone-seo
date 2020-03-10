import React from 'react';
import {
  render, wait,
} from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import mockAxios from 'jest-mock-axios';

import { shallow, configure, mount } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import ViewDataSources from './ViewDataSources';

configure({ adapter: new Adapter() });

const respMock = [{
  id: 1, name: 'FooDS', url: '/foo-ds', updated_at: new Date().toISOString(), created_at: new Date().toISOString(),
}];

jest.mock('react-router-dom', () => ({
  useHistory: () => ({
    push: jest.fn(),
  }),
}));

test('View Mismatch rule unit test', () => {
  const props = {
    match: { url: 'tes.com' },
  };
  const tree = mount(<ViewDataSources {...props} />);
  expect(tree).toMatchSnapshot();
});
/*
describe('ViewDataSources', () => {
  test('first load', async () => {
    const url = '/datasources';
    const {
      queryByText,
    } = render(<ViewDataSources match={{ url }} />);

    await wait(() => {
      expect(mockAxios.get).toHaveBeenCalledWith('/data_sources');
      mockAxios.mockResponse({ data: respMock });

      expect(queryByText(/Data Sources/)).toBeInTheDocument();
      expect(queryByText(/FooDS/)).toBeInTheDocument();
    });
  });
});
*/