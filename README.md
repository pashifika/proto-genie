# ProtoGenie

ProtoGenie automates the creation of boilerplate code from gRPC proto files, jumpstarting your microservice development.
Focus on your core logic while ProtoGenie generates consistent, ready-to-use service templates for efficient gRPC workflows.

# Generate config[YAML]

| Level 1    | Level 2       | Level 3  | Description                                                                       |
|:-----------|:--------------|:---------|:----------------------------------------------------------------------------------|
| version    |               |          | Current generator configuration file version,<br>set together with template file. |
| settings   |               |          | Control the overall behavior of the main generator.                               |
|            | template_root |          | Directory path with template files.                                               |
| generators |               |          | Control the overall behavior of the template generator.                           |
|            | slice         |          |                                                                                   |
|            |               | name     | Rule name.                                                                        |
|            |               | type     | Now support the: message / service                                                |
|            |               | template | Relative path of template file                                                    |
|            |               | output   | Output file path.                                                                 |
