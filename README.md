# Kuma

Kuma é um poderoso framework projetado para gerar estruturas de projetos para qualquer linguagem de programação com base em templates Go. Ele agiliza o processo de configuração de novos projetos, automatizando a criação de diretórios, arquivos e código base, garantindo consistência e economizando tempo valioso de desenvolvimento.

### Funcionalidades:

- Estruture os diretórios e arquivos do seu projeto de maneira customizada através de templates Go
- Integração com Github para poder baixar templates pré definidos pela comunidade ou para uso pessoal através de repositórios privados
- Possibilidade de criar workflows de comandos cli personalizados através de um arquivo YAML usando as runs
- Utilizar variáveis de forma dinâmica para serem aplicadas aos templetes. As variáveis podem ser extraídas de um arquivo YAML ou JSON que pode ser local ou baixado de uma URL pública. As variáveis também podem ser obtidas através de informações passadas pelo usuário durante a execução de uma run

### Crie seus próprios boilerplates utilizando Kuma

Para o funcionamento do Kuma, todos arquivos relacionados ao framework devem ficar dentro da pasta `.kuma-files`

O framework funciona com um parser de templates Go que são divididos em 3 tipos de arquivos

- **Builders**

- **Templates**

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
