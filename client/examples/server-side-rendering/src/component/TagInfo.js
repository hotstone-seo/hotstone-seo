import React from 'react';
import toString from 'react-element-to-jsx-string';

const toReactElement = ({ type, attributes, value }) => {
  return toString(React.createElement(type, attributes, value));
}

export default function TagInfo(props) {
  const { tags } = props;
  const rawTags = tags.map(toReactElement);
  return (
    <pre>{rawTags.join(`\n`)}</pre>
  );
}
