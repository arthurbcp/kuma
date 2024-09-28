# Kuma CLI

<!--![Kuma Logo](path_to_your_logo.png)  Replace with your actual logo path if available -->

**Kuma CLI** is a powerful command-line tool designed to generate project scaffolds based on Go templates. It streamlines the process of setting up new projects by automating the creation of directories, files, and boilerplate code, ensuring consistency and saving valuable development time.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
  - [Generate a Project Scaffold](#generate-a-project-scaffold)
  - [Parsing Configuration Files](#parsing-configuration-files)
- [Configuration](#configuration)
  - [KumaConfig File](#kumaconfig-file)
- [Example](#example)
  - [Go Template Example](#go-template-example)
  - [KumaConfig File Example](#kumaconfig-file-example)
- [Contributing](#contributing)
- [License](#license)

## Features

- Automated Scaffold Generation: Quickly generate project structures with predefined templates.
- Flexible Parsing: Supports multiple parsers like OpenAPI to process various configuration files.
- Customizable Templates: Easily define and modify Go templates to suit your project's needs.
- Extensible Architecture: Add new parsers and templates with minimal effort.

## Prerequisites

Before installing and using Kuma CLI, ensure you have the following prerequisites met:

- **Go**: Version 1.16 or higher. [Download Go](https://golang.org/dl/)
- **Git**: For cloning the repository and version control. [Download Git](https://git-scm.com/)

## Installation

You can install Kuma CLI using `go install` or by downloading precompiled binaries.

### Using Go Install

Set Go Environment Variables (if not already set):

```bash
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

Install Kuma CLI:

```bash
go install github.com/arthurbcp/kuma-cli@latest
```

### Downloading Precompiled Binaries

If you prefer not to build from source, you can download precompiled binaries from the [Releases](https://github.com/arthurbcp/kuma-cli/releases) page.

Extract the downloaded archive and move the kuma executable to a directory in your PATH:

```bash
tar -xzf kuma-cli-linux-amd64.tar.gz
sudo mv kuma /usr/local/bin/
```

## Usage

After installation, you can start using Kuma CLI to generate project scaffolds.

### Generate a Project Scaffold

The `generate` command is used to create a new project structure based on your configuration and templates.

```bash
kuma generate [flags]
```

#### Flags

- `--parser, -p`: Specify the parser helper to use (e.g., openapi).
- `--p-file`: Path to the file you want to parse.
- `--config, -c`: Path to the Kuma config file. (Default: "kuma-config.yaml")
- `--project-path, -p`: Path to the directory where the project will be generated. (Default: "kuma-generated")
- `--templates-path, -t`: Path to the directory containing Kuma templates. (Default: "kuma-templates")

#### Example

```bash
kuma generate --parser openapi --p-file api-spec.yaml --config kuma-config.yaml --project-path ./myproject --templates-path ./templates
```

### Parsing Configuration Files

The `parse` command allows you to process configuration files using specified parsers.

```bash
kuma parse [parser] [flags]
```

#### Supported Parsers

- `openapi`: Parses OpenAPI specification files.

#### Flags

- `--file, -f`: Path to the file you want to parse. (Required)

#### Example

```bash
kuma parse openapi --file api-spec.yaml
```

## Configuration

Kuma CLI relies on a configuration file (`KumaConfig`) to define project settings and structure. Below is an overview of how to set up and customize your `kuma-config.yaml`.

### KumaConfig File

The `kuma-config.yaml` file specifies the project name, repository, structure, and templates to be used during scaffold generation.

### Example

```yaml
Config:
  ProjectName: "myproject"
  ProjectRepository: "github.com/mycompany/myproject"

Structure:
  src:
    generated:
      providers:
        http:
          http_provider.ts:
            Template: HttpProvider.gtpl
          http_provider_interface.ts:
            Template: HttpProviderInterface.gtpl
          http_mock_provider.ts:
            Template: HttpMockProvider.gtpl
      services:
        {{range .Controllers }}{{ toSnake .Name }}:
          service.ts:
            Template: Service.gtpl
            Includes:
             - Index.gtpl
            Data:
              Name: {{ .Name }}
              FileName: {{ toSnake .Name }}_service.ts
              Endpoints:
                {{ range yaml .Endpoints }}{{ . }}
                {{ end }}
          service_interface.ts:
            Template: ServiceInterface.gtpl
            Data:
              Name: {{ .Name }}
              FileName: {{ toSnake .Name }}_service_interface.ts

          service_mock.ts:
            Template: ServiceMock.gtpl
            Data:
              Name: {{ .Name }}
              FileName: {{ toSnake .Name }}_service_mock.ts

        {{end}}
      dto:
        {{ range .Components }}{{ toSnake .Name }}.ts:
          Template: DTO.gtpl

        {{end}}
  index.ts:
    Template: Index.gtpl
  package.json:
    Template: PackageJson.gtpl
  tsconfig.json:
    Template: TsConfig.gtpl
```

## Contributing

Contributions are welcome! Please follow these steps to contribute to Kuma CLI:

1. Fork the Repository:
2. Clone Your Fork:
   ```bash
   git clone https://github.com/yourusername/kuma-cli.git
   cd kuma-cli
   ```
3. Create a New Branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```
4. Make Your Changes:
5. Commit Your Changes:
   ```bash
   git commit -m "Add feature: your-feature-description"
   ```
6. Push to Your Fork:
   ```bash
   git push origin feature/your-feature-name
   ```
7. Create a Pull Request:

## License

Distributed under the MIT License. See `LICENSE` for more information.
