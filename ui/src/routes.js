import React from 'react';
import {
  FormOutlined,
  DatabaseOutlined,
  TagsOutlined,
  AreaChartOutlined,
  PlayCircleOutlined,
} from '@ant-design/icons';

const Rule = React.lazy(() => import('views/Rule'));
const MismatchRule = React.lazy(() => import('views/MismatchRule'));
const DataSource = React.lazy(() => import('views/DataSource'));
const Analytic = React.lazy(() => import('views/Analytic'));
const Simulation = React.lazy(() => import('views/Simulation'));

const routes = [
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
  {
    path: '/analytic',
    name: 'Analytic',
    component: Analytic,
    icon: AreaChartOutlined,
  },
  {
    path: '/simulation',
    name: 'Simulation',
    component: Simulation,
    icon: PlayCircleOutlined,
  },
];

export default routes;
