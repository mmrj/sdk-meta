TO create an equivalent of sdkVersions.json for the docs site:
FOR each SDK in data/sdks.json
  GET all the associated repo and SDK ID
  IF monorepo:
    GET all releases for that repo + ID
  ELSE:
    GET all releases for that repo
  GET the u2c version from data/features.json
  LET sdk-name be SDK.name
  CREATE JSON OBJECT with properties:
    name: $sdk-name
    u2cVersion: $u2c-version
    versions: [
      Object {
        versionMajorMinor: parse the semantic version to drop the patch
        releaseYear: parse the date to get hte releaseYear
        releaseMonthDay: parse the date to get the month/day
      }
    ]
