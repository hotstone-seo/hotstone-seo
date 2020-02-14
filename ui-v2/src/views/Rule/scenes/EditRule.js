import React, { useState, useEffect } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import {
  Row, Col, message, Select, Button, Modal, Form,
} from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { RuleForm } from 'components/Rule';
import { TagList, TagForm } from 'components/Tag';
import { getRule, updateRule } from 'api/rule';
import { fetchTags, createTag, updateTag, deleteTag } from 'api/tag';
import locales from 'locales';
import styles from './AddRule.module.css';

const { Option } = Select;

function EditRule() {
  const { id } = useParams();
  const history = useHistory();
  const [tagForm] = Form.useForm();

  const [rule, setRule] = useState({});
  const [tags, setTags] = useState([]);
  const [locale, setLocale] = useState(locales[0] || '');
  const [tagFormVisible, setTagFormVisible] = useState(false);
  const [tagFormLoading, setTagFormLoading] = useState(false);

  useEffect(() => {
    getRule(id)
      .then(rule => { setRule(rule) })
      .catch(error => {
        message.error(error.message);
      });
  }, [id]);

  useEffect(() => {
    refreshTagList(id, locale);
  }, [id, locale])

  const editRule = (newRule) => {
    newRule.id = rule.id;
    updateRule(newRule)
      .then(() => {
        history.push('/rules');
      })
      .catch(error => {
        message.error(error.message);
      });
  };

  const submitTag = () => {
    tagForm.validateFields()
           .then((tag) => {
             setTagFormLoading(true);
             
             tag.rule_id = parseInt(id);
             let submitFunc = createTag;
             if (tag.id) {
               submitFunc = updateTag;
             }
             return submitFunc(tag);
           })
           .then(() => {
             tagForm.resetFields();
             setTagFormVisible(false);
             refreshTagList(id, locale);
           })
           .catch(error => {
             message.error(error.message);
           })
           .finally(() => {
             setTagFormLoading(false);
           });
  };

  const refreshTagList = (id, locale) => {
    fetchTags({ rule_id: id, locale: locale })
      .then(tags => { setTags(tags) })
      .catch(error => {
        message.error(error.message);
      });
  };

  const editTag = (tag) => {
    tagForm.setFieldsValue(tag);
    setTagFormVisible(true);
  }

  const removeTag = (tag) => {
    deleteTag(tag.id)
      .then(() => {
        setTags(tags.filter(item => item.id !== tag.id));
      })
      .catch(error => {
        message.error(error.message);
      });
  };

  return (
    <div>
      <Row>
        <Col className={styles.container} span={16} style={{ paddingTop: 24 }}>
          <RuleForm handleSubmit={editRule} rule={rule} />
        </Col>
      </Row>
      <Row style={{ marginTop: 24 }}>
        <Col className={styles.container} span={16} style={{ padding: 24 }}>
          <Select
            defaultValue={locale}
            onChange={(value) => setLocale(value)}
          >
            {locales.map(locale => (
              <Option value={locale}>{locale}</Option>
            ))}
          </Select>
          <TagList tags={tags} onEdit={editTag} onDelete={removeTag} />
          <Button
            type="dashed"
            onClick={() => setTagFormVisible(true)}
            style={{ width: '100%' }}
          >
            <PlusOutlined /> Add Tag
          </Button>
        </Col>
      </Row>
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
