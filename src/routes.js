import React from 'react';
import MetatagForm from './views/Metatag/MetatagForm';
 
const BrandButtons = React.lazy(() => import('./views/Buttons/BrandButtons'));
const ButtonDropdowns = React.lazy(() => import('./views/Buttons/ButtonDropdowns'));
const ButtonGroups = React.lazy(() => import('./views/Buttons/ButtonGroups'));
const Buttons = React.lazy(() => import('./views/Buttons/Buttons'));
const Dashboard = React.lazy(() => import('./views/Dashboard'));
const CoreUIIcons = React.lazy(() => import('./views/Icons/CoreUIIcons'));
const SimpleLineIcons = React.lazy(() => import('./views/Icons/SimpleLineIcons'));
const Users = React.lazy(() => import('./views/Users/Users'));
const User = React.lazy(() => import('./views/Users/User'));
const RuleList = React.lazy(() => import('./views/Rule/RuleList'));
const RuleForm = React.lazy(() => import('./views/Rule/RuleForm'));
const RuleDetail = React.lazy(() => import('./views/Rule/RuleDetail'));
const MetaTagList = React.lazy(() => import('./views/Metatag/Metatag'));
const MetatagPreview = React.lazy(() => import('./views/Metatag/MetatagPreview'));
const MetaTagEditForm = React.lazy(() => import('./views/Metatag/MetatagEditForm'));
const DataSourceList = React.lazy(() => import('./views/DataSource/DataSource'));
const DataSourceForm = React.lazy(() => import('./views/DataSource/DataSourceForm'));
const CanonicalList = React.lazy(() => import('./views/Canonical/Canonical'));
const CanonicalForm = React.lazy(() => import('./views/Canonical/CanonicalForm'));
const CanonicalEditForm = React.lazy(() => import('./views/Canonical/CanonicalEditForm'));
const TitletagList = React.lazy(() => import('./views/Titletag/Titletag'));
const TitletagForm = React.lazy(() => import('./views/Titletag/TitletagForm'));
const TitletagEditForm = React.lazy(() => import('./views/Titletag/TitletagEditForm'));
const ScripttagList = React.lazy(() => import('./views/Scripttag/Scripttag'));
const ScripttagForm = React.lazy(() => import('./views/Scripttag/ScripttagForm'));
const ScripttagEditForm = React.lazy(() => import('./views/Scripttag/ScripttagEditForm'));
const LanguageList = React.lazy(() => import('./views/Language/Language'));
const LanguageForm = React.lazy(() => import('./views/Language/LanguageForm'));

const routes = [
  { path: '/', exact: true, name: 'Home' },
  { path: '/dashboard', name: 'Dashboard', component: Dashboard },
  { path: '/buttons', exact: true, name: 'Buttons', component: Buttons },
  { path: '/buttons/buttons', name: 'Buttons', component: Buttons },
  { path: '/buttons/button-dropdowns', name: 'Button Dropdowns', component: ButtonDropdowns },
  { path: '/buttons/button-groups', name: 'Button Groups', component: ButtonGroups },
  { path: '/buttons/brand-buttons', name: 'Brand Buttons', component: BrandButtons },
  { path: '/icons', exact: true, name: 'Icons', component: CoreUIIcons },
  { path: '/icons/coreui-icons', name: 'CoreUI Icons', component: CoreUIIcons },
  { path: '/icons/simple-line-icons', name: 'Simple Line Icons', component: SimpleLineIcons },
  { path: '/users', exact: true,  name: 'Users', component: Users },
  { path: '/users/:id', exact: true, name: 'User Details', component: User },
  { path: '/rule', name: 'Rules', component: RuleList },
  { path: '/ruleForm', name: 'Rules', component: RuleForm },
  { path: '/ruleDetail/', exact: true, name: 'Rule Details', component: RuleDetail },
  { path: '/metatag', name: 'Metatag', component: MetaTagList },
  { path: '/metatagForm', name: 'Metatag', component: MetatagForm },
  { path: '/metatagEditForm', name: 'Metatag', component: MetaTagEditForm },
  { path: '/metatagPreview', name: 'MetatagPreview', component: MetatagPreview },
  { path: '/datasource', name: 'DataSource', component: DataSourceList },
  { path: '/dataSourceForm', name: 'DataSource', component: DataSourceForm },
  { path: '/canonical', name: 'Canonical', component: CanonicalList },
  { path: '/canonicalForm', name: 'Canonical', component: CanonicalForm },
  { path: '/canonicalEditForm', name: 'Canonical', component: CanonicalEditForm },
  { path: '/titletag', name: 'Titletag', component: TitletagList },
  { path: '/titletagForm', name: 'Titletag', component: TitletagForm },
  { path: '/titletagEditForm', name: 'Titletag', component: TitletagEditForm },
  { path: '/scripttag', name: 'Scripttag', component: ScripttagList },
  { path: '/scripttagForm', name: 'Scripttag', component: ScripttagForm },
  { path: '/scripttagEditForm', name: 'Scripttag', component: ScripttagEditForm },
  { path: '/language', name: 'Language', component: LanguageList },
  { path: '/languageForm', name: 'Language', component: LanguageForm },
];

export default routes;