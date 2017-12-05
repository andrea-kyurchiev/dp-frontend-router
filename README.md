dp-frontend-router
==================

### Configuration

| Environment variable          | Default                                 | Description
| ----------------------------- | --------------------------------------- | --------------------------------------
| BIND_ADDR                     | :20000                                  | The host and port to bind to.
| BABBAGE_URL                   | https://www.ons.gov.uk                  | The URL of the babbage instance to use
| RESOLVER_URL                  | http://localhost:20020                  | The URL of dp-content-resolver
| RENDERER_URL                  | http://localhost:20010                  | The URL of dp-frontend-renderer
| DATASET_CONTROLLER_URL        | http://localhost:20200                  | The URL of dp-frontend-dataset-controller
| FILTER_DATASET_CONTROLLER_URL | http://localhost:20001                  | The URL of dp-frontend-filter-dataset-controller
| ZEBEDEE_URL                   | http://localhost:8082                   | The URL of zebedee
| PATTERN_LIBRARY_ASSETS_PATH   | https://cdn.ons.gov.uk/sixteens/6cc1837 | The URL to the sixteens build to use
| SITE_DOMAIN                   | ons.gov.uk                              | The domain hosting the site
| SPLASH_PAGE                   |                                         | The splash page template to use
| DISABLED_PAGE                 |                                         | The disabled page template to use
| HOMEPAGE_AB_PERCENT           | 0                                       | Percentage of users who get version B
| DEBUG                         | false                                   | Whether to enable debug mode

### Licence

Copyright ©‎ 2016 - 2017, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
