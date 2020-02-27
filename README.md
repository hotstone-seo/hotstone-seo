

# hotstone-seo

[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)



### Usage

| Usage | Description |
|---|---|
|`hotstone-seo`|Run the application|


### Configuration

| Name | Type | Default | Required |
|---|---|---|:---:|
|APP_ADDRESS|string|:8089|Yes|
|APP_AUTH_LOGOUT_REDIRECT|string|||
|APP_COOKIE_SECURE|bool|||
|APP_DEFAULTLOCALE|string|id_ID||
|APP_JWT_SECRET|string|||
|APP_OAUTH2_GOOGLE_CALLBACK|string|||
|APP_OAUTH2_GOOGLE_CLIENT_ID|string|||
|APP_OAUTH2_GOOGLE_CLIENT_SECRET|string|||
|APP_OAUTH2_GOOGLE_HOSTED_DOMAIN|string|||
|APP_OAUTH2_GOOGLE_REDIRECT_FAILURE|string|||
|APP_OAUTH2_GOOGLE_REDIRECT_SUCCESS|string|||
|PG_DBNAME|string||Yes|
|PG_HOST|string|localhost||
|PG_PASSWORD|string|pgpass|Yes|
|PG_PORT|int|5432||
|PG_USER|string|postgres|Yes|
|REDIS_DB|int|0||
|REDIS_DIAL_TIMEOUT|Duration|5s|Yes|
|REDIS_HOST|string|localhost|Yes|
|REDIS_IDLE_CHECK_FREQUENCY|Duration|1m|Yes|
|REDIS_IDLE_TIMEOUT|Duration|5m|Yes|
|REDIS_MAX_CONN_AGE|Duration|30m|Yes|
|REDIS_PASSWORD|string|redispass||
|REDIS_POOL_SIZE|int|20|Yes|
|REDIS_PORT|string|6379|Yes|
|REDIS_READ_WRITE_TIMEOUT|Duration|3s|Yes|

----

## Development Guide

### Prerequisite

Install [Go](https://golang.org/doc/install) (It is recommend to install via [Homebrew](https://brew.sh/) `brew install go`)

### Quick Start

```bash
# equivalent with `docker-compose up -d` (if infrastructure not up)
./typicalw docker up 

# drop, create and migrate postgres database (if database not ready)
./typicalw pg reset 

# generate readme (if there is readme update)
./typicalw readme 

# generate mock (if require mock)
./typicalw mock 

# run test 
./typicalw test

# run the application
./typicalw run 

# release the distribution
./typicalw release 
```

### Commands
| Command | Description |
|---|---|
|`./typicalw build`|Build the binary|
|`./typicalw run`|Run the binary|
|`./typicalw clean`|Clean the project from generated file during build time|
|`./typicalw test`|Run the testing|
|`./typicalw mock`|Generate mock class|
|`./typicalw release`|Release the distribution|
|`./typicalw docker`|Docker utility|
|`./typicalw docker compose`|Generate docker-compose.yaml|
|`./typicalw docker up`|Spin up docker containers according docker-compose|
|`./typicalw docker down`|Take down all docker containers according docker-compose|
|`./typicalw docker wipe`|Kill all running docker container|
|`./typicalw readme`|Generate README Documentation|
|`./typicalw postgres`|Postgres Database Tool|
|`./typicalw postgres create`|Create New Database|
|`./typicalw postgres drop`|Drop Database|
|`./typicalw postgres migrate`|Migrate Database|
|`./typicalw postgres rollback`|Rollback Database|
|`./typicalw postgres seed`|Data seeding|
|`./typicalw postgres reset`|Reset Database|
|`./typicalw postgres console`|PostgreSQL Interactive|
|`./typicalw redis`|Redis Tool|
|`./typicalw redis console`|Redis Interactive|
