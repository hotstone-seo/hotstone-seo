import React, { useState, useEffect } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import {
  PageHeader, Row, Col, message,
} from 'antd';
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
  }, [id]);

  const handleEdit = (newDataSource) => {
    updateDataSource(newDataSource)
      .then(() => {
        history.push('/datasources');
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  return (
    <div>
      <PageHeader
        onBack={() => history.push('/datasources')}
        title={`Edit ${dataSource.name}`}
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <Row>
          <Col span={12} style={{ background: '#fff', paddingTop: 24 }}>
            <DataSourceForm handleSubmit={handleEdit} dataSource={dataSource} />
          </Col>
        </Row>
      </div>
    </div>
  );
}

export default EditDataSource;
