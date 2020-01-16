import React from 'react';
import { HotStone } from 'hotstone-client';
import TagInfo from './TagInfo';

export default function App(props) {
  const { rule, tags=[] } = props.data;
  return (
    <div>
      <HotStone tags={tags} />
      <h2>Tags Received</h2>
      <TagInfo tags={tags} />
    </div> 
  );
}
