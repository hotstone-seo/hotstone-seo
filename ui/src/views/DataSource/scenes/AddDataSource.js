import React from 'react';
import { useHistory } from 'react-router-dom';
import {
  PageHeader, Row, Col, message,
} from 'antd';
import { DataSourceForm } from 'components/DataSource';
import { createDataSource } from 'api/datasource';

function AddDataSource() {
  const history = useHistory();

  const handleCreate = (dataSource) => {
    createDataSource(dataSource)
      .then((newDataSource) => {
        history.push('/datasources', {
          message: {
            level: 'success',
            content: `${newDataSource.name} is successfully created`,
          },
        });
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  return (
    <div>
      <PageHeader
        onBack={() => history.push('/datasources')}
        title="Add new Data Source"
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <Row>
          <Col span={12} style={{ background: '#fff', paddingTop: 24 }}>
            <DataSourceForm handleSubmit={handleCreate} />
          </Col>
        </Row>
      </div>
    </div>
  );
}

export default AddDataSource;
