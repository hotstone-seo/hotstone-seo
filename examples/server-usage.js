import HotStone from '../src';

(async function() {
  const client = HotStone('http://localhost:4000');
  const pathContext = client.match('/any/path');

  // We can retrieve tags data using tags() method on a context object,
  // please note that retrieving tags is a Promise.
  const tags = await pathContext.tags()

  // After getting the data you can render the tags. This is done by
  // using React.
  const tagElements = pathContext.renderTags(tags);
})();
