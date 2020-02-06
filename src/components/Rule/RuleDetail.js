import React from 'react';
import { Link } from 'react-router-dom';
import { Descriptions } from 'antd';

function RuleDetail({ rule }) {
  const { name, url_pattern: urlPattern, data_source_id } = rule;
  return (
    <Descriptions title="Rule Detail">
      <Descriptions.Item label="Name">{name}</Descriptions.Item>
      <Descriptions.Item label="URL Pattern">{urlPattern}</Descriptions.Item>
      {data_source_id && (
        <Descriptions.Item label="Data Source">
          <Link to={`/data_sources/${data_source_id}`}>
            Data Source
          </Link>
        </Descriptions.Item>
      )}
    </Descriptions>
  );
}

export default RuleDetail;
