# sdk-meta

This repo contains metadata related to LaunchDarkly SDKs. 

The data is intended for consumption by downstream products and services.

| Data Product                                  | Description                                                     | Format  |
|-----------------------------------------------|-----------------------------------------------------------------|---------|
| [Database](./metadata.sqlite3)                | Database containing data from which other products are derived. | sqlite3 |
| [SDK Names](products/json/names.json)         | SDK friendly names for display.                                 | JSON    |
| [SDK Releases](products/json/releases.json)   | SDK major/minor releases with EOL dates.                        | JSON    |
| [SDK Types](products/json/types.json)         | SDK types for categorization.                                   | JSON    |
| [SDK Features](products/json/features.json)   | SDK features, including version introduced/deprecated.          | JSON    |
| [SDK Languages](products/json/languages.json) | Programming languages associated with SDKs.                     | JSON    |
| [SDK Repos](products/json/repos.json)         | SDK source repositories                                         | JSON    |


## structure

This repo contains an sqlite database containing a snapshot of SDK metadata
fetched from individual repos.

It also contains JSON files that are derived from the database. These are intended for
consumption by downstream products and services.

The JSON data products live in [`products`](./products) and the schemas for them live in [`schemas`](./schemas). 

## tooling

Data products can be validated against the schemas using `./scripts/ci/check-json-schemas.sh` on Linux.

Ensure that the JSON files are valid and formatted using `./scripts/ci/format-json.sh`.

To quickly create basic metadata for a new repo, run `./scripts/add-repo.sh <launchdarkly/repo-name>`. This will
clone the repo, check-in an `.sdk_metadata.json` file, and create a PR.
