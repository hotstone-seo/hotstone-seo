import {
  FormOutlined,
  DatabaseOutlined,
  TagsOutlined,
  AreaChartOutlined,
  PlayCircleOutlined,
  AuditOutlined,
  UserOutlined,
  LockOutlined,
  UsergroupAddOutlined,
  MenuOutlined,
  SettingOutlined,
} from '@ant-design/icons';

import Cookies from 'js-cookie';
import jwt from 'jsonwebtoken';

import Rule from 'views/Rule';
import MismatchRule from 'views/MismatchRule';
import DataSource from 'views/DataSource';
import Analytic from 'views/Analytic';
import Simulation from 'views/Simulation';
import AuditTrail from 'views/AuditTrail';
import ClientKey from 'views/ClientKey';
import GenericNotFound from 'views/GenericNotFound';
import User from './views/User';
import RoleType from './views/RoleType';
import Module from './views/Module';
import Setting from './views/Setting';

function match(text, patterns) {
  let flag = false;
  patterns.forEach((p) => {
    if (text.match(p)) {
      flag = true;
    }
  });
  return flag;
}

function filter(routes, menuPatterns) {
  const newRoutes = [];
  routes.forEach((r) => {
    if (r.name && match(r.name, menuPatterns)) {
      newRoutes.push(r);
    }
  });
  return newRoutes;
}

const defaultRoutes = [
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
    path: '/role-type',
    name: 'User Role',
    component: RoleType,
    icon: UsergroupAddOutlined,
    visible: true,
  },
  {
    path: '/modules',
    name: 'Modules',
    component: Module,
    icon: MenuOutlined,
    visible: true,
  },
  {
    path: '/client-keys',
    name: 'Client Keys',
    component: ClientKey,
    icon: LockOutlined,
    visible: true,
  },
  {
    path: '/setting',
    name: 'Setting',
    component: Setting,
    icon: SettingOutlined,
    visible: true,
  },
  {
    path: '*',
    component: GenericNotFound,
  },
];
const token = Cookies.get('token');
const tokenDecoded = token !== undefined ? jwt.decode(token) : undefined;

let isOldCookieVersion = false;
if (tokenDecoded !== undefined) isOldCookieVersion = tokenDecoded.menus === undefined;

let routes = defaultRoutes;

if (tokenDecoded !== undefined && isOldCookieVersion === false) { // status : already login
  const arrayStr = tokenDecoded.menus;
  let i = 0;
  let tempResult = '';
  const regResult = [];
  for (i = 0; i < arrayStr.length; i++) {
    tempResult = new RegExp(arrayStr[i], 'i');
    regResult.push(tempResult);
  }
  routes = filter(defaultRoutes, regResult);
}
export default routes;
