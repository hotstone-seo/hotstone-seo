import React from 'react';
import { Helmet } from 'react-helmet';

const toElements = (tags) => {
  return tags.map(({ type, attributes, value }) => (
    React.createElement(type, attributes, value)
  ));
}

export default function App(props) {
  const { rule, tags } = props;
  // TODO: Create a view for displaying rule and tags
  return (
    <div>
      <Helmet>{toElements(tags)}</Helmet>
    </div> 
  );
}
