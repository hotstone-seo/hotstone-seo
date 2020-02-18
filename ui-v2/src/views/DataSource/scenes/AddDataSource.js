import React from 'react';
import { useHistory } from 'react-router-dom';
import { message } from 'antd';
import { DataSourceForm } from 'components/DataSource';
import { createDataSource } from 'api/datasource';

function AddDataSource() {
  const history = useHistory();

  const handleCreate = (dataSource) => {
    createDataSource(dataSource)
      .then(() => {
        history.push('/datasources');
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  return (
    <DataSourceForm handleSubmit={handleCreate} />
  );
}

export default AddDataSource;
