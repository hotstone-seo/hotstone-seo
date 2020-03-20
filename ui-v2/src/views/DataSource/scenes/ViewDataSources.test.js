import React from "react";
import { render, wait } from "@testing-library/react";
import "@testing-library/jest-dom/extend-expect";
import mockAxios from "jest-mock-axios";
import userEvent from "@testing-library/user-event";
// import { createMemoryHistory } from 'history';
import ViewDataSources from "./ViewDataSources";

const respMock = [
  {
    id: 1,
    name: "FooDS",
    url: "/foo-ds",
    updated_at: new Date().toISOString(),
    created_at: new Date().toISOString()
  }
];

jest.mock("react-router-dom", () => ({
  useHistory: () => ({
    push: jest.fn()
  })
}));

describe("ViewDataSources", () => {
  test("first load", async () => {
    const url = "/datasources";
    const { queryByText } = render(<ViewDataSources match={{ url }} />);

    await wait(() => {
      expect(mockAxios.get).toHaveBeenCalledWith("/data_sources");
      mockAxios.mockResponse({ data: respMock });

      expect(queryByText(/Data Sources/)).toBeInTheDocument();
      expect(queryByText(/FooDS/)).toBeInTheDocument();
    });
  });

  // FIXME:
  // test('handle delete data source', async () => {
  //   const url = '/datasources';
  //   const {
  //     getByTestId, queryByText,
  //   } = render(<ViewDataSources match={{ url }} />);

  //   await wait(() => {
  //     expect(mockAxios.get).toHaveBeenCalledWith('/data_sources');
  //     mockAxios.mockResponse({ data: respMock });

  //     expect(queryByText(/FooDS/)).toBeInTheDocument();
  //     // const firstRender = asFragment();
  //     const saveBtn = getByTestId('btn-delete');
  //     userEvent.click(saveBtn);

  //     expect(saveBtn).toBeInTheDocument();
  //     console.log('lolos', 'aa1');
  //     // const history = createMemoryHistory({ initialEntries: ['/aa'] });
  //     // userEvent.click(queryByText.getByText('Delete'));
  //     // expect(history.location.search).toBe('?id=1');
  //     // // expect(firstRender).toMatchDiffSnapshot(asFragment());
  //     // expect(history.location.pathname).toBe('/aa');
  //   });
  //   // expect(mockAxios.delete).toHaveBeenCalledWith('/datasources', { id: 20 });
  //   // console.log('lolos', 'aa2');
  // });

  /*
  test('handle edit data source', async () => {
    const history = createMemoryHistory({ initialEntries: ['/'] });
    const queryByText = render('<Link history={history}><Button>Edit</Button></Link>');
    expect(queryByText(/Edit/)).toBeInTheDocument();
    // userEvent.click(renderResult.getByText('Edit'));
    expect(history.location.pathname).toBe('/');
  });
  */

  // test('test aja', () => {
  //   console.log('asik', 'aa');
  //   const {
  //     getByText,
  //   } = render(<div><button type="button">Edit</button></div>);
  //   expect(getByText('Edit')).toBeInTheDocument();
  //   console.log('asik', 'aa');
  // //     userEvent.click(saveBtn);
  // });
});
