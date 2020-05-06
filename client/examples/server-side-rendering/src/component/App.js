import React from 'react';
import { HotStone, HotStoneContext } from './hotstone-client';
import Layout from './Layout';
import TagInfo from './TagInfo';
import links from '../links';
import SyntaxHighlighter from 'react-syntax-highlighter';
import { docco } from 'react-syntax-highlighter/dist/cjs/styles/hljs';

export default function App(props) {
  const { tags=[], rawHTML } = props.data;
  return (
    <HotStone tags={tags} >
      <Layout links={links}>
        {/* <HotStoneContext.Consumer>
          {(value) => <TagInfo tags={value} />} 
        </HotStoneContext.Consumer> */}
        <SyntaxHighlighter language="html" style={docco}>
          {rawHTML}
        </SyntaxHighlighter>
      </Layout>
    </HotStone>
  );
}
