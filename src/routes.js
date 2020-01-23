import React from 'react';

const Rule = React.lazy(() => import('./views/Rule'));

const routes = [
  { path: '/', exact: true, name: 'Home' },
  { path: '/rules', name: 'Rules', component: Rule }
]

export default routes;
