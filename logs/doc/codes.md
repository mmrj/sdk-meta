# configuration

Conditions associated with SDK configuration.

## unknownOption

This message indicates that an unrecognized configuration option was provided. This primarily applied to languages that are not strongly typed.

| code | system | class |
|------|--------|-------|
| 0:3:2 | configuration | usageError |

### Message

`Ignoring unknown config option "${name}"`

| parameter | description |
|-----------|-------------|
| name | The option that was not recognized. |


## wrongType

This message indicates that a configuration option was not of the correct type. This primarily applies to languages that are not strongly typed.

| code | system | class |
|------|--------|-------|
| 0:3:1 | configuration | usageError |

### Message

`Config option "${name}" should be of type ${expectedType}, but received ${actualType}, using the default value (${defaultValue}).`

| parameter | description |
|-----------|-------------|
| actualType | The incorrect types used for the configuration option. |
| defaultValue | The default value of the configuration option. |
| expectedType | The correct type for the configuration option. |
| name | The name of the configuration option. |



# polling

This is responsible for polling.


# streaming

Responsible for handling the streaming connection to LaunchDarkly.


