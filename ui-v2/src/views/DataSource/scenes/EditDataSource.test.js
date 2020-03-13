import React from 'react';
import { MemoryRouter } from 'react-router-dom';
import {
  cleanup, render, wait,
} from '@testing-library/react';

import '@testing-library/jest-dom/extend-expect';
import userEvent from '@testing-library/user-event';
import mockAxios from 'jest-mock-axios';
import EditDataSource from './EditDataSource';
import { createMemoryHistory } from "history";

afterEach(cleanup);

jest.mock('axios');

const mockHistoryPush = jest.fn();

jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useHistory: () => ({
    push: mockHistoryPush,
  }),
}));

const data1 = [{
  id: 1, name: 'FooDS', url: '/foo-ds', updated_at: new Date().toISOString(), created_at: new Date().toISOString(),
}];

describe('EditDataSource', () => {
  test('test edit data source', async () => {
    const history = createMemoryHistory({ initialEntries: ['/data_sources'] });
    const {
      getByTestId,
    } = render(<EditDataSource id={5} />, { wrapper: MemoryRouter });

    // mockAxios.get.mockResolvedValueOnce({
    //  data1,
    // });
    const nameInput = getByTestId('input-name');
    await userEvent.type(nameInput, 'Datasource');
    expect(nameInput.value).toBe('Datasource');

    const urlInput = getByTestId('input-url');
    await userEvent.type(urlInput, '/ds-edit');
    expect(urlInput.value).toBe('/ds-edit');

    const saveBtn = getByTestId('btn-save');
    userEvent.click(saveBtn);

    await wait();

    expect(history.location.pathname).toBe('/data_sources');
    // const msgExpected = { message: { level: 'success', content: 'Datasource is successfully created' } };
    // expect(mockHistoryPush).toHaveBeenCalledWith(1, '/data_sources');
    // console.log(mockHistoryPush.mock.call);
  });
});
