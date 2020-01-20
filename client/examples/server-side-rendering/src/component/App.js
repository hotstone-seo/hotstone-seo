import React from 'react';
import { HotStone } from 'hotstone-client';
import Layout from './Layout';
import TagInfo from './TagInfo';

const links = [
  { to: '/airport', label: 'Static Link' },
  { to: '/airport/12', label: 'Dynamic Link' },
  { to: '/nonexistent', label: 'Non-existent Link' }
];

export default function App(props) {
  const { tags=[] } = props.data;
  return (
    <div>
      <HotStone tags={tags} />
      <Layout links={links} content={<TagInfo tags={tags} />} />
    </div> 
  );
}
