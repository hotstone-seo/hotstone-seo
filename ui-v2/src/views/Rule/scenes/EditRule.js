import React, { useState, useEffect } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import {
  Row, Col, message, Select, Button, Modal,
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

  const [rule, setRule] = useState({});
  const [tags, setTags] = useState([]);
  const [locale, setLocale] = useState(locales[0] || '');
  const [tagFormVisible, setTagFormVisible] = useState(false);

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

  const submitTag = (tag) => {
    let submitFunc = createTag;
    if (tag.id) {
      submitFunc = updateTag;
    }

    submitFunc(tag)
      .then(() => {
        refreshTagList(id, locale);
      })
      .catch(error => {
        message.error(error.message);
      });
  };

  const refreshTagList = (id, locale) => {
    fetchTags({ rule_id: id, locale: locale })
      .then(tags => { setTags(tags) })
      .catch(error => {
        message.error(error.message);
      });
  };

  const removeTag = (tag) => {
    deleteTag(tag.id)
      .then(() => {
        setTags(tags.filter(item => item.id !== tag.id));
      })
      .catch(error => {
        message.error(error.message);
      });
  };

  const changeLocale = (newLocale) => {
    setLocale(newLocale);
  };

  const openTagForm = () => {
    setTagFormVisible(true);
  }

  const closeTagForm = () => {
    setTagFormVisible(false);
  }

  return (
    <div>
      <Row>
        <Col className={styles.container} span={12} style={{ paddingTop: 24 }}>
          <RuleForm handleSubmit={editRule} rule={rule} />
        </Col>
      </Row>
      <Row style={{ marginTop: 24 }}>
        <Col className={styles.container} span={16} style={{ padding: 24 }}>
          <Select defaultValue={locale} onChange={changeLocale}>
            {locales.map(locale => (
              <Option value={locale}>{locale}</Option>
            ))}
          </Select>
          <TagList tags={tags} onDelete={removeTag} />
          <Button
            type="dashed"
            onClick={openTagForm}
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
        onCancel={closeTagForm}
      >
        <TagForm />
      </Modal>
    </div>
  );
}

export default EditRule;
