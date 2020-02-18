import React from "react";

const Rule = React.lazy(() => import("./views/Rule"));
const MismatchRule = React.lazy(() => import("./views/MismatchRule"));

const routes = [
  { path: "/", exact: true, name: "Home" },
  { path: "/rules", name: "Rules", component: Rule },
  { path: "/mismatch-rule", name: "MismatchRule", component: MismatchRule }
];

export default routes;
