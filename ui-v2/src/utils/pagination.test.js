import { buildQueryParam } from './pagination';

describe('pagination', () => {
  test('good case', () => {
    const pagination = { current: 2, pageSize: 3, total: 6 };
    const queryParam = buildQueryParam(pagination, null, null);
    expect(queryParam).toStrictEqual({ _limit: 6, _offset: 3 });
  });
});

describe('filter', () => {
  test('good case', () => {
    const filter = { name: 'fooname', url_pattern: 'foourl' };
    const queryParam = buildQueryParam(null, filter, null);
    expect(queryParam).toStrictEqual({ name: '%fooname%', url_pattern: '%foourl%' });
  });
});

describe('sort', () => {
  test('ascend', () => {
    const sort = { order: 'ascend', field: 'name' };
    const queryParam = buildQueryParam(null, null, sort);
    expect(queryParam).toStrictEqual({ _sort: 'name' });
  });
  test('descend', () => {
    const sort = { order: 'descend', field: 'url_pattern' };
    const queryParam = buildQueryParam(null, null, sort);
    expect(queryParam).toStrictEqual({ _sort: '-url_pattern' });
  });
});
