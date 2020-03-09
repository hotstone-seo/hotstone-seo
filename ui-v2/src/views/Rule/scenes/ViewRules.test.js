import React from 'react';
import {
  render, wait,
} from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import mockAxios from 'jest-mock-axios';
import ViewRules from './ViewRules';

const respMock = [{
  id: 1, name: 'Airport Rule', url: '/foo-ds', url_pattern: '/airport', updated_at: new Date().toISOString(), created_at: new Date().toISOString(),
}];

jest.mock('react-router-dom', () => ({
  useHistory: () => ({
    push: jest.fn(),
  }),
}));

describe('View Rules', () => {
  test('first load', async () => {
    const url = '/rules';
    const {
      queryByText,
    } = render(<ViewRules match={{ url }} />);

    await wait(() => {
      expect(mockAxios.get).toHaveBeenCalledWith('/rules');
      mockAxios.mockResponse({ data: respMock });

      expect(queryByText(/Rules/)).toBeInTheDocument();
      //expect(queryByText(/Airport Rule/)).toBeInTheDocument();
    });
  });
});
