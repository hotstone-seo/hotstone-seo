import React, { useState, useEffect } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import {
  PageHeader, Row, Col, message, Select, Button, Modal, Form,
} from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { RuleForm, RuleDetail } from 'components/Rule';
import { TagList, TagForm } from 'components/Tag';
import { getRule, updateRule } from 'api/rule';
import {
  fetchTags, createTag, updateTag, deleteTag,
} from 'api/tag';
import useDataSources from 'hooks/useDataSources';
import locales from 'locales';
import styles from './AddRule.module.css';

const { Option } = Select;

function EditRule() {
  const { id } = useParams();
  const history = useHistory();
  const [dataSources] = useDataSources();
  const [tagForm] = Form.useForm();

  const [rule, setRule] = useState({});
  const [tags, setTags] = useState([]);
  const [locale, setLocale] = useState(locales[0] || '');
  const [tagFormVisible, setTagFormVisible] = useState(false);
  const [tagFormLoading, setTagFormLoading] = useState(false);

  useEffect(() => {
    getRule(id)
      .then((newRule) => { setRule(newRule); })
      .catch((error) => {
        message.error(error.message);
      });
  }, [id]);

  useEffect(() => {
    fetchTags({ rule_id: id, locale })
      .then((newTags) => { setTags(newTags); })
      .catch((error) => {
        message.error(error.message);
      });
  }, [id, locale]);

  const editRule = (newRule) => {
    updateRule(newRule)
      .then(() => {
        history.push('/rules');
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
    setTagFormVisible(true);
  };

  const editTag = (tag) => {
    tagForm.setFieldsValue(tag);
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
        style={{ background: '#fff' }}
      >
        <RuleDetail rule={rule} />
      </PageHeader>

      <div style={{ padding: 24 }}>
        <Row>
          <Col className={styles.container} span={16} style={{ paddingTop: 24 }}>
            <RuleForm handleSubmit={editRule} rule={rule} dataSources={dataSources} />
          </Col>
        </Row>
        <Row style={{ marginTop: 24 }}>
          <Col className={styles.container} span={16} style={{ padding: 24 }}>
            <Select
              defaultValue={locale}
              onChange={(value) => setLocale(value)}
            >
              {locales.map((loc) => (
                <Option value={loc} key={loc}>{loc}</Option>
              ))}
            </Select>
            <TagList tags={tags} onEdit={editTag} onDelete={removeTag} />
            <Button
              type="dashed"
              onClick={addTag}
              style={{ width: '100%' }}
            >
              <PlusOutlined />
              Add Tag
            </Button>
          </Col>
        </Row>
      </div>

      <Modal
        title="Add/Edit Tag"
        visible={tagFormVisible}
        onOk={submitTag}
        onCancel={() => {
          setTagFormVisible(false);
          tagForm.resetFields();
        }}
        confirmLoading={tagFormLoading}
      >
        <TagForm form={tagForm} />
      </Modal>
    </div>
  );
}

export default EditRule;
