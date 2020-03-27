import React from 'react';
import { MemoryRouter } from 'react-router-dom';
import {
  cleanup, render, wait, act,
} from '@testing-library/react';

import '@testing-library/jest-dom/extend-expect';
import userEvent from '@testing-library/user-event';
import mockAxios from 'jest-mock-axios';
import AddRule from './AddRule';

afterEach(cleanup);

const respMock = [{
  id: 1, name: 'Foo DS', url: '/foo-ds-url', updated_at: new Date(), created_at: new Date(),
}];

describe('Add Rule', () => {
  test('good case', async () => {
    const {
      getByTestId,
    } = render(<AddRule />, { wrapper: MemoryRouter });

    act(() => {
      expect(mockAxios.get).toHaveBeenCalledWith('/data_sources');
      mockAxios.mockResponse({ data: respMock });
    });

    const nameInput = getByTestId('input-name');
    await userEvent.type(nameInput, 'Rule ABC');
    expect(nameInput.value).toBe('Rule ABC');

    const urlPatternInput = getByTestId('input-url-pattern');
    await userEvent.type(urlPatternInput, '/abc');
    expect(urlPatternInput.value).toBe('/abc');

    const saveBtn = getByTestId('btn-save');
    userEvent.click(saveBtn);

    await wait();

    expect(mockAxios.post).toHaveBeenCalledWith('/rules', { name: 'Rule ABC', url_pattern: '/abc' });
  });
});
