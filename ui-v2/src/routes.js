import React from 'react';
import {
  HomeOutlined, FormOutlined, DatabaseOutlined, TagsOutlined
} from '@ant-design/icons';

const Rule = React.lazy(() => import('./views/Rule'));
const MismatchRule = React.lazy(() => import('./views/MismatchRule'));
const DataSource = React.lazy(() => import('./views/DataSource'));

const routes = [
  {
    path: '/',
    exact: true,
    name: 'Home',
    icon: HomeOutlined,
  },
  {
    path: '/rules',
    name: 'Rules',
    component: Rule,
    icon: FormOutlined,
  },
  {
    path: '/datasources',
    name: 'Data Sources',
    component: DataSource,
    icon: DatabaseOutlined,
  },
  {
    path: '/mismatch-rule',
    name: 'Mismatch Rule',
    component: MismatchRule,
    icon: TagsOutlined,
  },
];

export default routes;
