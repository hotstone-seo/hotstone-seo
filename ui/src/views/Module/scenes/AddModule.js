import React from 'react';
import { useHistory } from 'react-router-dom';
import {
  PageHeader, Row, Col, message,
} from 'antd';
import { ModuleForm } from 'components/Module';

function AddModule() {
  const history = useHistory();

  const handleCreate = (module) => {
    if (module.name === undefined) {
      message.error(`Module ${module.name} already register`);
    } else {
      history.push('/modules', {
        message: {
          level: 'success',
          content: `${module.name} is successfully created`,
        },
      });
    }
  };

  return (
    <div>
      <PageHeader
        onBack={() => history.push('/modules')}
        title="Add new Module"
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <Row>
          <Col span={12} style={{ background: '#fff', paddingTop: 24 }}>
            <ModuleForm handleSubmit={handleCreate} />
          </Col>
        </Row>
      </div>
    </div>
  );
}

export default AddModule;
