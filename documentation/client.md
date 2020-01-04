# Client

The client is a javascript library responsible for retrieving data from HotStone
Provider as well as providing support for rendering elements.

## How to Use

You can install the library using your preferred package manager:

```sh
$ npm install hotstone-client
```

After successfully adding the dependency, you can create a HotStone client
instance by providing it with URL of running HotStone provider.

```js
import HotStone from 'hotstone-client';
const client = HotStone('https://hotstone-provider');
```

An instance of HotStone client works in terms of context. A context relates to a
path, in which we want to get tags for.

Please also note that since some of the client methods are asynchronous, you
should run it inside an async function block.

```js
const context = client.match('/current/path');
const data = await context.data();
    
// You can return a JSX elements containing the tags, the data is not
// necessarily nave to come from HotStone,
const tags = context.renderTags(data);

// Beside tag elements, you can also render content
const contents = context.renderContents(data);
```
