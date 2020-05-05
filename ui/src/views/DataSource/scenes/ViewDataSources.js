import React from 'react';
import PropTypes from 'prop-types';
import { useHistory } from 'react-router-dom';
import { PageHeader, Button, message } from 'antd';
import { fetchDataSources, deleteDataSource } from 'api/datasource';
import useAsync from 'hooks/useAsync';
import { DataSourceList } from 'components/DataSource';

import { PlusOutlined } from '@ant-design/icons';

function ViewDataSources({ match }) {
  const history = useHistory();
  const {
    pending, value: dataSources, setValue: setDataSources, error,
  } = useAsync(fetchDataSources);

  if (error) {
    message.error(error.message);
  }

  const showEditScene = (dataSource) => {
    history.push(`${match.url}/${dataSource.id}`);
  };

  const removeDataSource = (dataSource) => {
    deleteDataSource(dataSource.id)
      .then(() => {
        message.success(`Successfully deleted ${dataSource.name}`);
        setDataSources(
          dataSources.filter((item) => item.id !== dataSource.id),
        );
      })
      .catch((err) => {
        message.error(err.message);
      });
  };

  const addDataSource = () => {
    history.push(`${match.url}/new`);
  };

  return (
    <div>
      <PageHeader
        title="Data Sources"
        subTitle="Manage location for retrieving resources"
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <Button
          type="primary"
          style={{ marginBottom: 16 }}
          icon={<PlusOutlined />}
          onClick={() => addDataSource()}
        >
          Add New Data Source
        </Button>
        <DataSourceList
          dataSources={dataSources}
          loading={pending}
          onClick={showEditScene}
          onEdit={showEditScene}
          onDelete={removeDataSource}
        />
      </div>
    </div>
  );
}

ViewDataSources.propTypes = {
  match: PropTypes.shape({
    url: PropTypes.string,
  }).isRequired,
};

export default ViewDataSources;
