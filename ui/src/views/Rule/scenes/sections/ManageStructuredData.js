import React, { useState, useCallback } from 'react';
import PropTypes from 'prop-types';
import { message, Space, Button } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { fetchStructuredDatas, deleteStructuredData } from 'api/structuredData';
import useAsync from 'hooks/useAsync';
import { StructuredDataForm, StructuredDataList } from 'components/StructuredData';

// Aliasing for shorter name
const fetchStructs = fetchStructuredDatas;
const deleteStruct = deleteStructuredData;

function ManageStructuredData({ ruleID }) {
  // NOTE: Structs is the shorthand that we use for Structured Data
  const [focusStruct, setFocusStruct] = useState(null);
  const partialFetch = useCallback(
    () => fetchStructs({ rule_id: ruleID }),
    [ruleID],
  );
  const {
    value: structs, setValue: setStructs, pending, error, execute,
  } = useAsync(partialFetch);

  const refreshStruct = () => {
    execute().then(() => setFocusStruct(null));
  };

  const removeStruct = ({ id: structID }) => {
    deleteStruct(structID)
      .then(() => {
        setStructs(structs.filter((struct) => struct.id !== structID));
      })
      .catch((err) => {
        message.error(err.message);
      });
  };

  if (error) {
    message.error(error.message);
  }

  return (
    focusStruct ? (
      <StructuredDataForm
        structuredData={focusStruct}
        afterSubmit={refreshStruct}
        onCancel={() => setFocusStruct(null)}
      />
    ) : (
      <Space direction="vertical" style={{ width: '100%' }}>
        <Button
          type="dashed"
          onClick={() => setFocusStruct({ rule_id: ruleID })}
          style={{ width: '100%' }}
        >
          <PlusOutlined />
          Add Structured Data
        </Button>
        <StructuredDataList
          structuredDatas={structs}
          onEdit={(struct) => setFocusStruct(struct)}
          onDelete={removeStruct}
          loading={pending}
        />
      </Space>
    )
  );
}

ManageStructuredData.propTypes = {
  ruleID: PropTypes.number.isRequired,
};

export default ManageStructuredData;
