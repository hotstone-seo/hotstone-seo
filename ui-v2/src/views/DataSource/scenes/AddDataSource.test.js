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


describe('AddDataSource', () => {

  test('save new data source', () => {
  //  form.simulate('click');
  //  expect(saveDataSource).toHaveBeenNthCalledWith(1, dataSource);
  });
});
