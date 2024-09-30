# Kuma CLI

<!--![Kuma Logo](path_to_your_logo.png)  Replace with your actual logo path if available -->

Kuma CLI is a powerful command-line tool designed to generate project scaffolds based on Go templates. It streamlines the process of setting up new projects by automating the creation of directories, files, and boilerplate code, ensuring consistency and saving valuable development time.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  - [Using Go Install](#using-go-install)
  - [Downloading Precompiled Binaries](#downloading-precompiled-binaries)
- [Usage](#usage)
  - [Generate a Project Scaffold](#generate-a-project-scaffold)
    - [Flags](#flags)
    - [Example](#example)
  - [Parsing Configuration Files](#parsing-configuration-files)
    - [Supported Parsers](#supported-parsers)
    - [Flags](#flags-1)
    - [Example](#example-1)
- [Configuration](#configuration)
  - [KumaConfig File](#kumaconfig-file)
    - [Structure](#structure)
    - [Example](#example-2)
- [Example](#example-3)
  - [Go Template Example](#go-template-example)
    - [Explanation](#explanation)
  - [KumaConfig File Example](#kumaconfig-file-example)
    - [Explanation](#explanation-1)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Automated Scaffold Generation**: Quickly generate project structures with predefined templates.
- **Flexible Parsing**: Supports multiple parsers like OpenAPI to process various configuration files.
- **Customizable Templates**: Easily define and modify Go templates to suit your project's needs.
- **Extensible Architecture**: Add new parsers and templates with minimal effort.

## Prerequisites

Before installing and using Kuma CLI, ensure you have the following prerequisites met:

- **Go**: Version 1.16 or higher. [Download Go](https://golang.org/dl/)
- **Git**: For cloning the repository and version control. [Download Git](https://git-scm.com/downloads)

## Installation

You can install Kuma CLI using \`go install\` or by downloading precompiled binaries.

### Using Go Install

#### Set Go Environment Variables (if not already set):

Ensure that your \`GOPATH\` and \`GOBIN\` are correctly set. You can add \`$GOBIN\` to your \`PATH\` to access the \`kuma\` command globally.

```bash
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

#### Install Kuma CLI:

```bash
go install github.com/arthurbcp/kuma-cli@latest
```

This command fetches the latest version of Kuma CLI and installs it to your \`$GOBIN\` directory.

### Downloading Precompiled Binaries

If you prefer not to build from source, you can download precompiled binaries from the [Releases](https://github.com/arthurbcp/kuma-cli/releases) page.

#### Navigate to the Releases Page:

Go to [Kuma CLI Releases](https://github.com/arthurbcp/kuma-cli/releases) and download the appropriate binary for your operating system.

#### Extract and Install:

Extract the downloaded archive and move the \`kuma\` executable to a directory in your \`PATH\`, such as \`/usr/local/bin\`.

```bash
tar -xzf kuma-cli-linux-amd64.tar.gz
sudo mv kuma /usr/local/bin/
```

## Usage

After installation, you can start using Kuma CLI to generate project scaffolds.

### Generate a Project Scaffold

The \`generate\` command is used to create a new project structure based on your configuration and templates.

```bash
kuma generate [flags]
```

#### Flags

- \`--parser\`, \`-p\`: Specify the parser helper to use (e.g., \`openapi\`).
- \`--p-file\`: Path to the file you want to parse.
- \`--config\`, \`-c\`: Path to the Kuma config file. (Default: \`kuma-config.yaml\`)
- \`--project-path\`, \`-p\`: Path to the directory where the project will be generated. (Default: \`kuma-generated\`)
- \`--templates-path\`, \`-t\`: Path to the directory containing Kuma templates. (Default: \`kuma-templates\`)

#### Example

```bash
kuma generate --parser openapi --p-file api-spec.yaml --config kuma-config.yaml --project-path ./myproject --templates-path ./templates
```

### Parsing Configuration Files

The \`parse\` command allows you to process configuration files using specified parsers.

```bash
kuma parse [parser] [flags]
```

#### Supported Parsers

- \`openapi\`: Parses OpenAPI specification files.

#### Flags

- \`--file\`, \`-f\`: Path to the file you want to parse. **(Required)**

#### Example

```bash
kuma parse openapi --file api-spec.yaml
```

## Configuration

Kuma CLI relies on a configuration file (\`KumaConfig\`) to define project settings and structure. Below is an overview of how to set up and customize your \`kuma-config.yaml\`.

### KumaConfig File

The \`kuma-config.yaml\` file specifies the project name, repository, structure, and templates to be used during scaffold generation.

#### Structure

- **Config**: General project configurations like project name and repository.
- **Structure**: Defines the directory and file hierarchy along with associated templates and data.

#### Example

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

## Example

To illustrate how Kuma CLI works, let's walk through an example using a Go template and a \`kuma-config.yaml\` file.

### Go Template Example

Below is an example of a Go template used by Kuma CLI to generate a service file:

```go
import { HttpProvider } from '../providers/HttpProvider';
{{ template "Index.gtpl" }}
export class {{ .Data.Name }}Service {
  private http: HttpProvider;

  constructor(http: HttpProvider) {
    this.http = http;
  }

  {{ range .Data.Endpoints }}
  {{ if .Description }}//{{ .Description }}{{end}}
  {{ .Name }}(data?: any): Promise<any> {
    return this.http.{{ toLower .HttpMethod }}<any>("{{ .Route }}", data);
  }
  {{ end }}
}
```

#### Explanation

- **Imports**: Imports necessary dependencies.
- **Template Inclusion**: Includes another template named \`Index.gtpl\`.
- **Class Definition**: Defines a service class with a constructor and methods.
- **Endpoints**: Iterates over endpoints defined in the configuration to generate methods dynamically.

### KumaConfig File Example

The \`kuma-config.yaml\` file defines how the project structure should be generated.

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
        user:
          service.ts:
            Template: Service.gtpl
            Includes:
              - Index.gtpl
            Data:
              Name: User
              FileName: user_service.ts
              Endpoints:
                - name: GetUser
                  description: Retrieves a user by ID
                  httpMethod: GET
                  route: "/users/{id}"
                - name: CreateUser
                  description: Creates a new user
                  httpMethod: POST
                  route: "/users"
          service_interface.ts:
            Template: ServiceInterface.gtpl
            Data:
              Name: User
              FileName: user_service_interface.ts

          service_mock.ts:
            Template: ServiceMock.gtpl
            Data:
              Name: User
              FileName: user_service_mock.ts

      dto:
        user_dto.ts:
          Template: DTO.gtpl
  index.ts:
    Template: Index.gtpl
  package.json:
    Template: PackageJson.gtpl
  tsconfig.json:
    Template: TsConfig.gtpl
```

#### Explanation

- **Config**: Sets the project name and repository URL.
- **Structure**: Defines the directory and file hierarchy.
  - **Providers**: Specifies templates for HTTP providers.
  - **Services**: Dynamically generates service classes based on controllers.
    - **Includes**: Includes additional templates like \`Index.gtpl\`.
    - **Data**: Provides dynamic data such as service name, filename, and endpoints.
  - **DTO**: Generates Data Transfer Objects based on components.
- **Root Files**: Generates essential root files like \`index.ts\`, \`package.json\`, and \`tsconfig.json\` using respective templates.

## Contributing

Contributions are welcome! Please follow these steps to contribute to Kuma CLI:

1. **Fork the Repository**:

   Click the **Fork** button at the top-right corner of the repository page.

2. **Clone Your Fork**:

   ```bash
   git clone https://github.com/yourusername/kuma-cli.git
   cd kuma-cli
   ```

3. **Create a New Branch**:

   ```bash
   git checkout -b feature/your-feature-name
   ```

4. **Make Your Changes**:

   Implement your feature or fix.

5. **Commit Your Changes**:

   ```bash
   git commit -m "Add feature: your-feature-description"
   ```

6. **Push to Your Fork**:

   ```bash
   git push origin feature/your-feature-name
   ```

7. **Create a Pull Request**:

   Navigate to the original repository and click on **"Compare & pull request"**.

## License

Distributed under the MIT License. See \`LICENSE\` for more information.
`;
