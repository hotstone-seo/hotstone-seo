import React from 'react';
import ReactDOMServer from 'react-dom/server';
import { List, Button, Popconfirm } from 'antd';

function TagList({ tags, onEdit, onDelete }) {
  const sanitizeTag = (tag) => {
    if (tag.type === 'meta') {
      tag.value = null;
    }
    return tag;
  }
  return (
    <List
      dataSource={tags.map(sanitizeTag)}
      renderItem={(tag) => (
        <List.Item
          actions={[
            <Button
              type="link"
              onClick={() => onEdit(tag)}
              style={{ padding: 0 }}
            >
              Edit
            </Button>,
            <Popconfirm
              title="Are you sure to delete this tag?"
              placement="topRight"
              onConfirm={() => onDelete(tag)}
            >
              <Button type="link" danger style={{ padding: 0 }}>Delete</Button>
            </Popconfirm>
          ]}
        >
          <pre>
            {ReactDOMServer.renderToStaticMarkup(
              React.createElement(tag.type, tag.attributes, tag.value)
            )}
          </pre>
        </List.Item>
      )}
    />
  );
}

export default TagList;
