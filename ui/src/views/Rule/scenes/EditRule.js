import React, { useState, useEffect } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import {
  PageHeader,
  Row,
  Col,
  message,
  Select,
  Button,
} from 'antd';
import { EditOutlined, PlusOutlined, BarChartOutlined } from '@ant-design/icons';
import { RuleForm, RuleDetail } from 'components/Rule';
import { TagList, TagForm } from 'components/Tag';
import { getRule, updateRule } from 'api/rule';
import { fetchTags, deleteTag } from 'api/tag';
import useDataSources from 'hooks/useDataSources';
import locales from 'locales';

const { Option } = Select;

function EditRule() {
  const { id } = useParams();
  const history = useHistory();
  const [dataSources] = useDataSources();

  const [rule, setRule] = useState({});
  const [tags, setTags] = useState([]);
  const [currentTag, setCurrentTag] = useState(null);
  const [locale, setLocale] = useState(locales[0] || '');
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

  useEffect(() => {
    fetchTags({ rule_id: id, locale })
      .then((newTags) => {
        setTags(newTags);
      })
      .catch((error) => {
        message.error(error.message);
      });
  }, [id, locale]);

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

  const submitTag = () => {
    setCurrentTag(null);
    fetchTags({ rule_id: id, locale })
      .then((newTags) => {
        setTags(newTags);
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  const addTag = () => {
    setCurrentTag({ rule_id: parseInt(id, 10), locale });
  };

  const cancelTag = () => {
    setCurrentTag(null);
  };

  const editTag = (tag) => {
    setCurrentTag(tag);
  };

  const removeTag = (tag) => {
    deleteTag(tag.id)
      .then(() => {
        setTags(tags.filter((item) => item.id !== tag.id));
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
        <Row>
          <Col span={24} style={{ background: '#fff', padding: 24 }}>
            {currentTag ? (
              <TagForm tag={currentTag} onSubmit={submitTag} onCancel={cancelTag} />
            ) : (
              <>
                <Select
                  defaultValue={locale}
                  onChange={(value) => setLocale(value)}
                  style={{ float: 'right', marginBottom: 16 }}
                >
                  {locales.map((loc) => (
                    <Option value={loc} key={loc}>
                      {loc}
                    </Option>
                  ))}
                </Select>
                <Button
                  data-testid="btn-new-tag"
                  type="dashed"
                  onClick={addTag}
                  style={{ width: '100%', marginBottom: 16 }}
                >
                  <PlusOutlined />
                  Add Tag
                </Button>
                <TagList tags={tags} onEdit={editTag} onDelete={removeTag} />
              </>
            )}
          </Col>
        </Row>
      </div>
    </div>
  );
}

export default EditRule;
