import React from 'react';

const toRows = (tags) => {
  return tags.map(({ type, attributes, value }, index) => {
    const attrs = Object.keys(attributes).map((key, i) => (
      <li key={i}>
        {key}: {attributes[key]}
      </li>
    ));
    return (
      <tr key={index}>
        <td>{type}</td>
        <td>
          <ul>{attrs}</ul>
        </td>
        <td>{value}</td>
      </tr>
    );
  });
}

export default function TagInfo(props) {
  const { tags } = props;
  return (
    <table className="table">
      <thead>
        <tr>
          <th>Type</th>
          <th>Attributes</th>
          <th>Value</th>
        </tr>
      </thead>
      <tbody>
        {toRows(tags)}
      </tbody>
    </table>
  );
}
