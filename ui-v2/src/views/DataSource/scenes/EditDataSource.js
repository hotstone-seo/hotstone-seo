import React, { useState, useEffect } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import { message } from 'antd';
import { DataSourceForm } from 'components/DataSource';
import { getDataSource, updateDataSource } from 'api/datasource';

function EditDataSource() {
  const { id } = useParams();
  const history = useHistory();

  const [dataSource, setDataSource] = useState({});

  useEffect(() => {
    getDataSource(id)
      .then((newDataSource) => {
        setDataSource(newDataSource);
      })
      .catch((error) => {
        message.error(error.message);
      });
  });

  const handleCreate = (newDataSource) => {
    updateDataSource(newDataSource)
      .then(() => {
        history.push('/datasources');
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  return (
    <DataSourceForm handleSubmit={handleCreate} dataSource={dataSource} />
  );
}

export default EditDataSource;
