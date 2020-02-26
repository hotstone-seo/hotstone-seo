import React from "react";
import { Link } from "react-router-dom";
import _ from "lodash";
import useWindowSize from "../../hooks/useWindowSize";

export default function RawHTMLPreview({ ruleID, tags }) {
  const size = useWindowSize();
  if (_.isEmpty(tags)) {
    return (
      <div>
        No tags data. Register tags at{" "}
        <Link to={`/rule-detail/?id=${ruleID}`}>Rule Detail</Link>
      </div>
    );
  } else {
    const textAreaVal = tags
      .map(({ type, value, attributes }) => {
        let attributesStr = "";
        if (!_.isEmpty(attributes)) {
          if (_.isPlainObject(attributes)) {
            Object.entries(attributes).forEach(([key, value]) => {
              attributesStr += ` ${key}="${value}"`;
            });
          } else if (_.isArray(attributes)) {
            attributes.forEach(attributes => {
              Object.entries(attributes).forEach(([key, value]) => {
                attributesStr += ` ${key}="${value}"`;
              });
            });
          }
        }

        return `<${type}${attributesStr}>${value}</${type}>`;
      })
      .join("\n");

    return (
      <pre style={{ border: "1px solid grey", padding: 8 }}>
        <code>{textAreaVal}</code>
      </pre>
    );
  }
}
