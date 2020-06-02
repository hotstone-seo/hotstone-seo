import React from "react";
import { PageHeader, Card } from "antd";
import ViewSettings from "./ViewSettings";

function Setting({ match }) {
  return (
    <div className="Setting">
      <PageHeader
        title="Setting"
        subTitle="Application Setting"
        style={{ background: "#fff" }}
      />

      <div style={{ padding: 24 }}>
        <Card>
          <ViewSettings />
        </Card>
      </div>
    </div>
  );
}

export default Setting;
