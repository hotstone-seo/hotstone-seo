import React, { useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import { Descriptions, message } from 'antd';
import { getDataSource } from 'api/datasource';

function RuleDetail({ rule }) {
  const { name, url_pattern: urlPattern, data_source_id: dataSourceID } = rule;
  const [dataSource, setDataSource] = useState(null);

  useEffect(() => {
    if (dataSourceID) {
      getDataSource(dataSourceID)
        .then((retrievedValue) => {
          setDataSource(retrievedValue);
        })
        .catch((error) => {
          message.error(error.message);
        });
    }
  }, [dataSourceID]);

  return (
    <Descriptions>
      <Descriptions.Item key="name" label="Name">{name}</Descriptions.Item>
      <Descriptions.Item key="urlPattern" label="URL Pattern">{urlPattern}</Descriptions.Item>
      {dataSource && (
        <Descriptions.Item key="dataSource" label="Data Source">
          <Link to={`/datasources/${dataSource.id}`}>
            {dataSource.name}
          </Link>
        </Descriptions.Item>
      )}
    </Descriptions>
  );
}

RuleDetail.propTypes = {
  rule: PropTypes.shape({
    name: PropTypes.string,
    url_pattern: PropTypes.string,
    data_source_id: PropTypes.number,
  }).isRequired,
};

export default RuleDetail;
