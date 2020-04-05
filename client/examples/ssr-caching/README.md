# Tiket.com Mobile (SSR)

Tiket.com Mobile web app project

## Table of Contents
1. [Requirements](#requirements)
2. [Getting Started](#getting-started)
4. [CriticalCSS](#critical-css)
2. [Service Handling](#service-handling)

## Requirements
* node `8.11.1 (LTS)`
* npm `5.6.0`
* redis `^3.2.6`

## Configuration (Local Machine)
- Copy Environment configuration file from .env.example to .env
```bash
$ cp .env.example .env
```
- Modify it based on your configuration preference

## Configuration (Server Machine)
There are two method on how to set the environment configuration :
- Using .env File
  - Copy Environment configuration file from .env.example to .env
    ```bash
    $ cp .env.example .env
    ```
  
  - Modify it based on server configuration preference
- Using Machine Environment
  - Edit ~/.bash_profile or create a bash script that run the server
  - Put all the required environment configuration there
  - ``` 
    # NodeJS Environment
    NODE_ENV=production
    
    # Cluster worker count
    # Remove Comment to limit workers
    #RECLUSTER_WORKERS=1
    
    HOST=localhost
    PORT=3000
    SECRET=tiket.com!@
    
    # Host Configuration example: https://mpayment.tiket.com
    APPHOST=http://localhost:3000
    
    # Session Configuration
    # Pick betwen session storage method mysql | redis | memory
    SESSION_DRIVER=memory
    
    # Redis as session storage
    #SESSION_REDIS_SERVER=127.0.0.1
    #SESSION_REDIS_PORT=6379
    #SESSION_REDIS_SECRET=blabla
    
    # Mysql as session storage
    #SESSION_MYSQL_SERVER=127.0.0.1
    #SESSION_MYSQL_PORT=3306
    #SESSION_MYSQL_USER=root
    #SESSION_MYSQL_PASS=
    #SESSION_MYSQL_DB=session
    
    # NewRelic Configuration
    NEW_RELIC_ENABLED=false
    NEW_RELIC_NO_CONFIG_FILE=true
    NEW_RELIC_APP_NAME=payment-desktop-dev
    NEW_RELIC_LICENSE_KEY=
    NEW_RELIC_LOG_LEVEL=info
    
    # Log configuration
    #LOGDIR=/var/logs/
    #LOG_SIZE=500m
    #LOG_KEEP=5
    
    # Build Configuration
    SOURCE_MAP=true
    DEBUG=true
    

## Getting Started

After confirming that your development environment meets the specified [requirements](#requirements),
you can start the site by running these commands:

```bash
$ npm install                   # Install project dependencies
$ npm start                     # Compile and launch
```

While developing, you will probably rely mostly on `npm start`; however, there are additional scripts at your disposal:

|`npm run <script>`|Description|
|------------------|-----------|
|`start` |Serves your app at `localhost:3000`. HMR will be enabled and `client.css` file will not be generated.|
|`build`|Compiles the application to disk (`~/build` by default).|
|`serve` |Serves your builded app at `localhost:3000`. Using `./bin/cluster`, to simulate production behaviour.|
|`test`|Runs all tests in sequence|
|`lint:js`|Run javascript linter.|
|`lint:scss`|Run scss linter.|
|`release`|Runs test, lint:js, lint:scss, build and bump version then create and push tags.|
|`storybook`|Run storybook to view the UI components.|
|`sonar-scanner`|Run sonar-scanner and upload your test results to sonar.|
***Important note:***

Before you commit, make sure to always run:

```bash
$ npm test
```

and have all the tests pass.

## Service Handling

**CPU USAGE**  
By default app will use all of the CPU core to spawn it's worker,
this can be changed by modifying `RECLUSTER_WORKERS=1` to any CPU count you require

**GRACEFUL RESTART**  
To make sure that all process were not used before killing the process
and gracefully kill the child process use :  
``kill -s SIGUSR2 `cat /var/project-path/project-name.pid` ``  

to change PID file output location you could modify the env `PIDFILE=/var/project-name.pid`


**GRACEFUL KILL**  
``kill -s SIGTERM `cat /var/project-path/project-name.pid` ``

## Branches
There are following branches used in project:
* `master` that is what is running on production server
* `devel` is always on top of `master` (fast-forward) and that is branch you should fork when you start working on a new feature
That is also the branch you should rebase onto daily while you are working on your feature.
* `feature/*` each feature may have it's own branch and that branch may be deployed somewhere for testing as well.

## File and Folder naming convention
* React JS or JSX or Javascript Class shoud be named Pascal case, for example - OrderDetailPanel.jsx
* Js files should be named in camel case, for example - verifyPhone.js
* folder names should be all in small letters, if it's more than a word, then they should be seperated by -, for example -> core-library,
But if the folder is behave as React component index container then it should use Pascal case.

## Critical CSS
Why we do *critical css*? [See Here](https://www.sitepoint.com/how-and-why-you-should-inline-your-critical-css/)
* You'll found many `critical.scss` file in this codebase. We need to convert it to css first by running `lessc <filename.scss> > <outputFilename.css>`. If you want to have new critical css, first please rename your file to `critical.scss`, so other people will recognize it as a critical style. Also don't forget to remove the import from the file. Please choose wisely which style should be *inlined* and which style shouldn't.

* Copy the style inside the generated file, then copy and add it into `/core/html/critical.js`, but remember by adding to `critical.js` file, your style will be loaded in all pages and will **increase the DOM size**. If you don't want to, you can make a new file locally inside the particular routes. You can see the example in `routes/Landing/components/critical.js`. By doing so, your critical css won't be loaded in all pages and not damaging other pages performance.
