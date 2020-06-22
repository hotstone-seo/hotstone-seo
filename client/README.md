# HotStone Client

The client is a javascript library responsible for retrieving data from HotStone
API Server. 

## Installation

To install the client to your project simply run the npm install command:

``` bash
npm install hotstone-client
```

## Usage

``` javascript
import HotStone from 'hotstone-client';

(async function() {
  // Instantiate the client by providing the URL of HotStone provider
  const client = HotStone('http://localhost:4000');

  // Retrieve a Page SEO Rule by mathing its path
  // A rule is a colelction of tags or content that are tailored specifically
  // for the page
  const rule = await client.match('/any/path');

  // Get the tags associated with the rule, a tags is specific to a locale.
  // A tags is an array of tag object containing info of a specific HTML tag
  // Example:
  // [
  //   { type: "title", attributes: {}, value: "Page Title" },
  //   { type: "meta", attributes: { name: "description", content: "Page Description" } }
  // ]
  const tags = await rule.tags(rule, 'en-US');

  // Rendering tag element...
})();
```

## Changelog

- 0.5.2
  - `react` as peer dependency
  - `renderHelmetTags` on `'hotstone-client/lib/react'`
- 0.5.0
  - Support auth with client key to access HotStone service
  - Support HTTP caching