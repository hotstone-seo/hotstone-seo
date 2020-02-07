import React from "react";

const RuleList = React.lazy(() => import("./views/Rule/RuleList"));
const RuleDetail = React.lazy(() => import("./views/Rule/RuleDetail"));
const DataSourceList = React.lazy(() =>
  import("./views/DataSource/DataSource")
);
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
    path: "/rule-detail",
    name: "Rule Details",
    component: RuleDetail
  },
  { path: "/datasource", name: "DataSource", component: DataSourceList },
  { path: "/mismatchRule", name: "MismatchRule", component: MismatchRuleList },
  { path: "/analytic", name: "Analytic", component: AnalyticPage },
  { path: "/simulation", name: "Simulation", component: SimulationPage }
];

export default routes;
