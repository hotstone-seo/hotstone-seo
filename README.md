# hotstone-seo

[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)
![HotStone Build Status](https://github.com/hotstone-seo/hotstone-seo/workflows/HotStone/badge.svg)
![HotStoneUI Build Status](https://github.com/hotstone-seo/hotstone-seo/workflows/HotStoneUI/badge.svg)
![HotStoneClient Build Status](https://github.com/hotstone-seo/hotstone-seo/workflows/HotStoneClient/badge.svg)

Tag Management tool that support the server-side rendering and data sourcing.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites

For MacOS:
1. Install [Homebrew](https://brew.sh/)
2. Install [Go](https://golang.org/doc/install)
    ```bash
    brew install go
    ```
3. Install [NodeJS](https://nodejs.org/en/)
    ```bash
    brew install node
    ```
4. For editor, it is recommended to use [Visual Code](https://code.visualstudio.com/Download)
    - Install the recommended extensions

### Run 

Make sure you have a configuration set, use the template from .env.template

```bash
./typicalw docker up        # equivalent with `docker-compose up -d` (if infrastructure not up)
./typicalw main-db create   # 
./typicalw main-db migrate  # create and migrate postgres database

./typicalw mock             # generate mock (if require mock)
./typicalw test             # run test 

./typicalw run              # run the application
