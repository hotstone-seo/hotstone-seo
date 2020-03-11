import React from 'react';
import {
  render, wait,
} from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import mockAxios from 'jest-mock-axios';

import { configure, mount } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import ViewDataSources from './ViewDataSources';
import AddDataSource from './AddDataSource';

configure({ adapter: new Adapter() });

const respMock = [{
  id: 1, name: 'FooDS', url: '/foo-ds', updated_at: new Date().toISOString(), created_at: new Date().toISOString(),
}];

jest.mock('react-router-dom', () => ({
  useHistory: () => ({
    push: jest.fn(),
  }),
}));


/*
test('View Mismatch rule unit test', () => {
  const props = {
    match: { url: 'tes.com' },
  };
  const tree = mount(<ViewDataSources {...props} />);
  expect(tree).toMatchSnapshot();
}); */

const props = {
  onEdit: jest.fn(),
  onDelete: jest.fn(),
  handleSubmit: jest.fn()
};

describe('ViewDataSources', () => {
  test('first load data source list', async () => {
    const url = '/datasources';
    const {
      queryByText,
    } = render(<ViewDataSources match={{ url }} />);

    await wait(() => {
      expect(mockAxios.get).toHaveBeenCalledWith('/data_sources');
      mockAxios.mockResponse({ data: respMock });

      expect(queryByText(/Data Sources/)).toBeInTheDocument();
      expect(queryByText(/FooDS/)).toBeInTheDocument();
    });
  });

  test('event click back button', () => { 
    global.window.matchMedia = function () { return { matches: false, addListener() {}, removeListener() {} }; };
    const tree = mount(<AddDataSource />);
    expect(tree.find('.ant-page-header-back-button')).toHaveLength(2);
    tree.find('.ant-page-header-back-button').at(0).simulate('click');
    tree.find('.ant-page-header-back-button').at(1).simulate('click');
    tree.unmount();
  });
});
