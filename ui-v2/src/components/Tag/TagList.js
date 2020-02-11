import React from 'react';
import ReactDOMServer from 'react-dom/server';
import { List, Button } from 'antd';

// TODO: Make sure that tag with no children like <meta /> don't have
// "value" key in its object or has null or undefined as its value.
// It interferes with React's renderToStaticMarkup function
// which will create an empty children whenever the value key is set
// to empty string.
function TagList(props) {
  const tags = props.tags.map(tag => {
    if (tag.type === 'meta') {
      tag.value = null
    }
    return tag;
  });
  return (
    <List
      dataSource={tags}
      renderItem={({ type, attributes, value }) => (
        <List.Item
          actions={[
            <Button type="link" style={{ padding: 0 }}>Edit</Button>,
            <Button type="link" danger style={{ padding: 0 }}>Delete</Button>
          ]}
        >
          <pre>
            {ReactDOMServer.renderToStaticMarkup(
              React.createElement(type, attributes, value)
            )}
          </pre>
        </List.Item>
      )}
    />
  );
}

export default TagList;
