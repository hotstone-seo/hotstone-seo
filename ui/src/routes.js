import React from "react";

const RuleList = React.lazy(() => import("./views/Rule/RuleList"));
const RuleDetail = React.lazy(() => import("./views/Rule/RuleDetail"));
const MetatagPreview = React.lazy(() =>
  import("./views/Metatag/MetatagPreview")
);
const DataSourceList = React.lazy(() =>
  import("./views/DataSource/DataSource")
);
const LanguageList = React.lazy(() => import("./views/Language/Language"));
const MismatchRuleList = React.lazy(() =>
  import("./views/MismatchRule/MismatchRuleList")
);
const AnalyticPage = React.lazy(() => import("./views/Analytic/AnalyticPage"));
const SimulationPage = React.lazy(() =>
  import("./views/Simulation/SimulationPage")
);

const routes = [
  { path: "/", exact: true, name: "Home" },
  { path: "/rule", name: "Rules", component: RuleList },
  {
    path: "/ruleDetail",
    name: "Rule Details",
    component: RuleDetail
  },
  {
    path: "/metatagPreview",
    name: "MetatagPreview",
    component: MetatagPreview
  },
  { path: "/datasource", name: "DataSource", component: DataSourceList },
  { path: "/language", name: "Language", component: LanguageList },
  { path: "/mismatchRule", name: "MismatchRule", component: MismatchRuleList },
  { path: "/analytic", name: "Analytic", component: AnalyticPage },
  { path: "/simulation", name: "Simulation", component: SimulationPage }
];

export default routes;
