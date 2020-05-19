import React from 'react';
import { useHistory } from 'react-router-dom';
import {
  PageHeader, Row, Col, message,
} from 'antd';
import { ModuleForm } from 'components/Module';
import { createModule } from 'api/module';

function AddModule() {
  const history = useHistory();

  const handleCreate = (module) => {
    createModule(module)
      .then((newModule) => {
        history.push('/modules', {
          message: {
            level: 'success',
            content: `${newModule.name} is successfully created`,
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
