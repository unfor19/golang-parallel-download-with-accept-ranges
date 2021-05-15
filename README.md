# ops

**Work In Progress**- An attempt to create an application that serves as a one-stop-shop for DevOps Engineers.

## Features

- Download a file using efficiently with [Accept-Ranges](https://developer.mozilla.org/en-US/docs/Web/HTTP/Range_requests)
- Unzip file

## Requirements

- [Go](https://golang.org/doc/install) v1.16.x

## Build

1. Clone
    ```bash
    git clone https://github.com/unfor19/ops.git && \
    cd op
    ```
1. Get dependencies
   ```bash
   go mod download
   ```
1. Go Build
   ```bash
   go build
   # output file: ./ops
   ```

## Usage

Available commands

<!-- available_commands_start -->

```
A CLI for DevOps Engineers

Usage:
  ops [command]

Available Commands:
  download    Download a file efficiently
  help        Help about any command

Flags:
      --config string   config file (default is /.ops.yaml)
  -h, --help            help for ops
  -t, --toggle          Help message for toggle

Use "ops [command] --help" for more information about a command.
```

<!-- available_commands_end -->

## Docker

1. Clone
    ```bash
    git clone https://github.com/unfor19/ops.git && \
    cd ops
    ```
1. Build
   ```bash
   docker build -t unfor19/ops .
   # docker build -t unfor19/ops -f Dockerfile.scratch .
   ```
2. Run
   ```bash
   docker run --rm -it unfor19/ops
   ```

## Authors

Created and maintained by [Meir Gabay](https://github.com/unfor19)

This project is based on - [jacklin293/golang-parallel-download-with-accept-ranges](https://github.com/jacklin293/golang-parallel-download-with-accept-ranges)

## License

This project is licensed under the Apache License - see the [LICENSE](https://github.com/unfor19/ops/blob/master/LICENSE) file for details