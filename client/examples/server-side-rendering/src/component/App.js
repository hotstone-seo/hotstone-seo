import React from 'react';
import { HotStone, HotStoneContext } from 'hotstone-client';
import Layout from './Layout';
import TagInfo from './TagInfo';
import links from '../links';

export default function App(props) {
  const { tags=[] } = props.data;
  return (
    <HotStone tags={tags} >
      <Layout links={links}>
        <HotStoneContext.Consumer>
          {(value) => <TagInfo tags={value} />} 
        </HotStoneContext.Consumer>
      </Layout>
    </HotStone>
  );
}
