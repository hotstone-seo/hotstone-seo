import React, { useState } from 'react';
import PropTypes from 'prop-types';
import { useHistory } from 'react-router-dom';
import { PageHeader, Button, message } from 'antd';
import { deleteRoleType } from 'api/roleType';
import { RoleTypeList } from 'components/RoleType';
import { PlusOutlined } from '@ant-design/icons';

function ViewRoleTypes({ match }) {
  const history = useHistory();
  const [listRoleType, setListRoleType] = useState([]);

  const showEditForm = (roleType) => {
    history.push(`${match.url}/${roleType.id}`);
  };

  const handleDelete = (roleType) => {
    deleteRoleType(roleType.id)
      .then(() => {
        message.success(`Successfully deleted ${roleType.name}`);
        setListRoleType(listRoleType.filter((item) => item.id !== roleType.id));
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  const addDataRoleType = () => {
    history.push(`${match.url}/new`);
  };

  return (
    <div>
      <PageHeader
        title="Role User"
        subTitle="List of Role User"
        style={{ background: '#fff' }}
      />
      <div style={{ padding: 24 }}>
        <Button
          data-testid="btn-new-rule"
          type="primary"
          style={{ marginBottom: 16 }}
          icon={<PlusOutlined />}
          onClick={() => addDataRoleType()}
        >
          Add New Role
        </Button>
        <RoleTypeList
          onClick={showEditForm}
          onEdit={showEditForm}
          onDelete={handleDelete}
          listRoleType={listRoleType}
          setListRoleType={setListRoleType}
        />
      </div>
    </div>
  );
}

ViewRoleTypes.propTypes = {
  match: PropTypes.shape({
    url: PropTypes.string,
  }).isRequired,
};

export default ViewRoleTypes;
