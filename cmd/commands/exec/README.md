# Documentação das Runs no Kuma

## Índice

- [Introdução](#introdução)
- [Estrutura de uma Run](#estrutura-de-uma-run)
  - [Tipos de Ações](#tipos-de-ações)
    - [Input](#input)
    - [Log](#log)
    - [Create](#create)
    - [Cmd](#cmd)
    - [Load](#load)
    - [Run Aninhada](#run-aninhada)
- [Como Executar uma Run](#como-executar-uma-run)
  - [Usando o Comando CLI](#usando-o-comando-cli)
  - [Seleção Interativa de Runs](#seleção-interativa-de-runs)
- [Exemplos de Runs](#exemplos-de-runs)
  - [Run que extraí as variáveis de um arquivo swagger](#run-que-extraí-as-variáveis-de-um-arquivo-swagger)
- [Licença](#licença)

## Introdução

As **Runs** no Kuma são sequências de ações definidas para automatizar tarefas repetitivas durante o desenvolvimento de projetos. Elas permitem a execução de pipelines personalizados que podem incluir desde inputs de usuário até a execução de comandos no terminal, garantindo consistência e eficiência no fluxo de trabalho.

## Estrutura de uma Run

Uma Run é composta por uma sequência de passos que definem as ações a serem executadas. A seguir, detalhamos os principais componentes e tipos de ações que podem ser incluídos em uma Run.

### Tipos de Ações

#### Input

Solicita uma entrada do usuário durante a execução da Run.

```yaml
- input:
    label: "Qual é o nome do pacote do seu projeto?"
    out: packageName # Exemplo: github.com/arthurbcp/kuma-hello-world
```

**Campos:**

- `label`: A mensagem exibida para o usuário.
- `out`: A variável onde o valor inserido será armazenado.

**Opções Adicionais:**

- `options`: Lista de opções para seleção.
- `multi`: Flag que permite selecionar mais de uma opção. Retorna um array no `out`.
- `other`: Se nenhuma option for selecionada, exibe um atalho com a tecla **o** para abrir um input de texto
  **Exemplo com Opções e Seleção Múltipla:**

```yaml
select-runtime:
  description: "Select the runtime to use"
  steps: - input:
    label: "Select a runtime"
    multi: false
    other: false
    options:
      - label: Node
        value: node
      - label: Deno 2.0
        value: deno
      - label: Bun
        value: bun
    out: runtime
```

#### Log

Registra uma mensagem no console.

```yaml
- log: "Criando estrutura para {{.data.packageName}}" # Exemplo: Criando estrutura para github.com/arthurbcp/kuma-hello-world
```

**Campos:**

- `log`: A mensagem a ser registrada. Pode incluir variáveis dinâmicas.

#### Create

Cria a estrutura do projeto com base em um builder definido.

```yaml
- create:
    from: base.yaml
```

**Campos:**

- `from`: O arquivo YAML que define a estrutura e os templates a serem usados.

#### Cmd

Executa um comando no terminal.

```yaml
- cmd: npm install
```

**Campos:**

- `cmd`: O comando a ser executado. Pode incluir variáveis dinâmicas.

#### Load

Carrega variáveis a partir de um arquivo local ou URL.

```yaml
- load:
    from: variables.yaml
    out: vars
```

**Campos:**

- `from`: Caminho ou URL para o arquivo JSON ou YAML contendo a estrutura que será armazada dentro da variável `out`. Pode incluir variáveis dinâmicas.
- `out`: A variável onde os dados carregados serão armazenados.

#### Run Aninhada

Executa uma run dentro de outra run. As variáveis de uma run são passadas automaticamente para as runs aninhadas.

```yaml
main:
  description: "main run"
  steps:
    - log: "Executing main run"
    - run: nested

nested:
  description: "nested run"
  steps:
    - log: "Executing nested run"
```

**Campos:**

- `run`: Nome da Run a ser executada.

## Como Executar uma Run

### Usando o Comando CLI

Para executar uma Run específica, utilize o comando `exec` seguido do nome da Run.

```bash
kuma-cli exec --run=initial
```

**Flags:**

- `--run`, `-r`: Nome da Run que será executada.

### Seleção Interativa de Runs

Se o nome da Run não for especificado, o Kuma CLI apresentará uma interface interativa para selecionar qual Run deseja executar.

```bash
kuma-cli exec
```

**Passos:**

1. **Seleção de Run:** Uma lista de Runs disponíveis será exibida para seleção.
2. **Execução:** A Run selecionada será executada com base nos passos definidos.

## Exemplos avançados

### Run que extraí as variáveis de um arquivo swagger

Confira os builders e templates utilizados clonando esse [repositório](https://github.com/arthurbcp/typescript-rest-openapi-services) ou através do comando `kuma-cli get -t typescript-rest-openapi-services`

```yaml
initial:
  description: "Initial run"
  steps:
    - input:
        label: "What is your project name?"
        out: projectName
    - input:
        label: "What is your project repository?"
        out: projectRepository
    - create:
        from: base.yaml
    - log: Base structure created successfully!
    - input:
        label: "Enter the local file or the file URL in the Open API 2.0 format with the data you want to generate the library:"
        out: swagger
    - load:
        from: "{{.data.swagger}}"
        out: apiData
    - create:
        from: from-swagger.yaml
    - cmd: npm i
    - cmd: npm run format
```

## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
