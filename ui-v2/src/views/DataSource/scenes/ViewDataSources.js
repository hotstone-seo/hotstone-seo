import React from 'react';
import PropTypes from 'prop-types';
import { Link, useHistory } from 'react-router-dom';
import { PageHeader, Button, message } from 'antd';
import { deleteDataSource } from 'api/datasource';
import useDataSources from 'hooks/useDataSources';
import { DataSourceList } from 'components/DataSource';

import { PlusCircleOutlined } from '@ant-design/icons';

function ViewDataSources({ match }) {
  const [dataSources, setDataSources] = useDataSources();
  const history = useHistory();

  const showEditScene = (dataSource) => {
    history.push(`${match.url}/${dataSource.id}`);
  };

  const removeDataSource = (dataSource) => {
    deleteDataSource(dataSource.id)
      .then(() => {
        message.success(`Successfully deleted ${dataSource.name}`);
        setDataSources(dataSources.filter((item) => item.id !== dataSource.id));
      })
      .catch((error) => {
        message.error(error.message);
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
        <Button type="primary" style={{ marginBottom: 16 }} icon={<PlusCircleOutlined />}  onClick={() => addDataSource()}>
           Add New Data Source
        </Button>
        <DataSourceList
          dataSources={dataSources}
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
