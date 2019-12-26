import HotStone from 'hotstone-client';

(async function() {
  // Instantiate the client by providing the URL of HotStone provider
  const client = HotStone('http://localhost:4000');

  // Retrieve a Page SEO Rule by mathing its path
  // A rule is a colelction of tags or content that are tailored specifically
  // for the page
  const rule = client.match('/any/path');

  // Get the tags associated with the rule
  // A tags is an array of tag object containing info of a specific HTML tag
  // Example:
  // [
  //   { type: "title", attributes: {}, value: "Page Title" },
  //   { type: "meta", attributes: { name: "description", content: "Page Description" } }
  // ]
  const tags = await rule.tags();

  // Return a React element in order to be rendered
  const tagElements = tags.element();

  // Rendering tag element...
})();
