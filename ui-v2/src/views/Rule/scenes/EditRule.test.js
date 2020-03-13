import React from 'react';
import { MemoryRouter } from 'react-router-dom';
import {
  cleanup, render, wait, act,
} from '@testing-library/react';

import '@testing-library/jest-dom/extend-expect';
import userEvent from '@testing-library/user-event';
import mockAxios from 'jest-mock-axios';
import EditRule from './EditRule';

afterEach(cleanup);

const respMock = [{
  id: 1, name: 'Foo DS', url: '/foo-ds-url', updated_at: new Date(), created_at: new Date(),
}];

describe('Edit Rule', () => {
  test('good case', async () => {
    const ref = React.createRef();

    const {
      getByTestId, getByText
    } = render(<EditRule ref={ref} />, { wrapper: MemoryRouter });

    act(() => {
      expect(mockAxios.get).toHaveBeenCalledWith('/data_sources');
      mockAxios.mockResponse({ data: respMock });
    });

    expect(getByText('Manage Rule')).toBeInTheDocument();

    const nameInput = getByTestId('input-name');
    await userEvent.type(nameInput, 'Rule ABC Edit');
    expect(nameInput.value).toBe('Rule ABC Edit');

    const urlPatternInput = getByTestId('input-url-pattern');
    await userEvent.type(urlPatternInput, '/ruleedit');
    expect(urlPatternInput.value).toBe('/ruleedit');

    const saveBtn = getByTestId('btn-save');
    userEvent.click(saveBtn);

    await wait();
    expect(ref.current.getValue()).toEqual(ref.current.state.value);
    expect(mockAxios.post).toHaveBeenCalledWith('/rules', { name: 'Rule ABC Edit', url_pattern: '/ruleedit' });
  });
});
