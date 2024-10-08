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
  # msg = Hello, Kuma!

  #  Global variables to be used in all the templates.
  global:
    packageName: "{{.data.packageName}}"

  # Structure of folders and files that will be generated
  structure:
    # main.go
    main.go:
      # o template que será usado com base para a criação do arquivo main.go
      template: templates/Main.go

      # variáveis que serão utilizadas dentro do template
      data:
        # msg: Hello, Kuma!
        msg: "{{ .data.msg }}"
  ```

  &nbsp;&nbsp;

- **Templates**
  [Templates Go](https://pkg.go.dev/text/template) individuais para os arquivos que serão criados
  &nbsp;

  ```go
  package main

  import (
    "fmt"
  )

  func main() {
    // fmt.Print("Hello, Kuma!")
    fmt.Println("{{ .data.msg }}")
  }

  ```

&nbsp;&nbsp;

- **Runs**
  Arquivo YAML contendo um sequencias de ações que serem executadas ao chamar uma `run`. Desde logs, comandos de terminal, chamadas http, inputs de texto e múltipla escolha para terminal e ações para criar pastas arquivos baseados nos builders e templates.
  **Confira a documentação completa aqui:**
  &nbsp;

  ```yaml
  # Nome da run que será executada assim que o repositório for obtido através do comando `kuma-cli get`
  initial:
    # descrição da run
    description: "Initial run"
    # passos que serão executados quando a run for chamada
    steps:
      # input action
      - input:
          label: "What is your project package name?"
          out: packageName #github.com/arthurbcp/kuma-hello-world
      # input actions
      - input:
          label: "What message do you want to print?" #
          out: msg #Hello, Kuma!

      # log a message
      - log: "creating structure for {{.data.packageName}}" # creating structure for github.com/arthurbcp/kuma-hello-world

      # create the project structore using base.yaml builder
      - create:
          from: base.yaml

      #log a message
      - log: Base structure created successfully!

      # init the go package
      - cmd: go mod init {{.data.packageName}} #go mod init github.com/arthurbcp/kuma-hello-world

      # install dependencies
      - cmd: go mod tidy

      #exec the main.go file
      - cmd: go run main.go # Hello, Kuma!
  ```

### Comandos de terminal

#### Criar um boilerplate

O comando `create` é usado para criar um boilerplate baseado nos builders e templates que estão dentro da pasta `.kuma` e um arquivo JSON ou YAML contendo as variáveis para fazer as substituição nos templetes.

```
kuma-cli create --v=swagger.json --project=. --from=base.yaml
```

**Flags**
`--variables`, `-v`: Path or URL to the variables file
`--project`, `-p`: Path to the project you want to create
`--from`, `-f`: Path to the YAML file with the structure and templates

&nbsp;

#### Executar uma RUN

O comando `exec` é usado para iniciar o processo de uma run.

```
kuma-cli exec --run=initial
```

**Flags**
`--run`, `-r`: Nome da run que será executada

&nbsp;

#### Obter templates do GitHub

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
&nbsp;

##### Templates oficiais:

- **[OpenAPI 2.0 TypeScript services](github.com/arthurbcp/typescript-rest-openapi-services):** Crie uma library TypeScript com serviços tipados para todos os endpoints descrito em um arquivo Open API 2.0
