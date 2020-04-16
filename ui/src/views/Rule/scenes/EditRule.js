import React, { useState, useEffect } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import {
  PageHeader, Row, Col, message, Button,
} from 'antd';
import { EditOutlined, BarChartOutlined } from '@ant-design/icons';
import { RuleForm, RuleDetail } from 'components/Rule';
import { getRule, updateRule } from 'api/rule';
import useDataSources from 'hooks/useDataSources';
import ManageTags from './sections/ManageTags';


function EditRule() {
  const { id } = useParams();
  const history = useHistory();
  const [dataSources] = useDataSources();

  const [rule, setRule] = useState({});
  const [isEditingRule, setIsEditingRule] = useState(false);

  useEffect(() => {
    getRule(id)
      .then((newRule) => {
        setRule(newRule);
      })
      .catch((error) => {
        history.push('/rules', {
          message: {
            level: 'error',
            content: error.message,
          },
        });
      });
  }, [id, history]);

  const editRule = (newRule) => {
    updateRule(newRule)
      .then(() => {
        setRule(newRule);
        setIsEditingRule(false);
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  return (
    <div>
      <PageHeader
        onBack={() => history.push('/rules')}
        title="Manage Rule"
        subTitle="Organize tags to be rendered"
        style={{ background: '#fff', minHeight: 120 }}
        extra={[
          <Button
            data-testid="btn-edit"
            key="edit"
            type={isEditingRule ? 'default' : 'primary'}
            onClick={() => setIsEditingRule(!isEditingRule)}
            icon={<EditOutlined />}
          >
            {isEditingRule ? 'Cancel' : 'Edit Rule'}
          </Button>,
          <Button
            key="analytics"
            icon={<BarChartOutlined />}
            onClick={() => {
              history.push(`/analytic?ruleID=${rule.id}`);
            }}
          >
            Analytics
          </Button>,
        ]}
      >
        {isEditingRule ? (
          <RuleForm
            rule={rule}
            dataSources={dataSources}
            formLayout="inline"
            onSubmit={editRule}
          />
        ) : (
          <RuleDetail rule={rule} />
        )}
      </PageHeader>

      <div style={{ padding: 24 }}>
        <ManageTags ruleID={parseInt(id, 10)} />
      </div>
    </div>
  );
}

export default EditRule;
