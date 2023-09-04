# Configuration

On command start running, `falco` finds up `.falco.yml` file from the current directory.
If the file is found, load and set to CLI configuration.

## Configuration File Structure

Here is a full configuration file example:

```yaml
// .falco.yaml

## Basic configurations
include_paths: [".", "/path/to/include"] 
remote: true
max_backends: 5
max_acls: 1000

## Linter configurations
linter:
  verbose: warning
  rules:
    acl/syntax: error

## Simulator configuration
simulator:
  port: 3124
  max_backends: 100
  max_acls: 100

## Testing configuration
testing:
  timeout: 100
  max_backends: 100
  max_acls: 100

## Backend Overrides
override_backends:
  F_httpbin_org:
    host: example.com
    ssl: true
    unhealthy: true
```

falco cascades each setting from the order of `Default Setting` -> `Configuration File` -> `CLI Arguments` to override.
All configurations of configuration files and CLI arguments are described following table:

| Configuration Field                | Type          | Default | CLI Argument       | Description                                                                                                                |
|:-----------------------------------|:-------------:|:-------:|:------------------:|:--------------------------------------------------------------------------------------------------------------------------|
| include_paths                      | Array<String> | []      | -I, --include_path | Include VCL paths                                                                                                         |
| remote                             | Boolean       | false   | -r, --remote       | Fetch remote resources of Fastly                                                                                          |
| max_backends                       | Integer       | 5       | --max_backends     | Override Fastly's backend amount limitation                                                                               |
| max_acls                           | Integer       | 1000    | --max_acls         | Override Fastly's acl amount limitation                                                                                   |
| simulator                          | Object        | null    | -                  | Simulator configuration object                                                                                            |
| simulator.port                     | Integer       | 3124    | -p, --port         | Simulator server listen port                                                                                              |
| testing                            | Object        | null    | -                  | Testing configuration object                                                                                              |
| testing.timeout                    | Integer       | 10      | -t, --timeout      | Set timeout to stop testing                                                                                               |
| linter                             | Object        | null    | -                  | Override linter rules                                                                                                     |
| linter.verbose                     | String        | error   | -v, -vv            | Verbose level, `warning` or `info` is valid                                                                               |
| linter.rules                       | Object        | null    | -                  | Override linter rules                                                                                                     |
| linter.rules.[rule_name]           | String        | -       | -                  | Override linter error level for the rule name, see [rules](https://github.com/ysugimoto/falco/blob/develop/docs/rules.md) |
| override_backends                  | Object        | -       | -                  | Override backend settings in main VCL which correspond to the name                                                        |
| override_backends.[name]           | Object        | -       | -                  | Backend name to override                                                                                                  |
| override_backends.[name].host      | String        | -       | -                  | Backend host to override                                                                                                  |
| override_backends.[name].ssl       | Boolean       | true    | -                  | Use HTTPS when set `true`                                                                                                 |
| override_backends.[name].unhealthy | Boolean       | false   | -                  | Override backend to be unhealthy when set `true`                                                                          |





