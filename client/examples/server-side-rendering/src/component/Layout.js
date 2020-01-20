import React from 'react';
import { NavLink } from 'react-router-dom';

export default function Layout(props) {
  const { links } = props;
  const content = props.children;
  const navs = links.map(({ to, label, exact }) => {
    return (
      <li className="nav-item" key={to}>
        <NavLink to={to} exact={exact} className="nav-link" activeClassName="active">
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
