import {
  FormOutlined,
  DatabaseOutlined,
  TagsOutlined,
  AreaChartOutlined,
  PlayCircleOutlined,
  AuditOutlined,
  UserOutlined,
} from '@ant-design/icons';

import Cookies from 'js-cookie';
import jwt from 'jsonwebtoken';

import Rule from 'views/Rule';
import MismatchRule from 'views/MismatchRule';
import DataSource from 'views/DataSource';
import Analytic from 'views/Analytic';
import Simulation from 'views/Simulation';
import AuditTrail from 'views/AuditTrail';
import GenericNotFound from 'views/GenericNotFound';
import User from './views/User';

const COMPONENT_MAP = {
  rules: Rule,
  datasources: DataSource,
  mismatchrule: MismatchRule,
  analytic: Analytic,
  simulation: Simulation,
  audittrail: AuditTrail,
  user: User,
  notfound: GenericNotFound,
};

const ICON_MAP = {
  rules: FormOutlined,
  datasources: DatabaseOutlined,
  mismatchrule: TagsOutlined,
  analytic: AreaChartOutlined,
  simulation: PlayCircleOutlined,
  audittrail: AuditOutlined,
  user: UserOutlined,
};

const PATH_MAP = {
  rules: '/rules',
  datasources: '/datasources',
  mismatchrule: '/mismatch-rule',
  analytic: '/analytic',
  simulation: '/simulation',
  audittrail: '/audit-trail',
  user: '/users',
};

const LABEL_MAP = {
  rules: 'Rules',
  datasources: 'Data Sources',
  mismatchrule: 'Mismatch Rule',
  analytic: 'Analytic',
  simulation: 'Simulation',
  audittrail: 'Audit Trail',
  user: 'Users',
};

const token = Cookies.get('token');
const tokenDecoded = token !== undefined ? jwt.decode(token) : undefined;
const jsonModules = tokenDecoded !== undefined ? tokenDecoded.modules : [];
let routes = [];

let mn;
if (tokenDecoded !== undefined) {
  Object.keys(jsonModules).forEach((key) => {
    mn = jsonModules[key];
  });

  let arrMenu = [];
  mn.forEach((item, index) => {
    const tempMenu = [];
    tempMenu.path = PATH_MAP[item];
    tempMenu.name = LABEL_MAP[item];
    tempMenu.component = COMPONENT_MAP[item];
    tempMenu.icon = ICON_MAP[item];
    tempMenu.visible = true;
    arrMenu.push(tempMenu);
  });

  const menu404 = [];
  menu404.path = '*';
  menu404.component = COMPONENT_MAP['notfound'];
  arrMenu.push(menu404);

  routes = arrMenu;
} else {
  routes = [
    {
      path: '/rules',
      name: 'Rules',
      component: Rule,
      icon: FormOutlined,
      visible: true,
    },
    {
      path: '/datasources',
      name: 'Data Sources',
      component: DataSource,
      icon: DatabaseOutlined,
      visible: true,
    },
    {
      path: '/mismatch-rule',
      name: 'Mismatch Rule',
      component: MismatchRule,
      icon: TagsOutlined,
      visible: true,
    },
    {
      path: '/analytic',
      name: 'Analytic',
      component: Analytic,
      icon: AreaChartOutlined,
      visible: true,
    },
    {
      path: '/simulation',
      name: 'Simulation',
      component: Simulation,
      icon: PlayCircleOutlined,
      visible: true,
    },
    {
      path: '/audit-trail',
      name: 'Audit Trail',
      component: AuditTrail,
      icon: AuditOutlined,
      visible: true,
    },
    {
      path: '/users',
      name: 'User',
      component: User,
      icon: UserOutlined,
      visible: true,
    },
    {
      path: '*',
      component: GenericNotFound,
    },
  ];
}

export default routes;
