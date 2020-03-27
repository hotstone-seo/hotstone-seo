# HotStone UI

This project uses a running HotStone server as API endpoint. Please see the
[documentation](https://github.com/hotstone-seo/hotstone-seo/README.me) on how
to run the server

## How to start

This project was bootstrapped with [Create React
App](https://github.com/facebook/create-react-app), so all of the commands
available there works.

To run the available scripts, navigate to the ui project directory where
`package.json` is located.

### Start the development server

`npm start`

You also need to provide an `.env` file as the app configuration at the root
folder to indicate where your HotStone server is located

``` sh
REACT_APP_API_URL=http://localhost:8089
```

### Build for production

`npm run build` will bundles the application to `dist/` folder and you can deploy
the artifacts to your static server.
