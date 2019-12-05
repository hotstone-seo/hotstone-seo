import React from 'react';
import MetatagForm from './views/Base/Metatag/MetatagForm';
 
const BrandButtons = React.lazy(() => import('./views/Buttons/BrandButtons'));
const ButtonDropdowns = React.lazy(() => import('./views/Buttons/ButtonDropdowns'));
const ButtonGroups = React.lazy(() => import('./views/Buttons/ButtonGroups'));
const Buttons = React.lazy(() => import('./views/Buttons/Buttons'));
const Dashboard = React.lazy(() => import('./views/Dashboard'));
const CoreUIIcons = React.lazy(() => import('./views/Icons/CoreUIIcons'));
const Flags = React.lazy(() => import('./views/Icons/Flags'));
const FontAwesome = React.lazy(() => import('./views/Icons/FontAwesome'));
const SimpleLineIcons = React.lazy(() => import('./views/Icons/SimpleLineIcons'));
const Users = React.lazy(() => import('./views/Users/Users'));
const User = React.lazy(() => import('./views/Users/User'));
const RuleList = React.lazy(() => import('./views/Base/Rule/RuleList'));
const RuleForm = React.lazy(() => import('./views/Base/Rule/RuleForm'));
const RuleEditForm = React.lazy(() => import('./views/Base/Rule/RuleEditForm'));
const MetaTagList = React.lazy(() => import('./views/Base/Metatag/Metatag'));
const MetatagPreview = React.lazy(() => import('./views/Base/Metatag/MetatagPreview'));
const MetaTagEditForm = React.lazy(() => import('./views/Base/Metatag/MetatagEditForm'));
const DataSourceList = React.lazy(() => import('./views/Base/DataSource/DataSource'));
const DataSourceEditForm = React.lazy(() => import('./views/Base/DataSource/DataSourceEditForm'));
const DataSourceForm = React.lazy(() => import('./views/Base/DataSource/DataSourceForm'));
const CanonicalList = React.lazy(() => import('./views/Canonical/Canonical'));
const CanonicalForm = React.lazy(() => import('./views/Canonical/CanonicalForm'));
const CanonicalEditForm = React.lazy(() => import('./views/Canonical/CanonicalEditForm'));
const TitletagList = React.lazy(() => import('./views/Base/Titletag/Titletag'));
const TitletagForm = React.lazy(() => import('./views/Base/Titletag/TitletagForm'));
const TitletagEditForm = React.lazy(() => import('./views/Base/Titletag/TitletagEditForm'));
const ScripttagList = React.lazy(() => import('./views/Base/Scripttag/Scripttag'));
const ScripttagForm = React.lazy(() => import('./views/Base/Scripttag/ScripttagForm'));
const ScripttagEditForm = React.lazy(() => import('./views/Base/Scripttag/ScripttagEditForm'));
const LanguageList = React.lazy(() => import('./views/Base/Language/Language'));
const LanguageForm = React.lazy(() => import('./views/Base/Language/LanguageForm'));
const LanguageEditForm = React.lazy(() => import('./views/Base/Language/LanguageEditForm'));

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
  { path: '/icons/flags', name: 'Flags', component: Flags },
  { path: '/icons/font-awesome', name: 'Font Awesome', component: FontAwesome },
  { path: '/icons/simple-line-icons', name: 'Simple Line Icons', component: SimpleLineIcons },
  { path: '/users', exact: true,  name: 'Users', component: Users },
  { path: '/users/:id', exact: true, name: 'User Details', component: User },
  { path: '/base/rule', name: 'Rules', component: RuleList },
  { path: '/base/ruleForm', name: 'Rules', component: RuleForm },
  { path: '/base/ruleEditForm', name: 'Rules', component: RuleEditForm },
  { path: '/base/metatag', name: 'Metatag', component: MetaTagList },
  { path: '/base/metatagForm', name: 'Metatag', component: MetatagForm },
  { path: '/base/metatagEditForm', name: 'Metatag', component: MetaTagEditForm },
  { path: '/base/metatagPreview', name: 'MetatagPreview', component: MetatagPreview },
  { path: '/base/datasource', name: 'DataSource', component: DataSourceList },
  { path: '/base/DataSourceForm', name: 'DataSource', component: DataSourceForm },
  { path: '/base/DataSourceEditForm', name: 'DataSource', component: DataSourceEditForm },
  { path: '/canonical', name: 'Canonical', component: CanonicalList },
  { path: '/canonicalForm', name: 'Canonical', component: CanonicalForm },
  { path: '/canonicalEditForm', name: 'Canonical', component: CanonicalEditForm },
  { path: '/base/titletag', name: 'Titletag', component: TitletagList },
  { path: '/base/titletagForm', name: 'Titletag', component: TitletagForm },
  { path: '/base/titletagEditForm', name: 'Titletag', component: TitletagEditForm },
  { path: '/base/scripttag', name: 'Scripttag', component: ScripttagList },
  { path: '/base/scripttagForm', name: 'Scripttag', component: ScripttagForm },
  { path: '/base/scripttagEditForm', name: 'Scripttag', component: ScripttagEditForm },
  { path: '/base/Language', name: 'Language', component: LanguageList },
  { path: '/base/LanguageForm', name: 'Language', component: LanguageForm },
  { path: '/base/LanguageEditForm', name: 'Language', component: LanguageEditForm },
];

export default routes;
