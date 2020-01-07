import express from 'express';
import React from 'react';
import { renderToString } from 'react-dom/server';
import { Helmet } from 'react-helmet';
import HotStone from 'hotstone-client';
import App from './app';

const server = express()
// Instantiate the client by providing the URL of HotStone provider
const client = new HotStone('http://localhost:4000');

const template = ({ body, head }) => {
  return `
    <!DOCTYPE html>
    <html ${head.htmlAttributes.toString()}>
      <head>
        ${head.title.toString()}
        ${head.meta.toString()}
        ${head.link.toString()}
      </head>
      <body ${head.bodyAttributes.toString()}>
        <div id="root">${body}</div>
      </body>
      <script src="/public/bundle.js"></script>
    </html>
  `
}

server.use('/public', express.static('../public'));
server.get('/', (req, res, next) => {
 (async function() {
   try {
     // Retrieve a Page SEO Rule by mathing its path
     // A rule is a colelction of tags or content that are tailored specifically
     // for the page
     const rule = await client.match(req.path);
     
     // Get the tags associated with the rule and locale, you can also provide it with data
     // to override using server provided data.
     // A tags is an array of tag object containing info of a specific HTML tag
     // Example:
     // [
     //   { type: "title", attributes: {}, value: "Page Title" },
     //   { type: "meta", attributes: { name: "description", content: "Page Description" } }
     // ]
     const tags = await client.tags(rule, "en-US");

     // Rendering element...
     const helmet = Helmet.renderStatic();
     const appString = renderToString(<App rule={rule} tags={tags} />);
     res.send(template({
       body: appString,
       head: helmet,
     }));
   } catch(error) {
     next(error);
   }
 })();
});

server.listen(8080)
