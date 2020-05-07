import React, { useState, useEffect } from 'react';
import { NavItem, NavLink } from 'react-router-dom';
import Dropdown from 'react-bootstrap/Dropdown'
import Iframe from 'react-iframe'
import HtmlPreview from './HtmlPreview'

export default function Layout(props) {

  const [currentURL, setCurrentURL] = useState("");

  useEffect(() => {
    if (__isBrowser__) {
      setCurrentURL(window.location.href)
    }
  });

  const renderLink = (label, to) => {
    return (
      <li className="nav-item" key={to}>
        <a className="nav-link" href={to}>{label}</a>
      </li>
    )
  }

  const { links } = props;
  const content = props.children;
  const navs = links.map((link) => {
    if ('to' in link) {
      return renderLink(link.label, link.to)
    }
    if ('children' in link) {
      const {label, children} = link
      return (
        <Dropdown size="sm" style={{ padding: "0.1em" }}>
          <Dropdown.Toggle>{label}</Dropdown.Toggle>
          <Dropdown.Menu>
            {children.map(({label, desc, to}) => (
              <Dropdown.Item key={to} href={to}>
                <div>
                  Rule:<code>{label}</code>{" "}
                  URL:<strong>{desc}</strong> 
                </div>
              </Dropdown.Item>
            ))}
          </Dropdown.Menu>
        </Dropdown>
      )
    }
  });

  const currentURLEnc = encodeURIComponent(currentURL)
  const tools = [
    { name: "Google Structured Data Testing Tool", url: `https://search.google.com/structured-data/testing-tool/u/0/#url=${currentURLEnc}` },
  ].map(({ name, url }) => {
    return (
      <li className="nav-item" key={name}>
        <a target="_blank" className="nav-link" href={url}>{name}</a>
      </li>
    )
  })
  return (
    <div className="container">
      <ul className="nav nav-pills" style={{ padding: "1em" }}>
        {navs}
        {tools}
      </ul>
      <div className="card">
        <div className="card-header">
          Response
        </div>
        <div className="card-body">
          {content}
        </div>
      </div>
      {/* <Iframe url={`https://search.google.com/structured-data/testing-tool/u/0/#url=${currentURLEnc}`}
            width="100%"
            height="100%"
            id="myId"
            className="myClassname"
            display="initial"
            position="relative"/> */}
      <HtmlPreview url={currentURL} /> 
    </div>
  );
}
