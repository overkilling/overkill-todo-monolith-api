# overkill-todo-monolith-api

> Beat the drum, beat the drum, beat forever on the endless march - Motörhead

![CI](https://github.com/overkilling/overkill-todo-monolith-api/workflows/CI/badge.svg?branch=master)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/overkilling/overkill-todo-monolith-api)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/overkilling/overkill-todo-monolith-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/overkilling/overkill-todo-monolith-api)](https://goreportcard.com/report/github.com/overkilling/overkill-todo-monolith-api)
[![GitHub license](https://img.shields.io/github/license/overkilling/overkill-todo-monolith-api)](https://github.com/overkilling/overkill-todo-monolith-api/blob/master/LICENSE)


An overkill implementation of a backend API for serving todos, implemented in Go and backed by a Postgres database.
It is a monolith as it handles todos, users, events and anything else the Todo frontend needs.
It is an overkill as it is far too overengineered for a pet project, implementing authentication, pact tests and observability.

## Overview

The current architecture of the Todo application follows the Box-Box-Container model, or BBC for short.
It splits the architecture concerns into three layers: presentation, business logic and persistence.
In the context of the Todo application:

* The presentation layer is a [React Single Page Application (SPA)](https://github.com/overkilling/overkill-todo-spa-frontend)
* The business logic is a [Golang API](https://github.com/overkilling/overkill-todo-monolith-api)
* The persistence is a [Postgres database](https://www.postgresql.org/).

Below is colorful architecture diagram:

![Diagram](/.github/diagram.png?raw=true)

This is a fairly [standard architecture style](https://martinfowler.com/bliki/PresentationDomainDataLayering.html), and it could considered a good starting point for most applications.
If/when an application grows in complexity and features, other patterns could be considered, such as  microservices, event sourcing, CQRS and many more.
Hopefully these other patterns will be explored in other overkill implementations, so they can be evaluated and compared.

In this simple architecture, the SPA and API components where implemented in a slighlty overkill fashion.
This means that, even though the problem space can be considered quite simple, the solution has been overengineered to exercise and highlight certain techniques and methodologies.
E.g. both `BasicAuth` and `JWT` authentication methods will be supported, it will include REST and `GraphQL` APIs, logs and metrics will be gathered.
The intent is to provide some practice to the developer and, perhaps, some education to the readers.

Although this repository only contains the API component, there's an [infrastructure repository](https://github.com/overkilling/overkill-todo-infrastructure) which ties the whole application together.
It allows to start the application locally through `docker-compose`, but it doesn't deploy it yet to any real environment.

## API endpoints

Note that, as there are currently no notion of users, all endpoints fetch and modify a global state.

### `GET /todos`

Lists all todos.

Example response:

```
[
  { "todo": "Some todo" },
  { "todo": "Another todo" },
  { "todo": "Yet another todo" }
]
```

## Getting started

The code has been developed and test with Golang version 1.14.3.

To get started with the API, clone the codebase to your local environment and install all prequisites:

```
git clone git@github.com:overkilling/overkill-todo-monolith-api.git
cd overkill-todo-monolith-api
go mod download
```

### Running the application

There are a few ways of running the application, from a quick and dirty way to running the "full stack".

To quickly start the application from the cloned code, you can run the following (assuming you have a Postgres database already running):

```
make run
```

It is possible to also run it through docker, using `docker-compose`. The benefit is that it will also start up a database, although it will take a bit longer to build the image. The command is:

```
docker-compose --build up
```

There is also the [infrastructure repository](https://github.com/overkilling/overkill-todo-infrastructure), which contains instructions and code to run the whole application locally, including the frontend and the observability stack (`fluentd`, `elasticsearch`, etc).
For more info, checkout the repo.

### Configuration

It is possible to configure the application through a config file or environment variables.
The possible configuration settings are :

| Name | Config | Environment Variable | Default value |
| ---- | ------ | -------------------- | ------------- |
| Database host | `db.host` | `DB_HOST` | `localhost` |
| Database port | `db.port`| `DB_PORT` | `5432` |
| Database name | `db.database` | `DB_DATABASE` | `todo` |
| Database user | `db.username` | `DB_USERNAME` | `postgres` |
| Database password | `db.password` | `DB_PASSWORD` | `postgres` |
| Pretty logging | `log.pretty` | `LOG_PRETTY` | `false` |


#### Example configuration file

```
db:
  host: localhost
  port: 5432
  database: some-todos
  username: some-user
  password: super-secret
log:
  pretty: true
```


### Testing

There three levels of tests: unit, integration and pact tests. You can run them by:

```
make only_unit # only runs unit tests using the -short flag
make test # run both unit and integration tests
make pact
make ci # runs test and pact
```

### Github Actions

TODO
