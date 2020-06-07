# overkill-todo-monolith-api

> Beat the drum, beat the drum, beat forever on the endless march - Motörhead

![CI](https://github.com/overkilling/overkill-todo-monolith-api/workflows/CI/badge.svg?branch=master)

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

To quickly start the application from the cloned code, you can run the following:

```
make run
```

To run it through Docker, you can use the following `docker-compose` command (the `--build` will generate a new docker image):

```
docker-compose --build up
```

There is also the [infrastructure repository](https://github.com/overkilling/overkill-todo-infrastructure), which contains instructions and code to run the whole application locally, including the frontend and the observability stack (`fluentd`, `elasticsearch`, etc).


### Testing

There three levels of tests: unit, integration and pact tests. You can run them by:

```
make unit
make integration
make pact
make all # runs all of the above
```

### Github Actions