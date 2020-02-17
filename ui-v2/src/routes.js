import React from 'react';

const Rule = React.lazy(() => import('./views/Rule'));

const DataSource = React.lazy(() => import('./views/DataSource'));

const routes = [
  { path: '/', exact: true, name: 'Home' },
  { path: '/rules', name: 'Rules', component: Rule },
  { path: '/datasources', name: 'Data Sources', component: DataSource },
];

export default routes;
