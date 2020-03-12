import React from 'react';
import { MemoryRouter } from 'react-router-dom';
import {
  cleanup, render, wait,
} from '@testing-library/react';

import '@testing-library/jest-dom/extend-expect';
import userEvent from '@testing-library/user-event';
import mockAxios from 'jest-mock-axios';
import AddDataSource from './AddDataSource';

afterEach(cleanup);

const mockHistoryPush = jest.fn();

jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useHistory: () => ({
    push: mockHistoryPush,
  }),
}));

describe('AddDataSource', () => {
  test('test add new data source', async () => {
    const {
      getByTestId,
    } = render(<AddDataSource />, { wrapper: MemoryRouter });

    const nameInput = getByTestId('input-name');
    await userEvent.type(nameInput, 'Datasource');
    expect(nameInput.value).toBe('Datasource');

    const urlInput = getByTestId('input-url');
    await userEvent.type(urlInput, '/ds');
    expect(urlInput.value).toBe('/ds');

    const saveBtn = getByTestId('btn-save');
    userEvent.click(saveBtn);

    await wait();

    // expect(mockAxios.post).toHaveBeenCalledWith('/data_sources', { name: 'Datasource', url: '/ds' });

    // const msgExpected = { message: { level: 'success', content: 'Datasource is successfully created' } };
    // expect(mockHistoryPush).toHaveBeenCalledWith('/data_sources', msgExpected);
  });
});
