import React from 'react';
import { Helmet } from 'react-helmet';

const toElements = (tags) => {
  return tags.map(({ type, attributes, value }) => (
    React.createElement(type, attributes, value)
  ));
}

const toRows = (tags) => {
  return tags.map(({ type, attributes, value }, index) => {
    const attrs = Object.keys(attributes).map((key, i) => (
      <li key={i}>
        {key}: {attributes[key]}
      </li>
    ))
    return (
      <tr key={index}>
        <td>{type}</td>
        <td>
          <ul>{attrs}</ul>
        </td>
        <td>{value}</td>
      </tr>
    );
  })
}

export default function App(props) {
  const { rule, tags } = props;
  return (
    <div>
      <Helmet>{toElements(tags)}</Helmet>
      <h1>Sample Application using HotStone</h1>
      <table>
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
    </div> 
  );
}
