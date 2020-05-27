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
} from "@ant-design/icons";

import Cookies from "js-cookie";
import jwt from "jsonwebtoken";

import Rule from "views/Rule";
import MismatchRule from "views/MismatchRule";
import DataSource from "views/DataSource";
import Analytic from "views/Analytic";
import Simulation from "views/Simulation";
import AuditTrail from "views/AuditTrail";
import ClientKey from "views/ClientKey";
import GenericNotFound from "views/GenericNotFound";
import User from "./views/User";
import RoleType from "./views/RoleType";
import Module from "./views/Module";
import Setting from "./views/Setting";

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
  setting: Setting,
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
  setting: SettingOutlined,
};

// TODO : Label menu will use label from API
const LABEL_MAP = {
  rule: "Rules",
  datasources: "Data Sources",
  mismatchrule: "Mismatch Rule",
  analytic: "Analytic",
  simulation: "Simulation",
  audittrail: "Audit Trail",
  user: "Users",
  clientkey: "Client Keys",
  roletype: "Role User",
  module: "Modules",
  setting: "Setting",
};

const token = Cookies.get("token");
const tokenDecoded = token !== undefined ? jwt.decode(token) : undefined;

let isOldCookieVersion = false;
if (tokenDecoded !== undefined)
  isOldCookieVersion = tokenDecoded.modules === undefined;

let routes = [];

if (tokenDecoded !== undefined && isOldCookieVersion === false) {
  // status : already login
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
    tempMenu.path = "/".concat(item.path);
    if (isAnyLabel) tempMenu.name = item.label;
    else tempMenu.name = LABEL_MAP[item.name];
    tempMenu.component = COMPONENT_MAP[item.name];
    tempMenu.icon = ICON_MAP[item.name];
    tempMenu.visible = true;
    arrMenu.push(tempMenu);
  });

  const menu404 = [];
  menu404.path = "*";
  menu404.component = COMPONENT_MAP["notfound"];
  arrMenu.push(menu404);

  routes = arrMenu;
}

export default routes;
