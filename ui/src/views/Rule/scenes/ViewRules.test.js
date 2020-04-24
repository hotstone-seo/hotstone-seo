import React from 'react';
import { MemoryRouter } from 'react-router-dom';
import {
  cleanup, render, wait,
} from '@testing-library/react';

import '@testing-library/jest-dom/extend-expect';
import mockAxios from 'jest-mock-axios';
import ViewRules from './ViewRules';

afterEach(cleanup);

const respMock = [{
  id: 1, name: 'Airport Rule', url_pattern: '/airport', updated_at: '2020-04-01', created_at: new Date(),
}];

describe('View Rules', () => {
  test('first load', async () => {
    const {
      queryByText,
    } = render(<ViewRules match={{ url: '/rules' }} />, { wrapper: MemoryRouter });

    const queryParam = { params: { _limit: 10, _offset: 0 } };
    expect(mockAxios.get).toHaveBeenCalledWith('/rules', queryParam);
    mockAxios.mockResponse({ data: respMock });

    await wait();

    const firstRow = queryByText('Airport Rule');
    expect(firstRow).not.toBeNull();
  });
});
