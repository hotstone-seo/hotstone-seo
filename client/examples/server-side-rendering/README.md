# Sample React application with SSR

This sample provides some idea on how hotstone can be utilized.

## Installation

Navigate to server-side-rendering root directory and install the dependencies by
running this command:

``` sh
npm install
```

Build the bundled js file both for server-side and application by running the
script:

``` sh
npm run build
```

Finally, run the application by running the script:

``` sh
npm run start
```

The application should now be running at the default 3000 port. If you wish to
use another port, simply provide the environment variable.

``` sh
PORT=4000 npm run start
```

## Run While Editing `hotstone-client`

```
cd ../../ && npm run build && npm pack && cd - && npm i && npm run start
```