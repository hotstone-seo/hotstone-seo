import React from 'react';
import {
  FormOutlined,
  DatabaseOutlined,
  TagsOutlined,
  AreaChartOutlined,
  PlayCircleOutlined,
  AuditOutlined,
} from '@ant-design/icons';

const Rule = React.lazy(() => import('views/Rule'));
const MismatchRule = React.lazy(() => import('views/MismatchRule'));
const DataSource = React.lazy(() => import('views/DataSource'));
const Analytic = React.lazy(() => import('views/Analytic'));
const Simulation = React.lazy(() => import('views/Simulation'));
const AuditTrail = React.lazy(() => import('views/AuditTrail'));

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
  {
    path: '/audit-trail',
    name: 'Audit Trail',
    component: AuditTrail,
    icon: AuditOutlined,
  },
];

export default routes;
