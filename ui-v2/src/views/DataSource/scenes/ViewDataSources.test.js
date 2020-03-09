import React from 'react';
import {
  render, wait,
} from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import mockAxios from 'jest-mock-axios';
import ViewDataSources from './ViewDataSources';

const respMock = [{
  id: 1, name: 'FooDS', url: '/foo-ds', updated_at: new Date().toISOString(), created_at: new Date().toISOString(),
}];

jest.mock('react-router-dom', () => ({
  useHistory: () => ({
    push: jest.fn(),
  }),
}));

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
