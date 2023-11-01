# Origin Backend Take-Home Assignment

[![Run tests on PR -> Ubuntu](https://github.com/victormagalhaess/origin-backend-take-home-assignment/actions/workflows/main.yml/badge.svg)](https://github.com/victormagalhaess/origin-backend-take-home-assignment/actions/workflows/main.yml)
[![Deploy](https://github.com/victormagalhaess/origin-backend-take-home-assignment/actions/workflows/deploy.yml/badge.svg)](https://github.com/victormagalhaess/origin-backend-take-home-assignment/actions/workflows/deploy.yml)

This is an implementation of [Origin's Backend Take Home Assignment](https://github.com/OriginFinancial/origin-backend-take-home-assignment).
It aims to determine the user’s insurance needs by asking personal and risk-related questions and gathering information about the user’s vehicle and house.


## How to run
There are two ways for running the application: 
- Using local installation of Go
- Using docker

The application was built using [Go 1.16](https://go.dev/dl/), so in order to run it using a local installation of Go, you must install the proper Go version.

To run using docker, you must have Docker installed. 

Both ways were intended to be run on Ubuntu 20.04 or newer, however it may run without too much problem on any other system that contains the requirements. Keep in mind that running the application using ***Docker on Ubuntu*** is probably the best (and recommended) way of running it.

To run it locally, open the project folder in a terminal and run:

```sh
make run
```

To run it using Docker, run:
```sh
make docker/run
```

> Both commands take care of the installation of dependencies and build of binaries, however you can just build the binary (or the image, in Docker's case) through the commands:

```sh
make build
```

> Or for docker:

```sh
make docker/build
```

After running, the application will be live on the localhost at the port number set in Makefile's variable PORT. You can change it as you please, but the default port is 8080.

## How to test
The application counts with a series of automated tests. They can also be run using a local installation of Go, or using the Docker. 

To run the tests locally, open the project folder in a terminal and run:

```sh
make test
```

To run it using Docker, run:
```sh
make docker/test
```

## Documentation
The project counts with automated documentation for the api, using [Open Api's Swagger](https://swagger.io/specification/). The documentation is already built, but can be regenerated (using the local Go installation) using the command:

```sh
make docs
```

It will update the swagger documentation. After running the aplication on a given port (default 8080), the swagger will be live on the [localhost:8080/swagger/index.html](localhost:8080/swagger/index.html). It counts with the routes specification and models definitions as well.

# Using the api
The api routes, input and output examples are better show on the swagger page, however a brief explanation of the routes follows:

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
Examples of cUrl requests for the endpoints:

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

# Discussion

## Relevant comments about the project
This project raised the opportunity to use a lot of knowledge that isn't code, such as:
- [Structurized task management](https://trello.com/b/bGp3yzhe/take-home-assignment)
- [Automated tests and deploy](https://github.com/victormagalhaess/origin-backend-take-home-assignment/actions) (the public instance of the api is accessible on [this endpoint](https://origin-take-home-assignment.herokuapp.com/))
- [Standartized PR's](https://github.com/victormagalhaess/origin-backend-take-home-assignment/pulls?q=is%3Apr+is%3Aclosed)

It was a good exercise to use this knowledge on a take-home assignment.

## Technical decisions
The main technical decisions all starts with the use of Go. It is a "*good-to-write, fast-to-run*" language, that allowed a few interesting decisions along the way. It has a really powerful way to serialize and desserialize json to Go structs, using the [struct tags](https://github.com/victormagalhaess/origin-backend-take-home-assignment/blob/515ff2c6144b0371ea3c98c277576c72b6f8eebb/pkg/model/user.go#L13).

### Extensible risk calculator
One of the main concerns were to make the risk engine extensible to validations, since the process of validating an insurance can grow and change with time.
The solution to this problem was to build three packages that allowed the application to run a stream of steps, allowing the developer to change the stream adding or removing steps.

<p align="center">
  <img src="https://i.ibb.co/Kq28qmb/asas-drawio.png" width="300" alt="Risk engine flux">
</p>

The packages responsible for the flow showed in the chart above are:
- [steps](https://github.com/victormagalhaess/origin-backend-take-home-assignment/tree/main/pkg/pipeline/steps):
    - The steps package is responsible to hold all the functions that implements the actual risk validations. Each one can process the user data and update a transitory risk profile, which after being run across all the required steps turns into the final risk profile.

- [pipeline](https://github.com/victormagalhaess/origin-backend-take-home-assignment/tree/main/pkg/pipeline/):
    - The pipeline package implements a struct that holds steps. It is a holder that the **engine** uses to save the steps that will be executed on the user input.

- [engine](https://github.com/victormagalhaess/origin-backend-take-home-assignment/tree/main/pkg/engine):
    - The engine package is the actual "engine" for the risk calculation. It initializes the pipeline with all the steps that will be used to build the risk profile. Then, on receiving an user information object it build a risk profile. Finally it calls each step to build a risk profile based on the steps results. 

Using these packages it is really simple to add or remove steps to the risk algorithm. The engine fills the pipeline with the required steps, so in order to change it, you just need to write any new steps using the pattern of the other steps and add or remove them when the [engine is fueled](https://github.com/victormagalhaess/origin-backend-take-home-assignment/blob/515ff2c6144b0371ea3c98c277576c72b6f8eebb/pkg/engine/engine.go#L11).

### Project structure
Other interesting approach used was the project structure.
In order to ensure the extensibility of the project to new routes, the api flow was divided in services and controllers.

The controllers were responsible to deal with the REST interface and operations, where all the business logic and communication with other system packages must be in the services.

The [api](https://github.com/victormagalhaess/origin-backend-take-home-assignment/tree/main/pkg/api/) package holds the routes initialization, and can be easily extend to hold even more routes and middlewares just with a new controller function.

In order to add a new route, a new controller is needed. It may also require a service and new models, so each of these features are divided on it's own package.

## Non-ideal situations
Given the characteristics of this take-home assignment, a few decisions are taken to save on time, complexity and money.
An example is the [current deploy](https://origin-take-home-assignment.herokuapp.com/), running on the free cloud platform [Heroku](https://heroku.com), suffers a lot from start latency and can not be easily scaled. These are characteristics of the cost trade-off of using a free platform. However all these non-ideal situations raises the opportunity to improve on the future. A few ideas about this topic were raised and listed here:

- Use a better cloud provider
- Use a custom domain to host the application
- Since the user data could be sensitive, implement authentication and encrypt the requests on the sensitive routes
- Implement tests focused on load
- Implement system tests
- Since the risk engine runs through a lot of steps, it would be good to add some [Fuzzing tests](https://go.dev/doc/tutorial/fuzz) to generate more inputs and help the process of finding edge cases
- Add better logging. The default Go log package isn't bad, but it can produce only really simple log messages. It would be better if the logs were more rich and contentful


