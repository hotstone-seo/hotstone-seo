import React from 'react';
import { MemoryRouter } from 'react-router-dom';
import {
  render, wait,
} from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import mockAxios from 'jest-mock-axios';
import ViewMismatchRules from './ViewMismatchRules';

const respMock = [{
  url: 1, first_seen: new Date(), last_seen: new Date(), count: 2,
}];

describe('View mismatched rule', () => {
  test('first load', async () => {
    const url = '/metrics/mismatched';
    const {
      queryByText, container,
    } = render(<ViewMismatchRules match={{ url }} />, { wrapper: MemoryRouter });

    await wait(() => {
      const queryParam = { params: { _limit: 5, _next_key: '-count', _sort: '-last_seen' } };
      expect(mockAxios.get).toHaveBeenCalledWith('metrics/mismatched', queryParam);
      mockAxios.mockResponse({ data: respMock });
    });
  });
});
