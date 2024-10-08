# Kuma

Kuma é um poderoso framework projetado para gerar estruturas de projetos para qualquer linguagem de programação com base em [templates Go](https://pkg.go.dev/text/template). Ele agiliza o processo de configuração de novos projetos, automatizando a criação de diretórios, arquivos e código base, garantindo consistência e economizando tempo valioso de desenvolvimento.

## Tabela de Conteúdos

- [Funcionalidades](#funcionalidades)
- [Instalação](#instalação)
- [Crie seus Próprios Boilerplates](#crie-seus-próprios-boilerplates)
  - [Builders](#builders)
  - [Templates](#templates)
  - [Runs](#runs)
- [Comandos de Terminal](#comandos-de-terminal)
  - [Criar um Boilerplate](#criar-um-boilerplate)
  - [Executar uma Run](#executar-uma-run)
  - [Obter Templates do GitHub](#obter-templates-do-github)
    - [Templates Oficiais](#templates-oficiais)
- [Contribuição](#contribuição)
- [Licença](#licença)

## Funcionalidades

- Estruture os diretórios e arquivos do seu projeto de maneira customizada através de [templates Go](https://pkg.go.dev/text/template).
- Integração com GitHub para baixar templates pré-definidos pela comunidade ou para uso pessoal através de repositórios privados.
- Possibilidade de criar workflows de comandos CLI personalizados através de um arquivo YAML usando as runs.
- Utilização de variáveis de forma dinâmica para serem aplicadas aos templates. As variáveis podem ser extraídas de um arquivo YAML ou JSON que pode ser local ou baixado de uma URL pública. As variáveis também podem ser obtidas através de informações passadas pelo usuário durante a execução de uma run.

## Instalação

Para instalar o Kuma, utilize o comando `go install`. Siga os passos abaixo:

### Requisitos

- [Go](https://golang.org/dl/) versão 1.23 ou superior.
- Git instalado e configurado no seu sistema (necessário para `go install`).

### Passo a Passo

1. **Execute o comando de instalação:**

   ```bash
   go install github.com/arthurbcp/kuma@latest
   ```

2. **Verifique a instalação:**

   ```bash
   kuma-cli --help
   ```

   Você deve ver a ajuda do Kuma CLI, confirmando que a instalação foi bem-sucedida.

## Crie seus Próprios Boilerplates

Para o funcionamento do Kuma, todos os arquivos relacionados ao framework devem ficar dentro da pasta `.kuma`.

O framework utiliza um parser de [templates Go](https://pkg.go.dev/text/template) que são divididos em três tipos de arquivos:

### Builders

Templates no formato YAML que contêm a estrutura de pastas e arquivos a serem criados.

```yaml
# .kuma/base.yaml

# Variáveis globais a serem usadas em todos os templates.
global:
  packageName: "{{.data.packageName}}"

# Estrutura de diretórios e arquivos que serão gerados
structure:
  # main.go
  main.go:
    # O template que será usado como base para a criação do arquivo main.go
    template: templates/Main.go

    # Variáveis que serão utilizadas dentro do template
    data:
      # msg: Hello, Kuma!
      msg: "{{ .data.msg }}"
```

### Templates

[Templates Go](https://pkg.go.dev/text/template) individuais para os arquivos que serão criados.

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

### Runs

Arquivo YAML contendo uma sequência de ações que serão executadas ao chamar uma `run`. Inclui logs, comandos de terminal, chamadas HTTP, inputs de texto e múltipla escolha para terminal, além de ações para criar pastas e arquivos baseados nos builders e templates.

**Confira a documentação completa [aqui](#).**

```yaml
# Nome da run que será executada assim que o repositório for obtido através do comando `kuma-cli get`
initial:
  # Descrição da run
  description: "Run inicial para configurar o projeto"

  # Passos que serão executados quando a run for chamada
  steps:
    # Ação de input
    - input:
        label: "Qual é o nome do pacote do seu projeto?"
        out: packageName # Exemplo: github.com/arthurbcp/kuma-hello-world

    # Outra ação de input
    - input:
        label: "Qual mensagem você deseja imprimir?"
        out: msg # Exemplo: Hello, Kuma!

    # Log de mensagem
    - log: "Criando estrutura para {{.data.packageName}}" # Exemplo: Criando estrutura para github.com/arthurbcp/kuma-hello-world

    # Cria a estrutura do projeto usando o builder base.yaml
    - create:
        from: base.yaml

    # Log de sucesso
    - log: "Estrutura base criada com sucesso!"

    # Inicializa o módulo Go
    - cmd: go mod init {{.data.packageName}} # Exemplo: go mod init github.com/arthurbcp/kuma-hello-world

    # Instala dependências
    - cmd: go mod tidy

    # Executa o arquivo main.go
    - cmd: go run main.go # Exemplo: Hello, Kuma!
```

## Comandos de Terminal

### Criar um Boilerplate

O comando `create` é usado para criar um boilerplate baseado nos builders e templates que estão dentro da pasta `.kuma` e em um arquivo JSON ou YAML contendo as variáveis para fazer as substituições nos templates.

```bash
kuma-cli create --variables=swagger.json --project=. --from=base.yaml
```

**Flags:**

- `--variables`, `-v`: Caminho ou URL para o arquivo de variáveis.

- `--project`, `-p`: Caminho para o projeto onde o boilerplate será criado.

- `--from`, `-f`: Caminho para o arquivo YAML com a estrutura e templates.

### Executar uma Run

O comando `exec` é usado para iniciar o processo de uma run.

```bash
kuma-cli exec --run=initial
```

**Flags:**

- `--run`, `-r`: Nome da run que será executada.

### Obter Templates do GitHub

Obtenha templates e runs através de um repositório GitHub.

Você pode obter os templates de qualquer repositório através do comando:

```bash
kuma-cli get --repo=github.com/arthurbcp/typescript-rest-openapi-services
```

Ou utilizar um dos nossos templates oficiais com:

```bash
kuma-cli get --template=typescript-rest-openapi-services
```

**Flags:**

- `--repo`, `-r`: Nome do repositório GitHub.
- `--template`, `-t`: Nome do template oficial.

#### Templates Oficiais

- **[OpenAPI 2.0 TypeScript Services](https://github.com/arthurbcp/typescript-rest-openapi-services):** Crie uma library TypeScript com serviços tipados para todos os endpoints descritos em um arquivo Open API 2.0.

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests para melhorar o Kuma.

1. **Fork o repositório.**
2. **Crie uma branch para a sua feature:**

   ```bash
   git checkout -b minha-nova-feature
   ```

3. **Comite suas mudanças:**

   ```bash
   git commit -m "Adiciona nova feature"
   ```

4. **Envie para o branch:**

   ```bash
   git push origin minha-nova-feature
   ```

5. **Abra um Pull Request.**

## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
