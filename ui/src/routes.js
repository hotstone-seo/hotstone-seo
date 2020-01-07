import React from "react";
import MetatagForm from "./views/Metatag/MetatagForm";

const Dashboard = React.lazy(() => import("./views/Dashboard"));
const RuleList = React.lazy(() => import("./views/Rule/RuleList"));
const RuleForm = React.lazy(() => import("./views/Rule/RuleForm"));
const RuleDetail = React.lazy(() => import("./views/Rule/RuleDetail"));
const MetaTagList = React.lazy(() => import("./views/Metatag/Metatag"));
const MetatagPreview = React.lazy(() =>
  import("./views/Metatag/MetatagPreview")
);
const MetaTagEditForm = React.lazy(() =>
  import("./views/Metatag/MetatagEditForm")
);
const DataSourceList = React.lazy(() =>
  import("./views/DataSource/DataSource")
);
const DataSourceForm = React.lazy(() =>
  import("./views/DataSource/DataSourceForm")
);
const CanonicalList = React.lazy(() => import("./views/Canonical/Canonical"));
const CanonicalForm = React.lazy(() =>
  import("./views/Canonical/CanonicalForm")
);
const CanonicalEditForm = React.lazy(() =>
  import("./views/Canonical/CanonicalEditForm")
);
const TitletagList = React.lazy(() => import("./views/Titletag/Titletag"));
const TitletagForm = React.lazy(() => import("./views/Titletag/TitletagForm"));
const TitletagEditForm = React.lazy(() =>
  import("./views/Titletag/TitletagEditForm")
);
const ScripttagList = React.lazy(() => import("./views/Scripttag/Scripttag"));
const ScripttagForm = React.lazy(() =>
  import("./views/Scripttag/ScripttagForm")
);
const ScripttagEditForm = React.lazy(() =>
  import("./views/Scripttag/ScripttagEditForm")
);
const LanguageList = React.lazy(() => import("./views/Language/Language"));
const LanguageForm = React.lazy(() => import("./views/Language/LanguageForm"));

const MismatchRuleList = React.lazy(() => import("./views/MismatchRule/MismatchRuleList"));

const routes = [
  { path: "/", exact: true, name: "Home" },
  { path: "/dashboard", name: "Dashboard", component: Dashboard },
  { path: "/rule", name: "Rules", component: RuleList },
  { path: "/ruleForm", name: "Rules", component: RuleForm },
  {
    path: "/ruleDetail?ruleId=:id",
    exact: true,
    name: "Rule Details",
    component: RuleDetail
  },
  { path: "/metatag", name: "Metatag", component: MetaTagList },
  { path: "/metatagForm", name: "Metatag", component: MetatagForm },
  { path: "/metatagEditForm", name: "Metatag", component: MetaTagEditForm },
  {
    path: "/metatagPreview",
    name: "MetatagPreview",
    component: MetatagPreview
  },
  { path: "/datasource", name: "DataSource", component: DataSourceList },
  { path: "/dataSourceForm", name: "DataSource", component: DataSourceForm },
  { path: "/canonical", name: "Canonical", component: CanonicalList },
  { path: "/canonicalForm", name: "Canonical", component: CanonicalForm },
  {
    path: "/canonicalEditForm",
    name: "Canonical",
    component: CanonicalEditForm
  },
  { path: "/titletag", name: "Titletag", component: TitletagList },
  { path: "/titletagForm", name: "Titletag", component: TitletagForm },
  { path: "/titletagEditForm", name: "Titletag", component: TitletagEditForm },
  { path: "/scripttag", name: "Scripttag", component: ScripttagList },
  { path: "/scripttagForm", name: "Scripttag", component: ScripttagForm },
  {
    path: "/scripttagEditForm",
    name: "Scripttag",
    component: ScripttagEditForm
  },
  { path: "/language", name: "Language", component: LanguageList },
  { path: "/languageForm", name: "Language", component: LanguageForm },
  { path: "/mismatchRule", name: "MismatchRule", component: MismatchRuleList }
];

export default routes;
