import React from 'react';
import PropTypes from 'prop-types';
import { Link, useHistory } from 'react-router-dom';
import { Button, message } from 'antd';
import { deleteDataSource } from 'api/datasource';
import useDataSources from 'hooks/useDataSources';
import { DataSourceList } from 'components/DataSource';

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

  return (
    <div>
      <Button type="primary" style={{ marginBottom: 16 }}>
        <Link to={`${match.url}/new`}>Add new Data Source</Link>
      </Button>
      <DataSourceList
        dataSources={dataSources}
        onClick={showEditScene}
        onEdit={showEditScene}
        onDelete={removeDataSource}
      />
    </div>
  );
}

ViewDataSources.propTypes = {
  match: PropTypes.shape({
    url: PropTypes.string,
  }).isRequired,
};

export default ViewDataSources;
