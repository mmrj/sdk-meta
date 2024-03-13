# sdk-meta

This repo contains metadata related to LaunchDarkly SDKs. 

The data is intended for consumption by downstream products and services.

| Data Product                             | Description                                                     | Format  |
|------------------------------------------|-----------------------------------------------------------------|---------|
| [Database](./metadata.sqlite3)           | Database containing data from which other products are derived. | sqlite3 |
| [SDK Names](products/names.json)         | SDK friendly names for display.                                 | JSON    |
| [SDK Types](products/types.json)         | SDK types for categorization.                                   | JSON    |
| [SDK Features](products/features.json)   | SDK features, including version introduced/deprecated.          | JSON    |
| [SDK Languages](products/languages.json) | Programming languages associated with SDKs.                     | JSON    |
| [SDK Repos](products/repos.json)         | SDK source repositories                                         | JSON    |


## structure

This repo is essentially a JSON database hosted on Github. 

The "tables" live in [`data`](./products) and the schemas for those tables live in [`schemas`](./schemas). When adding
a new table, ensure it has a corresponding schema.

Data can be validated against the schemas using `./scripts/ci/check-json-schemas.sh` on Linux.

Ensure that the JSON files are valid and formatted using `./scripts/ci/format-json.sh`.
