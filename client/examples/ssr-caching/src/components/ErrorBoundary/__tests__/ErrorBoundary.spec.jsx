import React, { PureComponent } from 'react';
import ErrorBoundary from '../';
import ErrorView from '../ErrorView';

const TestError = () => <div>test</div>;

test('Should render ErrorBoundary correctly', () => {
  const tree = enzyme.mount(
    <ErrorBoundary render={(error, info) => <ErrorView error={error} info={info} />}>
      <TestError />
    </ErrorBoundary>
  );

  expect(tree).toMatchSnapshot();
});
