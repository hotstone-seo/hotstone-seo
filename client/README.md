# HotStone Client

The client is a javascript library responsible for retrieving data from HotStone
API Server.

## Installation

To install the client to your project simply run the npm install command:

``` bash
$ npm install hotstone-client

# to use 'renderHelmetTags'
$ npm install react
```

## Server Side Usage

`HotStoneClient` is server-side library providing API to match and to get data of HTML tags of given path. To initiate, it requires two mandatory arguments and last optional argument:

```js
const client = new HotStoneClient(hotstoneURL, clientKey, fetchOpts)
```

- (mandatory) `hotstoneURL`: URL of **HotStone Provider**
- (mandatory) `clientKey`: To interact with **HotStone Provider**, it needs client key. Get client key from **HotStone UI** / Dashboard on menu `Client Keys`
- (optional) `fetchOpts`: Under the hood, **HotStone Client** uses [make-fetch-happen](https://www.npmjs.com/package/make-fetch-happen). We can pass [extra fetch options](https://www.npmjs.com/package/make-fetch-happen#extra-options) to get additional features, i.e. local caching with `{ cacheManager: './hotstone-local-cache' }`

``` javascript
import { HotStoneClient } from 'hotstone-client';

(async function() {
  const hotstoneURL = 'http://localhost:8089'
  const clientKey = 'xxxxxxx.yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy'
  const fetchOpts = { cacheManager: `./hotstone-local-cache` }
  // Instantiate the client by providing the URL of HotStone provider and client key
  const client = new HotStoneClient(hotstoneURL, clientKey, fetchOpts);

  // Retrieve a Page SEO Rule by matching its path
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