import express from 'express';
import cors from 'cors';
import favicon from 'serve-favicon'
import path from 'path'
import React from 'react';
import { renderToString } from 'react-dom/server';
import { StaticRouter } from 'react-router-dom';
import { Helmet } from 'react-helmet';
import { HotStoneClient } from 'hotstone-client';
import serialize from 'serialize-javascript';
import App from '../component/App';

const server = express();

server.use(cors());
 
// Instantiate the client by providing the URL of HotStone provider
const hotstoneURL = process.env.HOTSTONE_URL || 'http://localhost:8089'
const client = new HotStoneClient(hotstoneURL, {cacheManager: `./test-local-cache`});

const template = ({ body, head }, data) => {
  return `
    <!DOCTYPE html>
    <html ${head.htmlAttributes.toString()}>
      <head>
        <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" rel="stylesheet"
          integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh"
          crossorigin="anonymous" />
        <link href="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.53.2/codemirror.min.css" rel="stylesheet" />

        ${head.title.toString()}
        ${head.meta.toString()}
        ${head.link.toString()}
        ${head.script.toString()}
      </head>
      <body ${head.bodyAttributes.toString()}>
        <div id="root">${body}</div>
        <script>window.__INITIAL_DATA__ = ${serialize(data)}</script>
        <script src="/public/bundle.js"></script>
      </body>
    </html>
  `
}

const port = process.env.PORT || 3000

server.use('/public', express.static('public'));
server.use(favicon(path.join(__dirname, 'public', 'favicon.ico')))

server.get('*', (req, res, next) => {
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
     const tags = await client.tags(rule, "en_US");
    //  const dataWithoutRawHTML = { rule, tags }

     // Rendering element...
    //  const helmet = Helmet.renderStatic();
    //  const rawHTML = template({ body: "", head: helmet }, dataWithoutRawHTML)

    //  const data = {rule, tags, rawHTML}
    
     const data = {rule, tags}
     const appString = renderToString(
      <StaticRouter location={req.url} context={{}} >
        <App data={data} />
      </StaticRouter>
    );
    const helmet = Helmet.renderStatic();

     res.send(template({ body: appString, head: helmet }, data));
   } catch(error) {
     next(error);
   }
 })();
});

server.listen(port, () => {
  console.log(`Listening on port: ${port}`)
})
