import React, {useState, useEffect} from 'react';
import { NavLink } from 'react-router-dom';

export default function Layout(props) {

  const [currentURL, setCurrentURL] = useState("");

  useEffect(() => {
    if (__isBrowser__) {
      setCurrentURL(window.location.href)
    }
  });

  const { links } = props;
  const content = props.children;
  const navs = links.map(({ to, label, exact }) => {
    return (
      <li className="nav-item" key={to}>
        {/* <NavLink to={to} exact={exact} className="nav-link" activeClassName="active">
          {label}
        </NavLink> */}
        <a class="nav-link" href={to}>{label}</a>
      </li>
    )
  });

  const currentURLEnc = encodeURIComponent(currentURL)
  const tools = [
    {name: "Google Structured Data Testing Tool", url: `https://search.google.com/structured-data/testing-tool/u/0/#url=${currentURLEnc}`},
  ].map(({name, url}) => {
    return (
      <li className="nav-item" key={name}>
        <a target="_blank" class="nav-link" href={url}>{name}</a>
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
          Tag Response
        </div>
        <div className="card-body">
          {content}
        </div>
      </div>
    </div>
  );
}
