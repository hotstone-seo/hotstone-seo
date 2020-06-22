[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)
![HotStone Build Status](https://github.com/hotstone-seo/hotstone-seo/workflows/HotStone/badge.svg)
![HotStoneUI Build Status](https://github.com/hotstone-seo/hotstone-seo/workflows/HotStoneUI/badge.svg)
![HotStoneClient Build Status](https://github.com/hotstone-seo/hotstone-seo/workflows/HotStoneClient/badge.svg)

# Hotstone SEO

Hotstone is a service that let you decide how meta information for each of your web page should look like. It works by using a set of rules that you define: what meta tag to render whenever my web application navigate to a certain page. A Hotstone client in your web application will then make sure that the appropriate meta tags is being retrieved.

## Getting Started

This section will guide you to get the project up and running both for development and production.

### Requirements

This project uses **Golang (>1.12)** and **NodeJS (v13)**, so make sure to have both installed on your machine to build the project.

To run this project, you need both **PostgreSQL** and **Redis** available in the runtime environment.

(Optional) You can also install Docker to let Hotstone instantiate its own PostgreSQL and Redis using its provided build tool.

### How to Run

Hotstone uses [Typical-Go](https://github.com/typical-go/typical-go) as the build tool for this project. You can observe all the available command by typing `./typical help` from the project's root directory.

#### Local Environment

To run Hotstone in local environment, it is recommended that you use Docker and let typical-go run the containers for the required database and caching services.

``` bash
./typicalw docker up        # Equivalent with `docker-compose up -d` if the container is not already exists
./typicalw main-db create   # 
./typicalw main-db migrate  # Create and migrate postgres database

./typicalw test             # Run tests

./typicalw run              # Run the application server
./typicalw ui start         # Run the dashboard UI
```

#### Production Environment

For production environment, you might want to manage the build artifact yourself and run it using your own database infrastructure. To do this, you can do the following.

1. Build the binary for Hotstone server using `./typicalw compile` command:

2. Bundle the dashboard UI by running the command `npm build` on the `ui/` directory

3. Run the built artifacts according to your convention

### Configuration

Hotstone only supports using environment variables as runtime configuration. The required variables can be observed at the `.env.template` file.
