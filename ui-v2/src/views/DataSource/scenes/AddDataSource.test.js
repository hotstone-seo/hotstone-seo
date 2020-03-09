import React from 'react';
import {
  render, fireEvent,
} from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import mockAxios from 'jest-mock-axios';
import AddDataSource from './AddDataSource';
import { Form, Input } from "antd";

const dataSource = [{
  name: 'data source flight', url: 'localhost/flight',
}];

jest.mock('react-router-dom', () => ({
  useHistory: () => ({
    push: jest.fn(),
  }),
}));

const form = render(<AddDataSource />);
const setup = () => {
  const utils = form;
  const input = utils.queryByText('Name').toBeInTheDocument();
  return {
    input,
    ...utils,
  };
};

describe('AddDataSource', () => {
  test('name of data source', () => {
    const { input } = setup();
    fireEvent.change(input, { target: { value: 'data source airport' } });
    expect(input.value).toBe('data source airport');
  });

  test('save new data source', () => {
  //  form.simulate('click');
  //  expect(saveDataSource).toHaveBeenNthCalledWith(1, dataSource);
  });
});
