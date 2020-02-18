import React, { useEffect, useMemo } from "react";
import { useTable, useSortBy, useFilters, usePagination } from "react-table";
import { useTokenPagination } from "react-table/src/utility-hooks/useTokenPagination";
import { fetchMismatched } from "api/metric";
import _ from "lodash";

// Define a default UI for filtering
function DefaultColumnFilter({ column: props }) {
  const { Header, filterValue, preFilteredRows, setFilter } = props;

  // console.log("@@@ PROPS: ", props);

  const count = preFilteredRows.length;

  return (
    <input
      value={filterValue || ""}
      onChange={e => {
        console.log("e.target.value: ", e.target.value);
        setFilter(e.target.value);
      }}
      placeholder={`Search ${Header}`}
    />
  );
}

function Table({
  columns,
  data,
  fetchData,
  loading,
  pageSize,
  setPageSize,
  initialSortBy
}) {
  const defaultColumn = useMemo(
    () => ({
      // Let's set up our default Filter UI
      Filter: DefaultColumnFilter
    }),
    []
  );

  const props = useTable(
    {
      columns,
      data,
      defaultColumn,
      initialState: { pageIndex: 0, pageSize: pageSize, sortBy: initialSortBy }, // Pass our hoisted table state
      manualPagination: true, // Tell the usePagination
      manualSortBy: true
      // hook that we'll handle our own data fetching
    },

    useFilters,

    useSortBy,
    // usePagination,
    useTokenPagination
  );

  const {
    getTableProps,
    getTableBodyProps,
    headerGroups,
    prepareRow,
    rows,

    setNextPageToken,
    pageToken,
    // pageIndex,
    previousPage,
    nextPage,
    canPreviousPage,
    canNextPage,
    resetPagination,

    state: { pageIndex, sortBy, filters }
  } = props;

  console.log("useTokenPagination props: ", props);

  useEffect(() => {
    console.log("!!! FILTERS: ", filters);
  }, [filters]);

  // Listen for changes in pagination and use the state to fetch our new data
  useEffect(() => {
    fetchData({
      pageIndex,
      pageSize,
      pageToken,
      setNextPageToken,
      sortBy,
      filters
    });
  }, [
    fetchData,
    pageIndex,
    pageSize,
    pageToken,
    setNextPageToken,
    sortBy
    // filters
  ]);

  // Render the UI for your table
  return (
    <>
      <pre>
        <code>
          {JSON.stringify(
            {
              pageIndex,
              pageSize,
              pageToken,
              sortBy,
              filters,
              canNextPage,
              canPreviousPage
            },
            null,
            2
          )}
        </code>
      </pre>
      <table {...getTableProps()}>
        <thead>
          {headerGroups.map(headerGroup => (
            <tr {...headerGroup.getHeaderGroupProps()}>
              {headerGroup.headers.map(column => (
                <th {...column.getHeaderProps(column.getSortByToggleProps())}>
                  {column.render("Header")}
                  <span>
                    {column.isSorted
                      ? column.isSortedDesc
                        ? " 🔽"
                        : " 🔼"
                      : ""}
                  </span>
                  <div>{column.canFilter ? column.render("Filter") : null}</div>
                </th>
              ))}
            </tr>
          ))}
        </thead>
        <tbody {...getTableBodyProps()}>
          {rows.map(row => {
            prepareRow(row);
            return (
              <tr {...row.getRowProps()}>
                {row.cells.map(cell => {
                  return (
                    <td {...cell.getCellProps()}>{cell.render("Cell")}</td>
                  );
                })}
              </tr>
            );
          })}
          <tr>
            {loading ? (
              // Use our custom loading state to show a loading indicator
              <td colSpan="10000">Loading...</td>
            ) : null}
          </tr>
        </tbody>
      </table>
      {/* 
        Pagination can be built however you'd like. 
        This is just a very basic UI implementation:
      */}
      <div className="pagination">
        <button onClick={() => resetPagination()} disabled={!canPreviousPage}>
          {"Latest"}
        </button>{" "}
        <button onClick={() => previousPage()} disabled={!canPreviousPage}>
          {"Prev"}
        </button>{" "}
        <button onClick={() => nextPage()} disabled={!canNextPage}>
          {"Next"}
        </button>{" "}
      </div>
    </>
  );
}

