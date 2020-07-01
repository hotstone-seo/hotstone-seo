# HotStone Client

The client is a javascript library responsible for retrieving data from HotStone
API Server.

## Documentation

Read [HotStone Client Usage](https://hotstone-seo.github.io/getting-started/hotstone-client-usage.html) and [HotStone Client Config](https://hotstone-seo.github.io/advanced/hotstone-client-config.html)

## Changelog

- 0.5.3
  - Avoid global handling error; Only handle (try-catch) invalid response format
  - Fix: strip trailing slash of baseURL
- 0.5.2
  - `react` as peer dependency
  - `renderHelmetTags` on `'hotstone-client/lib/react'`
- 0.5.0
  - Support auth with client key to access HotStone service
  - Support HTTP caching