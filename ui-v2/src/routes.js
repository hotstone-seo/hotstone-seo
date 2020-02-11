import React from 'react';

const Rule = React.lazy(() => import('./views/Rule'));

const TagList = React.lazy(() => import('./components/Tag/TagList'));

const routes = [
  { path: '/', exact: true, name: 'Home' },
  { path: '/rules', name: 'Rules', component: Rule },
  { path: '/tags', name: 'Tags', component: TagList }
]

export default routes;
