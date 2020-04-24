import {
  FormOutlined,
  DatabaseOutlined,
  TagsOutlined,
  AreaChartOutlined,
  PlayCircleOutlined,
  AuditOutlined,
} from '@ant-design/icons';

import Rule from 'views/Rule';
import MismatchRule from 'views/MismatchRule';
import DataSource from 'views/DataSource';
import Analytic from 'views/Analytic';
import Simulation from 'views/Simulation';
import AuditTrail from 'views/AuditTrail';

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
