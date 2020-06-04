import React, { useState, useEffect } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import {
  PageHeader, Row, Col, message,
} from 'antd';
import { ModuleForm } from 'components/Module';
import { getModule, updateModule } from 'api/module';

function EditModule() {
  const { id } = useParams();
  const moduleID = parseInt(id, 10);
  const history = useHistory();

  const [module, setModule] = useState({});

  useEffect(() => {
    getModule(moduleID)
      .then((newModule) => {
        setModule(newModule);
      })
      .catch((error) => {
        history.push('/modules', {
          message: {
            level: 'error',
            content: error.message,
          },
        });
      });
  }, [moduleID, history]);

  const handleEdit = (newModule) => {
    history.push('/modules', {
      message: {
        level: 'success',
        content: `Module ${newModule.name} is successfully edit`,
      },
    });
  };

  return (
    <div>
      <PageHeader
        onBack={() => history.push('/modules')}
        title={`Edit Module ${module.name}`}
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <Row>
          <Col span={12} style={{ background: '#fff', paddingTop: 24 }}>
            <ModuleForm handleSubmit={handleEdit} module={module} />
          </Col>
        </Row>
      </div>
    </div>
  );
}

export default EditModule;
