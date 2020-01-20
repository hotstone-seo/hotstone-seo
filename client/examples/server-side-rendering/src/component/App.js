import React from 'react';
import { HotStone } from 'hotstone-client';
import Layout from './Layout';
import TagInfo from './TagInfo';

const links = [
  { to: '/airports', label: 'Static Link', exact: true },
  { to: '/airports/12', label: 'Dynamic Link' },
  { to: '/nonexistent', label: 'Non-existent Link', exact: true }
];

export default function App(props) {
  const { tags=[] } = props.data;
  return (
    <div>
      <HotStone tags={tags} />
      <Layout links={links}>
        <TagInfo tags={tags} />
      </Layout>
    </div> 
  );
}
