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

const COMPONENT_MAP = {
  rule: Rule,
  datasources: DataSource,
  mismatchrule: MismatchRule,
  analytic: Analytic,
  simulation: Simulation,
  audittrail: AuditTrail,
  user: User,
  clientkey: ClientKey,
  roletype: RoleType,
  module: Module,
  notfound: GenericNotFound,
};

const ICON_MAP = {
  rule: FormOutlined,
  datasources: DatabaseOutlined,
  mismatchrule: TagsOutlined,
  analytic: AreaChartOutlined,
  simulation: PlayCircleOutlined,
  audittrail: AuditOutlined,
  user: UserOutlined,
  clientkey: LockOutlined,
  roletype: UsergroupAddOutlined,
  module: MenuOutlined,
};

// TODO : Label menu will use label from API
const LABEL_MAP = {
  rule: 'Rules',
  datasources: 'Data Sources',
  mismatchrule: 'Mismatch Rule',
  analytic: 'Analytic',
  simulation: 'Simulation',
  audittrail: 'Audit Trail',
  user: 'Users',
  clientkey: 'Client Keys',
  roletype: 'Role User',
  module: 'Modules',
};

const token = Cookies.get('token');
const tokenDecoded = token !== undefined ? jwt.decode(token) : undefined;

let isOldCookieVersion = false;
if (tokenDecoded !== undefined) isOldCookieVersion = tokenDecoded.modules === undefined;

let routes = [];

if (tokenDecoded !== undefined && isOldCookieVersion === false) { // status : already login
  let jsonModules = tokenDecoded !== undefined ? tokenDecoded.modules : [];
  let mn;
  jsonModules = JSON.parse(jsonModules);

  Object.keys(jsonModules).forEach((key) => {
    mn = jsonModules[key];
  });

  const arrMenu = [];
  mn.forEach((item) => {
    const tempMenu = [];
    const isAnyLabel = item.label !== undefined;
    tempMenu.path = '/'.concat(item.path);
    if (isAnyLabel) tempMenu.name = item.label;
    else tempMenu.name = LABEL_MAP[item.name];
    tempMenu.component = COMPONENT_MAP[item.name];
    tempMenu.icon = ICON_MAP[item.name];
    tempMenu.visible = true;
    arrMenu.push(tempMenu);
  });

  const menu404 = [];
  menu404.path = '*';
  menu404.component = COMPONENT_MAP['notfound'];
  arrMenu.push(menu404);

  routes = arrMenu;
} else { // status : still not login
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
