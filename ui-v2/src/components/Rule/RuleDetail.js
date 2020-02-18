import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import { Descriptions } from 'antd';

function RuleDetail({ rule }) {
  const { name, url_pattern: urlPattern, data_source_id: dataSourceID } = rule;
  return (
    <Descriptions>
      <Descriptions.Item label="Name">{name}</Descriptions.Item>
      <Descriptions.Item label="URL Pattern">{urlPattern}</Descriptions.Item>
      {dataSourceID && (
        <Descriptions.Item label="Data Source">
          <Link to={`/data_sources/${dataSourceID}`}>
            Data Source
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
