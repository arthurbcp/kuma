# Kuma

Kuma é um poderoso framework projetado para gerar estruturas de projetos para qualquer linguagem de programação com base em [templates Go](https://pkg.go.dev/text/template). Ele agiliza o processo de configuração de novos projetos, automatizando a criação de diretórios, arquivos e código base, garantindo consistência e economizando tempo valioso de desenvolvimento.

### Funcionalidades:

- Estruture os diretórios e arquivos do seu projeto de maneira customizada através de [templates Go](https://pkg.go.dev/text/template)
- Integração com Github para poder baixar templates pré definidos pela comunidade ou para uso pessoal através de repositórios privados
- Possibilidade de criar workflows de comandos cli personalizados através de um arquivo YAML usando as runs
- Utilizar variáveis de forma dinâmica para serem aplicadas aos templetes. As variáveis podem ser extraídas de um arquivo YAML ou JSON que pode ser local ou baixado de uma URL pública. As variáveis também podem ser obtidas através de informações passadas pelo usuário durante a execução de uma run

### Crie seus próprios boilerplates

Para o funcionamento do Kuma, todos arquivos relacionados ao framework devem ficar dentro da pasta `.kuma`

O framework funciona com um parser de [templates Go](https://pkg.go.dev/text/template) que são divididos em 3 tipos de arquivos

- **Builders**
  Templates no formato YAML que contém a estrutura de pastas e arquivos a serem criados
  &nbsp;

  ```yaml
  # .kuma/base.yaml

  # variables:
  # packageName = github.com/arthurbcp/kuma-hello-world
  # serviceName = customPrint
  # msg = Hello, Kuma!

  #  Global variables to be used in all the templates.
  global:
    packageName: "{{.data.packageName}}"

  # Structure of folders and files that will be generated
  structure:

    # src
    src:

      # src/services
      services:

        # src/services/custom_print.go
      {{toSnakeCase .data.serviceName}}.ts:
        # o template que será usado com base para a criação do arquivo custom_print.go
        template: templates/Service.go

        # variáveis que serão utilizadas dentro do template
        data:
        # name: CustomPrint
          name: {{toPascalCase .data.serviceName}}

    # main.go
    main.go:
        # o template que será usado com base para a criação do arquivo main.go
        template: templates/Main.go

      # variáveis que serão utilizadas dentro do template
      data:
        # name: CustomPrint
        service: {{toPascalCase .data.serviceName}}

        # msg: Hello, Kuma!
        msg: {{.data.msg}}
  ```

  &nbsp;&nbsp;

- **Templates**
  [Templates Go](https://pkg.go.dev/text/template) individuais para os arquivos que serão criados
  &nbsp;

  ```go
  // .kuma/templates/Main.go
  package main

  import (
  "fmt"
  //github.com/arthurbcp/kuma-hello-world/src/services/custom_print.go
  "{{.global.packageName}}/src/services/{{.data.service}}"
  )

  func main() {
  service := new{{data.service}}Service()
      // service.Print("Hello, Kuma!")
      service.Print("{{.data.msg}}")
  }

  ```

  ```go
  // .kuma/templates/Service.go
  package services

  import (
  "fmt"
  )

  // type CustomPrint struct {
  type {{.data.name}}Service struct {
  Print(msg string)
  }

  // func (s *CustomPrintService) Print(msg string) {
  func (s \*{{.data.name}}Service) Print(msg string) {
  fmt.Print(msg)
  }

  ```

- **Runs**

#### Executar uma RUN

O comando `exec` é usado para iniciar o processo de uma run.

```
kuma-cli exec --run=initial
```

**Flags**
`--run`, `-r`: Nome da run que será executada

#### Obter template do GitHub

Obtenha templates e runs através de um repositório GitHub.

Você pode obter os templates de qualquer repositório através do comando

```
kuma get --repo=github.com/arthurbcp/typescript-rest-openapi-services
```

Ou utilizar um de nosso templates oficiais com

```
kuma get --template=typescript-rest-openapi-services
```

**Flags**
`--repo`, `-r`: Nome do repositório GitHub
`--template`, `-t`: Nome do template oficial

##### Templates oficiais:

- **[OpenAPI 2.0 TypeScript services](github.com/arthurbcp/typescript-rest-openapi-services):** Crie uma library TypeScript com serviços tipados para todos os endpoints descrito em um arquivo Open API 2.0

```

```
