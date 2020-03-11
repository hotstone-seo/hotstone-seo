import React from 'react';
import {
  render, queryByPlaceholderText,
} from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import mockAxios from 'jest-mock-axios';
import { mount, configure } from 'enzyme';

import Adapter from 'enzyme-adapter-react-16';

import { JSDOM } from 'jsdom';
import AddDataSource from './AddDataSource';

// const dataSource = [{
//  name: 'data source flight', url: 'localhost/flight',
// }];

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
  test('Should render Name correctly', () => {
    global.window.matchMedia = function () {
      return {
        matches: false,
        addListener() {},
        removeListener() {},
      };
    };
    const tree = mount(<AddDataSource />);
    expect(tree).toMatchSnapshot();
  });

  test('name of data source', () => {
    global.window.matchMedia = function () {
      return {
        matches: false,
        addListener() {},
        removeListener() {},
      };
    };

    const wrapper = mount(<AddDataSource />);
    // expect(wrapper.contains('data-cy')).toBe(true);
    // console.log(wrapper.html(), 'wrapp');

    // expect(wrapper.find('Input [data-cy=name]').length).toBe(1); // Has Result component
    // expect(wrapper.find('input [data-cy="name"]')).toHaveLength(1); // Has Result component
    // expect(wrapper.find('#url')).to.have.toLengthOf(1); // Has Result component
    // expect(wrapper.find('Name').length).toBe(1); // Has Result component

    // expect(wrapper.contains(<input placeholder="My Data Source" maxlength="100" data-cy="name" type="text" id="name" class="ant-input" value="" />)).toBe(true);
    // console.log(wrapper.find('input').at(0).name(), 'iddd');
    // expect(wrapper.find('input').at(0).i.toBe('name');
    //  expect(wrapper.find(<input placeholder="My Data Source" />).length).toBe(1);

  // expect(wrapper.find('input').at(0).name).toEqual('text');
  // expect(wrapper.state().data).toBe('something');
  // console.log(wrapper.find('').html(), 'tes');
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

  /*
  it('simulates click events', () => {
    const onButtonClick = sinon.spy();
    const wrapper = mount((
      <Foo onButtonClick={onButtonClick} />
    ));
    wrapper.find('button').simulate('click');
    expect(onButtonClick).to.have.property('callCount', 1);
  });

    */
  test('save new data source', () => {
    // const expectedArg = 'success save';
    const wrapper = mount(<AddDataSource />);
    wrapper.find('form').simulate('submit');
    // expect(window.alert).toHaveBeenCalledWith(expectedArg);
  });

  test('validasi', () => {
    const tree = mount(<AddDataSource />);
    const event = {
      preventDefault() {},
      target: { value: '1234' }
    };
    tree.find('input[id="name"]').simulate('change', event);
    tree.unmount();
  });
});
