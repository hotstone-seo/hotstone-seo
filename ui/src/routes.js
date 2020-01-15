import React from "react";
import MetatagForm from "./views/Metatag/MetatagForm";

const Dashboard = React.lazy(() => import("./views/Dashboard"));
const RuleList = React.lazy(() => import("./views/Rule/RuleList"));
const RuleForm = React.lazy(() => import("./views/Rule/RuleForm"));
const RuleDetail = React.lazy(() => import("./views/Rule/RuleDetail"));
const MetatagPreview = React.lazy(() =>
  import("./views/Metatag/MetatagPreview")
);

const DataSourceList = React.lazy(() =>
  import("./views/DataSource/DataSource")
);
const DataSourceForm = React.lazy(() =>
  import("./views/DataSource/DataSourceForm")
);
const CanonicalForm = React.lazy(() =>
  import("./views/Canonical/CanonicalForm")
);
const TitletagForm = React.lazy(() => import("./views/Titletag/TitletagForm"));
const TitletagEditForm = React.lazy(() =>
  import("./views/Titletag/TitletagEditForm")
);
const ScripttagForm = React.lazy(() =>
  import("./views/Scripttag/ScripttagForm")
);
const ScripttagEditForm = React.lazy(() =>
  import("./views/Scripttag/ScripttagEditForm")
);
const LanguageList = React.lazy(() => import("./views/Language/Language"));
const LanguageForm = React.lazy(() => import("./views/Language/LanguageForm"));

const MismatchRuleList = React.lazy(() =>
  import("./views/MismatchRule/MismatchRuleList")
);

const AnalyticPage = React.lazy(() => import("./views/Analytic/AnalyticPage"));

const routes = [
  { path: "/", exact: true, name: "Home" },
  { path: "/dashboard", name: "Dashboard", component: Dashboard },
  { path: "/rule", name: "Rules", component: RuleList },
  { path: "/ruleForm", name: "Rules", component: RuleForm },
  {
    path: "/ruleDetail",
    name: "Rule Details",
    component: RuleDetail
  },
  { path: "/metatagForm", name: "Metatag", component: MetatagForm },
  {
    path: "/metatagPreview",
    name: "MetatagPreview",
    component: MetatagPreview
  },
  { path: "/datasource", name: "DataSource", component: DataSourceList },
  { path: "/dataSourceForm", name: "DataSource", component: DataSourceForm },
  { path: "/canonicalForm", name: "Canonical", component: CanonicalForm },
  { path: "/titletagForm", name: "Titletag", component: TitletagForm },
  { path: "/titletagEditForm", name: "Titletag", component: TitletagEditForm },
  { path: "/scripttagForm", name: "Scripttag", component: ScripttagForm },
  {
    path: "/scripttagEditForm",
    name: "Scripttag",
    component: ScripttagEditForm
  },
  { path: "/language", name: "Language", component: LanguageList },
  { path: "/languageForm", name: "Language", component: LanguageForm },
  { path: "/mismatchRule", name: "MismatchRule", component: MismatchRuleList },
  { path: "/analytic", name: "Analytic", component: AnalyticPage }
];

export default routes;
