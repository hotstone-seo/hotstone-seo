import React from 'react';
import PropTypes from 'prop-types';
import { renderToStaticMarkup } from 'react-dom/server';

function TagPreview({ tag }) {
  const { type, attributes, value } = tag;
  return (type
    ? (
      <pre>
        {renderToStaticMarkup(
          React.createElement(type, attributes, value === '' ? null : value),
        )}
      </pre>
    ) : null
  );
}

TagPreview.propTypes = {
  tag: PropTypes.shape({
    type: PropTypes.string.isRequired,
    attributes: PropTypes.object,
    value: PropTypes.string,
  }).isRequired,
};

export default TagPreview;
