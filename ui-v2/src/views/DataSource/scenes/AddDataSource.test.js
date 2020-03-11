import React from 'react';
import {
  render, fireEvent, queryByPlaceholderText,
} from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import mockAxios from 'jest-mock-axios';
import { Form, Input } from 'antd';
import { exportAllDeclaration } from '@babel/types';
import { mount, configure } from 'enzyme';

import Adapter from 'enzyme-adapter-react-16';

import { JSDOM } from 'jsdom';
import AddDataSource from './AddDataSource';

const dataSource = [{
  name: 'data source flight', url: 'localhost/flight',
}];
configure({ adapter: new Adapter() });

jest.mock('react-router-dom', () => ({
  useHistory: () => ({
    push: jest.fn(),
  }),
}));

const jsdom = new JSDOM(
  `<!DOCTYPE html>
    <head>
      <script></script>
    </head>
    <body><div id="app"><!--APP--></div><div id="modal-root" ><!--MODAL--></div></body>
  </html>`,
);
const { window } = jsdom;


window.matchMedia = window.matchMedia
    || function () {
      return {
        matches: false,
        addListener() {},
        removeListener() {},
      };
    };

global.window = window;
 

jest.mock('react-router-dom', () => ({
  useHistory: () => ({
    push: jest.fn(),
  }),
}));

describe('AddDataSource', () => {
  test('name of data source', () => {
    global.window.matchMedia = function () {
      return {
        matches: false,
        addListener() {},
        removeListener() {},
      };
    };

    const wrapper = mount(<AddDataSource />);
    console.log(wrapper, "wrapp");
    // expect(wrapper.find('Name').length).toBe(1); // Has Result component
    // expect(wrapper.find('Name').prop('type')).toBe('success');
    // const inpt = wrapper.find(Form.Item).at(0);
    // fireEvent.change(inpt, { target: { value: 'test' } });
    // expect(inpt.value).toBe('test');
    // const { input } = setup();
    // console.log('AA');
    // fireEvent.change(input, { target: { value: 'data source airport' } });
    // console.log('AAA');
    // expect(input.value).toBe('data source airport');
  });


  test('save new data source', () => {
    const expectedArg = 'success save';
    const wrapper = mount(<AddDataSource />);
    wrapper.find('form').simulate('submit');
    // expect(window.alert).toHaveBeenCalledWith(expectedArg);
  });
});
