import React, { useState, useEffect } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import {
  PageHeader,
  Row,
  Col,
  message,
  Select,
  Button,
  Modal,
  Form,
} from 'antd';
import {
  EditOutlined,
  PlusOutlined,
  BarChartOutlined,
} from '@ant-design/icons';
import { RuleForm, RuleDetail } from 'components/Rule';
import { TagList, TagForm } from 'components/Tag';
import { getRule, updateRule } from 'api/rule';
import {
  fetchTags, createTag, updateTag, deleteTag,
} from 'api/tag';
import useDataSources from 'hooks/useDataSources';
import locales from 'locales';

const { Option } = Select;

function EditRule() {
  const { id } = useParams();
  const history = useHistory();
  const [dataSources] = useDataSources();
  const [tagForm] = Form.useForm();

  const [rule, setRule] = useState({});
  const [tags, setTags] = useState([]);
  const [locale, setLocale] = useState(locales[0] || '');
  const [isEditingRule, setIsEditingRule] = useState(false);
  const [tagFormTitle, setTagFormTitle] = useState('');
  const [tagFormVisible, setTagFormVisible] = useState(false);
  const [tagFormLoading, setTagFormLoading] = useState(false);

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
    tagForm
      .validateFields()
      .then((tag) => {
        setTagFormLoading(true);
        let submitFunc = createTag;
        if (tag.id) {
          submitFunc = updateTag;
        }
        return submitFunc(tag);
      })
      .then(() => {
        tagForm.resetFields();
        setTagFormVisible(false);
        return fetchTags({ rule_id: id, locale });
      })
      .then((newTags) => {
        setTags(newTags);
      })
      .catch((error) => {
        message.error(error.message);
      })
      .finally(() => {
        setTagFormLoading(false);
      });
  };

  const addTag = () => {
    tagForm.setFieldsValue({ rule_id: parseInt(id, 10) });
    tagForm.setFieldsValue({ locale });
    setTagFormTitle('Add Tag');
    setTagFormVisible(true);
  };

  const editTag = (tag) => {
    tagForm.setFieldsValue(tag);
    setTagFormTitle('Edit Tag');
    setTagFormVisible(true);
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
              type="dashed"
              onClick={addTag}
              style={{ width: '100%', marginBottom: 16 }}
            >
              <PlusOutlined />
              Add Tag
            </Button>
            <TagList tags={tags} onEdit={editTag} onDelete={removeTag} />
          </Col>
        </Row>
      </div>

      <Modal
        title={tagFormTitle}
        visible={tagFormVisible}
        onCancel={() => {
          setTagFormVisible(false);
          tagForm.resetFields();
        }}
        confirmLoading={tagFormLoading}
        destroyOnClose
        footer={[
          <Button
            key="back"
            onClick={() => {
              setTagFormVisible(false);
              tagForm.resetFields();
            }}
          >
            Cancel
          </Button>,
          <Button
            key="submit"
            type="primary"
            onClick={submitTag}
          >
            Save
          </Button>,
        ]}
      >
        <TagForm form={tagForm} />
      </Modal>
    </div>
  );
}

export default EditRule;
