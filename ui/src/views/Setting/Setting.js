import React, { useState, useEffect } from "react";
import { PageHeader, Card, Empty, Table, Space } from "antd";
import { fetchSettings } from "api/settings";

function Setting({ match }) {
  const [settings, setSettings] = useState([]);

  useEffect(() => {
    fetchSettings().then((settings) => {
      setSettings(settings);
    });
  }, []);

  let settingsView;
  if (settings.length < 1) {
    settingsView = <Empty image={Empty.PRESENTED_IMAGE_SIMPLE} />;
  } else {
    settingsView = renderSettings(settings);
  }

  return (
    <div className="Setting">
      <PageHeader
        title="Setting"
        subTitle="Application Setting"
        style={{ background: "#fff" }}
      />

      <div style={{ padding: 24 }}>
        <Card>{settingsView}</Card>
      </div>
    </div>
  );
}

function renderSettings(settings) {
  const columns = [
    { title: "Key", dataIndex: "key", key: "key" },
    { title: "Value", dataIndex: "value", key: "value" },
    {
      title: "Action",
      key: "action",
      render: (text, record) => (
        <Space size="middle">
          <a href="#">Edit</a>
        </Space>
      ),
    },
  ];
  return (
    <Table
      dataSource={settings}
      columns={columns}
      pagination={{
        total: settings.length,
        pageSize: settings.length,
        hideOnSinglePage: true,
      }}
    />
  );
}

export default Setting;
