import React from 'react';
import { MemoryRouter } from 'react-router-dom';
import {
  cleanup, render, wait, fireEvent,
} from '@testing-library/react';

import '@testing-library/jest-dom/extend-expect';
import mockAxios from 'jest-mock-axios';
import AddRule from './AddRule';

afterEach(cleanup);

const respMock = [{
  id: 1, name: 'Foo DS', url: '/foo-ds-url', updated_at: new Date(), created_at: new Date(),
}];

describe('Add Rule', () => {
  test('good case', async () => {
    const {
      queryByText,
      getByTestId,
      debug,
    } = render(<AddRule />, { wrapper: MemoryRouter });

    expect(mockAxios.get).toHaveBeenCalledWith('/data_sources');
    mockAxios.mockResponse({ data: respMock });

    await wait();

    const nameInput = getByTestId('input-name');
    fireEvent.change(nameInput, { target: { value: 'Rule ABC' } });
    expect(nameInput.value).toBe('Rule ABC');

    const urlPatternInput = getByTestId('input-url-pattern');
    fireEvent.change(urlPatternInput, { target: { value: '/abc' } });
    expect(urlPatternInput.value).toBe('/abc');

    await wait();

    const saveBtn = getByTestId('btn-save');
    fireEvent.click(saveBtn);

    await wait();

    expect(mockAxios.post).toHaveBeenCalledWith('/rules', { name: 'Rule ABC', url_pattern: '/abc' });
  });
});
