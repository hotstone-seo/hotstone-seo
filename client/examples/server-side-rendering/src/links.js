const links = [
  { label: 'Static Page', children: [
    { to: '/airports', label: '/airports' }
  ]},
  { label: 'Page with Parameter', children: [
    { to: '/events/bigbang', label: '/events/bigbang'}
  ]},
  { to: '/nonexistent', label: 'Non-existent Page', exact: true }
];

export default links;
