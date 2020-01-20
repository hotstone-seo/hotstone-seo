import React from 'react';
import { NavLink } from 'react-router-dom';

export default function Layout(props) {
  const { content, links } = props;
  const navs = links.map(({ to, label }) => {
    return (
      <li class="nav-item">
        <NavLink to={to} className="nav-link" activeClassName="active">
          {label}
        </NavLink>
      </li>
    )
  });
  return (
    <div className="container"> 
      <ul className="nav nav-pills">
        {navs}
      </ul>
      {content}
    </div>
  );
}
