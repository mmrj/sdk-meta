# LaunchDarkly Log Codes

## Introduction 

Log codes provide a standardized way to reference different log conditions across LaunchDarkly SDKs.

## Codes

### configuration

Conditions associated with SDK configuration.

#### 0:3:0 - wrongType

This message indicates that a configuration option was not of the correct type. This primarily applies to languages that are not strongly typed.

| code | system | class |
|------|--------|-------|
| 0:3:0 | configuration | usageError |

##### Message

`Config option "${name}" should be of type ${expectedType}, but received ${actualType}, using the default value (${defaultValue}).`

| parameter | description |
|-----------|-------------|
| defaultValue | The default value of the configuration option. |
| expectedType | The correct type for the configuration option. |
| name | The name of the configuration option. |
| actualType | The incorrect types used for the configuration option. |


### Troubleshooting

This is a step that you need to troubleshoot!
#### 0:3:1 - unknownOption

This message indicates that an unrecognized configuration option was provided. This primarily applied to languages that are not strongly typed.

| code | system | class |
|------|--------|-------|
| 0:3:1 | configuration | usageError |

##### Message

`Ignoring unknown config option "${name}"`

| parameter | description |
|-----------|-------------|
| name | The option that was not recognized. |



### polling

Conditions related to polling.

#### 1:3:0 - badPollThing

You did a bad polling thing.

| code | system | class |
|------|--------|-------|
| 1:3:0 | polling | usageError |

##### Message

`The ${blank} was how ${bad}.`

| parameter | description |
|-----------|-------------|
| bad | The bad thing. |
| blank | The blank thing. |




