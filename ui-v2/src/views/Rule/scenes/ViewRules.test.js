import React from 'react';
import { createMemoryHistory } from 'history';
import { MemoryRouter, Router, Route } from 'react-router-dom';
import {
  cleanup, act,
  render, wait,
} from '@testing-library/react';
import { prettyDOM } from '@testing-library/dom';

import '@testing-library/jest-dom/extend-expect';
import mockAxios from 'jest-mock-axios';
import ViewRules from './ViewRules';

afterEach(cleanup);

const respMock = [{
  id: 1, name: 'Airport Rule', url: '/foo-ds', url_pattern: '/airport', updated_at: new Date().toISOString(), created_at: new Date().toISOString(),
}];

// jest.mock('react-router-dom', () => ({
//   useHistory: () => ({
//     push: jest.fn(),
//   }),
// }));

function renderWithRouter(children, historyConf = {}) {
  const history = createMemoryHistory(historyConf);
  return render(<Router history={history}>{children}</Router>);
}

describe('View Rules', () => {
  test('first load', async () => {
    const url = '/rules';
    const {
      queryByText, container,
    } = renderWithRouter(<ViewRules match={{ url }} />, { initialEntries: [url] });


    // act(() => {
    //   // example: click a <Link> to /products?id=1234
    // });

    // expect(queryByText('Rules')).toBeInTheDocument();
    // expect(queryByText('Airport Rule')).toBeInTheDocument();


    // const elem = getByText(container, 'Goodbye world'); // will fail by throwing error


    const queryParam = { params: { _limit: 10, _offset: 0 } };
    expect(mockAxios.get).toHaveBeenCalledWith('/rules', queryParam);
    mockAxios.mockResponse({ data: respMock });

    // console.log(`## CONTAINER: ${prettyDOM(container)}`);


    await wait(() => {
    //   const queryParam = { params: { _limit: 10, _offset: 0 } };
    //   expect(mockAxios.get).toHaveBeenCalledWith('/rules', queryParam);
    //   mockAxios.mockResponse({ data: respMock });

      // expect(queryByText(/No Data/)).toBeInTheDocument();

      // console.log(`## CONTAINER: ${prettyDOM(container)}`);

    //   // expect(queryByText(/Rules/)).toBeInTheDocument();
    });
  });
});
