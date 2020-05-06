const links = [
  { label: 'Static Page', children: [
    { to: '/airports', label: '/airports', desc: '/airports' }
  ]},
  { label: 'Page with Parameter', children: [
    { to: '/events/bigbang', label: '/events/<eventName>', desc: '/events/bigbang'}
  ]},
  { to: '/nonexistent', label: 'Non-existent Page', exact: true }
];

export default links;
