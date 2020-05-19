import React, { useState } from 'react';
import PropTypes from 'prop-types';
import { useHistory } from 'react-router-dom';
import { PageHeader, Button, message } from 'antd';
import { deleteModule } from 'api/module';
import { ModuleList } from 'components/Module';
import { PlusOutlined } from '@ant-design/icons';

function ViewModules({ match }) {
  const history = useHistory();
  const [listModule, setListModule] = useState([]);

  const showEditForm = (module) => {
    history.push(`${match.url}/${module.id}`);
  };

  const handleDelete = (module) => {
    deleteModule(module.id)
      .then(() => {
        message.success(`Successfully deleted ${module.name}`);
        setListModule(listModule.filter((item) => item.id !== module.id));
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  const addDataModule = () => {
    history.push(`${match.url}/new`);
  };

  return (
    <div>
      <PageHeader
        title="Modules"
        subTitle="List of Modules"
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <Button
          data-testid="btn-new-rule"
          type="primary"
          style={{ marginBottom: 16 }}
          icon={<PlusOutlined />}
          onClick={() => addDataModule()}
        >
          Add New Module
        </Button>
        <ModuleList
          onClick={showEditForm}
          onEdit={showEditForm}
          onDelete={handleDelete}
          listModule={listModule}
          setListModule={setListModule}
        />
      </div>
    </div>
  );
}

ViewModules.propTypes = {
  match: PropTypes.shape({
    url: PropTypes.string,
  }).isRequired,
};

export default ViewModules;
