import React, { useState, useEffect } from "react";
import { Empty, Table, Space, Button, Modal, Input } from "antd";
import { fetchSettings, updateSetting } from "api/settings";

function ViewSettings() {
  const [settings, setSettings] = useState([]);
  const [modal, setModal] = useState(false);
  const [data, setData] = useState({ key: "", value: "" });

  useEffect(() => {
    fetchSettings().then((settings) => {
      setSettings(settings);
    });
  }, [modal]);

  const handleOk = (e) => {
    updateSetting(data.key, data)
      .then(() => {
        setModal(false);
      })
      .catch((error) => {
        console.log(error);
      });
  };

  const columns = [
    { title: "Key", dataIndex: "key", key: "key" },
    { title: "Value", dataIndex: "value", key: "value" },
    {
      title: "Action",
      key: "action",
      render: (text, data) => {
        return (
          <Space size="middle">
            <Button
              type="primary"
              onClick={(e) => {
                setModal(true);
                setData(data);
              }}
            >
              Edit
            </Button>
          </Space>
        );
      },
    },
  ];

  if (settings.length < 1) {
    return <Empty image={Empty.PRESENTED_IMAGE_SIMPLE} />;
  }

  return (
    <>
      <Table
        dataSource={settings}
        columns={columns}
        pagination={{
          total: settings.length,
          pageSize: settings.length,
          hideOnSinglePage: true,
        }}
      />
      <Modal
        title="Basic Modal"
        visible={modal}
        onOk={handleOk}
        onCancel={() => setModal(false)}
      >
        <Space>
          {data.key}
          <Input
            value={data.value}
            onChange={(e) => setData({ key: data.key, value: e.target.value })}
            onFocus={(e) => e.target.select()}
          />
        </Space>
      </Modal>
    </>
  );
}

export default ViewSettings;
