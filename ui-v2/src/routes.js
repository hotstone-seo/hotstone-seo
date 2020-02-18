import React from 'react';

const Rule = React.lazy(() => import('./views/Rule'));
const MismatchRule = React.lazy(() => import('./views/MismatchRule'));
const DataSource = React.lazy(() => import('./views/DataSource'));

const routes = [
  { path: '/', exact: true, name: 'Home' },
  { path: '/rules', name: 'Rules', component: Rule },
  { path: '/datasources', name: 'Data Sources', component: DataSource },
  { path: '/mismatch-rule', name: 'Mismatch Rule', component: MismatchRule },
];

export default routes;
