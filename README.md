# ops

[![Dockerhub pulls](https://img.shields.io/docker/pulls/unfor19/ops)](https://hub.docker.com/r/unfor19/ops)

**Work In Progress**- An attempt to create an application that serves as a one-stop-shop for DevOps Engineers.

## Features

- Download a file efficiently with [Accept-Ranges](https://developer.mozilla.org/en-US/docs/Web/HTTP/Range_requests)
- Unzip file

## Requirements

- [Go](https://golang.org/doc/install) v1.16.x

## Usage

- Cross platform support - macOS, Windows and Linux.
- Check the [Releases](https://github.com/unfor19/ops/releases) page to download the binary file

### Available commands

<!-- available_commands_start -->

```
A CLI for DevOps Engineers

Usage:
  ops [command]

Available Commands:
  download    Download a file efficiently
  help        Help about any command
  unzip       

Flags:
      --config string   config file (default is /.ops.yaml)
  -h, --help            help for ops
  -t, --toggle          Help message for toggle

Use "ops [command] --help" for more information about a command.
```

<!-- available_commands_end -->

## Docker

### Run

- Latest version
  ```bash
  docker run --rm -it unfor19/ops
  ```
- Specific version - Check the [Releases](https://github.com/unfor19/ops/releases) to see available versions
  ```bash
  docker run --rm -it unfor19/ops:0.0.11rc
  # docker run --rm -it unfor19/ops:0.0.11rc-scratch
  ```

### Build Locally

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
1. Run
   ```bash
   docker run --rm -it unfor19/ops
   ```

## Contributing

Report issues/questions/feature requests on the [Issues](https://github.com/unfor19/ops/issues) section.

Pull requests are welcome! Ideally, create a feature branch and issue for every single change you make. These are the steps:

1. Fork this repo
1. Create your feature branch from master (`git checkout -b my-new-feature`)
1. Get dependencies
   ```bash
   go mod download
   ```
1. Go Build
   ```bash
   go build
   # output file: ./ops
   ```
1. Add the code of your new feature
1. Commit your remarkable changes (`git commit -am 'Added new feature'`)
1. Push to the branch (`git push --set-up-stream origin my-new-feature`)
1. Create a new Pull Request and tell us about your changes

## Authors

Created and maintained by [Meir Gabay](https://github.com/unfor19)

This project is based on - [jacklin293/golang-parallel-download-with-accept-ranges](https://github.com/jacklin293/golang-parallel-download-with-accept-ranges)

## License

This project is licensed under the Apache License - see the [LICENSE](https://github.com/unfor19/ops/blob/master/LICENSE) file for details