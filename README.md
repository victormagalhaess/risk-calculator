# Risk Calculator

Aluno: Victor Hugo Faria Dias Magalhães

[![Run tests on PR -> Ubuntu](https://github.com/victormagalhaess/origin-backend-take-home-assignment/actions/workflows/main.yml/badge.svg)](https://github.com/victormagalhaess/origin-backend-take-home-assignment/actions/workflows/main.yml)

Esta é uma implementação de uma Calculadora de Risco para seguros. Seu objetivo é determinar as necessidades de seguro do usuário fazendo perguntas pessoais e relacionadas ao risco e coletando informações sobre o veículo e a residência do usuário.

## Como executar

Existem duas maneiras de executar a aplicação:

Usando a instalação local do Go
Usando Docker
A aplicação foi construída usando o Go 1.16, portanto, para executá-la usando a instalação local do Go, você deve instalar a versão adequada do Go.

Para executar usando o Docker, você deve ter o Docker instalado.

Ambas as maneiras foram projetadas para serem executadas no Ubuntu 20.04 ou mais recente, no entanto, é possível que ela funcione em outros sistemas que atendam aos requisitos. Tenha em mente que executar a aplicação usando o Docker no Ubuntu é provavelmente a melhor (e recomendada) maneira de executá-la.

Para executá-la localmente, abra a pasta do projeto em um terminal e execute:

```sh
make run
```

Para rodar usando docker:

```sh
make docker/run
```

> Ambos os comandos cuidam da instalação das dependências e da compilação dos binários, no entanto, você pode simplesmente compilar o binário (ou a imagem, no caso do Docker) com os comandos:

```sh
make build
```

> Ou para o docker:

```sh
make docker/build
```

Depois de executar, a aplicação estará ativa em localhost na porta configurada na variável PORT do Makefile. Você pode alterá-la como desejar, mas a porta padrão é 8080.

## Como testar

A aplicação conta com uma série de testes automatizados. Eles também podem ser executados usando a instalação local do Go ou o Docker.

Para executar os testes localmente, abra a pasta do projeto em um terminal e execute:

```sh
make test
```

Para rodar usando docker:

```sh
make docker/test
```

## Documentação

O projeto conta com documentação automatizada para a API, usando o [Swagger da Open Api](https://swagger.io/specification/). A documentação já está construída, mas pode ser regenerada (usando a instalação local do Go) com o comando:

```sh
make docs
```

Isso atualizará a documentação do Swagger. Após executar a aplicação em uma determinada porta (padrão 8080), o Swagger estará disponível em [localhost:8080/swagger/index.html](localhost:8080/swagger/index.html). Ele inclui a especificação das rotas e definições de modelos.

# Usando a api

As rotas da API, exemplos de entrada e saída são melhor apresentados na página do Swagger, no entanto, uma breve explicação das rotas segue abaixo:

[GET] **/api/v1/healthcheck**

- Receives: none
- Produces: plain/text

[POST] **/api/v1/risk**

- Receives: body -> application/json

```js
{
 "age": integer,
 "dependents": integer,
 "house": {
   "ownership_status": "mortgaged" || "owned"
 },
 "income": integer,
 "marital_status": "married" || "single",
 "risk_questions": [3]boolean,
 "vehicle": {
   "year": integer
 }
}
```

- Produces: application/json

```js
{
  "auto": "economic",
  "disability": "regular",
  "home": "ineligible",
  "life": "responsible"
}
```

Exemplos de solicitações cUrl para os endpoints:

- /api/v1/healthcheck

```bash
curl --location --request GET 'localhost:8080/api/v1/healthcheck'
```

- /api/v1/risk

```bash
curl --location --request POST 'localhost:8080/api/v1/risk' \
--data-raw '{
  "age": 35,
  "dependents": 2,
  "house": {"ownership_status": "owned"},
  "income": 0,
  "marital_status": "married",
  "risk_questions": [0, 1, 0],
  "vehicle": {"year": 2018}
}'
```

# Discussão

## Comentários relevantes sobre o projeto

Este projeto proporcionou a oportunidade de usar muito conhecimento que não é apenas código, como:

- [Testes automatizados](https://github.com/victormagalhaess/origin-backend-take-home-assignment/actions)
- [PR's padronizados](https://github.com/victormagalhaess/origin-backend-take-home-assignment/pulls?q=is%3Apr+is%3Aclosed)

## Decisões técnicas

As principais decisões técnicas começam com o uso do Go. É uma linguagem "_boa para escrever, rápida para executar_", que permitiu algumas decisões interessantes ao longo do caminho. Ela possui uma maneira muito poderosa de serializar e desserializar JSON em structs do Go, usando as [struct tags](https://github.com/victormagalhaess/origin-backend-take-home-assignment/blob/515ff2c6144b0371ea3c98c277576c72b6f8eebb/pkg/model/user.go#L13).

### Calculadora extensível de risco

Uma das principais preocupações foi tornar o mecanismo de risco extensível para validações, uma vez que o processo de validação de um seguro pode crescer e mudar com o tempo. A solução para esse problema foi construir três pacotes que permitiram que a aplicação executasse uma sequência de etapas, permitindo que o desenvolvedor alterasse a sequência adicionando ou removendo etapas.

<p align="center">
  <img src="https://i.ibb.co/Kq28qmb/asas-drawio.png" width="300" alt="Risk engine flux">
</p>

Os pacotes responsáveis pelo fluxo mostrado no gráfico acima são::

- [steps](https://github.com/victormagalhaess/origin-backend-take-home-assignment/tree/main/pkg/pipeline/steps):

  - O pacote steps é responsável por conter todas as funções que implementam as validações reais de risco. Cada uma delas pode processar os dados do usuário e atualizar um perfil de risco transitório, que, após ser executado em todas as etapas necessárias, se torna o perfil de risco final.

- [pipeline](https://github.com/victormagalhaess/origin-backend-take-home-assignment/tree/main/pkg/pipeline/):

  - O pacote pipeline implementa uma estrutura que mantém as etapas. É um recipiente que a **engine** usa para salvar as etapas que serão executadas com base nos dados do usuário.

- [engine](https://github.com/victormagalhaess/origin-backend-take-home-assignment/tree/main/pkg/engine):
  - O pacote engine é o "motor" real para o cálculo de risco. Ele inicializa o pipeline com todas as etapas que serão usadas para construir o perfil de risco. Em seguida, ao receber um objeto de informações do usuário, ele cria um perfil de risco. Finalmente, ele chama cada etapa para construir um perfil de risco com base nos resultados das etapas.

Usando esses pacotes, é muito simples adicionar ou remover etapas para o algoritmo de risco. O motor preenche o pipeline com as etapas necessárias, portanto, para alterá-lo, basta escrever novas etapas usando o padrão das outras etapas e adicioná-las ou removê-las quando a [engine é alimentada](https://github.com/victormagalhaess/origin-backend-take-home-assignment/blob/515ff2c6144b0371ea3c98c277576c72b6f8eebb/pkg/engine/engine.go#L11).

### Estrutura do projeto

Outra abordagem interessante usada foi a estrutura do projeto. Para garantir a extensibilidade do projeto para novas rotas, o fluxo da API foi dividido em serviços e controladores.

Os controllers eram responsáveis por lidar com a interface REST e operações, onde toda a lógica de negócios e comunicação com outros pacotes do sistema deve estar nos serviços.

O pacote [api](https://github.com/victormagalhaess/origin-backend-take-home-assignment/tree/main/pkg/api/) mantém a inicialização das rotas e pode ser facilmente estendido para incluir ainda mais rotas e middlewares apenas com uma nova função de controlador.

Para adicionar uma nova rota, é necessário um novo controller. Isso também pode exigir um serviço e novos modelos, portanto, cada um desses recursos é dividido em seu próprio pacote.