function ViewMismatchRules() {
  const columns = React.useMemo(
    () => [
      {
        Header: "URL",
        accessor: "request_path",
        disableSortBy: true,
        disableFilters: true
      },
      {
        Header: "Since",
        accessor: "since",
        disableSortBy: true,
        disableFilters: true
      },
      {
        Header: "Count",
        accessor: "count",
        disableSortBy: true,
        disableFilters: true
      }
    ],
    []
  );

  // We'll start our table without any data
  const [data, setData] = React.useState([]);
  const [pageSize, setPageSize] = React.useState(10);
  const [loading, setLoading] = React.useState(false);
  const fetchIdRef = React.useRef(0);

  const initialSortBy = useMemo(() => {
    return [
      { id: "since", desc: true }
      // { id: "count", desc: true } // '-count' as next_key
    ];
  }, []);

  const fetchData = React.useCallback(
    ({ pageSize, pageIndex, pageToken, setNextPageToken, sortBy, filters }) => {
      // // This will get called when the table needs new data
      // // You could fetch your data from literally anywhere,
      // // even a server. But for this example, we'll just fake it.

      // // Give this fetch an ID
      // const fetchId = ++fetchIdRef.current;

      // // Set the loading state
      // setLoading(true);

      // // We'll even set a delay to simulate a server here
      // setTimeout(() => {
      //   // Only update the data if this is the latest fetch
      //   if (fetchId === fetchIdRef.current) {
      //     const startRow = pageSize * pageIndex;
      //     const endRow = startRow + pageSize;
      //     // setData(serverData.slice(startRow, endRow));

      //     setLoading(false);
      //   }
      // }, 1000);

      // try {
      //   const queryParam = buildQueryParam(pageSize, sortBy, filters);
      //   const data = await fetchMismatched({ params: queryParam });
      //   // setData(data);
      // } catch (err) {
      //   console.log("ERR: ", err);
      // }

      console.log(
        "PAGE_INDEX: ",
        pageIndex,
        " PAGE_TOKEN: ",
        pageToken,
        " SET_NEXT_PAGE_TOKENB: ",
        setNextPageToken
      );
      const nextKey = { id: "count", desc: true };
      const queryParam = buildQueryParam(
        pageSize,
        sortBy,
        filters,
        nextKey,
        pageToken
      );
      // queryParam["_next_key"] = "-count";
      fetchMismatched({ params: queryParam }).then(data => {
        console.log("$$$ DATA: ", data);
        console.log("$$$ queryParam: ", queryParam);

        if (!_.isEmpty(data)) {
          const lastRow = data[data.length - 1];

          const nextPageToken = createPageToken(lastRow, sortBy, nextKey);
          // setNextPageToken(nextPageToken);
        }

        setData(data);
      });
    },
    []
  );

  return (
    <Table
      columns={columns}
      data={data}
      fetchData={fetchData}
      loading={loading}
      pageSize={pageSize}
      setPageSize={setPageSize}
      initialSortBy={initialSortBy}
    />
  );
}

export default ViewMismatchRules;

function createPageToken(lastRow, sortBy, nextKey) {
  var pageToken = [];
  sortBy.map(({ id, desc }) => {
    pageToken.push({ id: id, desc: desc, val: lastRow[id] });
  });

  pageToken.push({ ...nextKey, val: lastRow[nextKey.id] });
  return pageToken;
}

function buildQueryParam(pageSize, sortBy, filters, nextKey, pageToken) {
  console.log(">>> pageSize: ", pageSize);
  console.log(">>> sortBy: ", sortBy);
  console.log(">>> filters: ", filters);

  var queryParam = {};

  sortBy.map(({ id, desc }) => {
    const orderSign = desc ? "-" : "";
    var _sort = queryParam["_sort"];
    _sort = (_sort == undefined ? "" : `${_sort},`) + `${orderSign}${id}`;
    queryParam["_sort"] = _sort;
  });

  if (!_.isEmpty(nextKey)) {
    const orderSign = nextKey.desc ? "-" : "";
    queryParam["_next_key"] = `${orderSign}${nextKey.id}`;
  }

  if (!_.isEmpty(pageToken)) {
    pageToken.map(({ id, desc, val }) => {
      var _next = queryParam["_next"];
      _next = (_next == undefined ? "" : `${_next},`) + `${val}`;
      queryParam["_next"] = _next;
    });
  }

  // Object.entries(filters).forEach(([key, value]) => {
  //   if (!_.isEmpty(value)) {
  //     queryParam[key] = `%${value[0]}%`;
  //   }
  // });

  queryParam["_limit"] = pageSize;

  console.log(">>> queryParam: ", queryParam);
  return queryParam;
}
