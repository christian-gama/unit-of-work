# Unit of Work Example

This repository contains an example implementation of the Unit of Work pattern in Go using GORM and Gin. Please keep in mind that this is just a PoC and it's not following the best practices of architecture or Go, as it was made quickly for demo purposes.
Some services or repositories may look too dumb, but as I said, this is just a PoC and the purpose is to show how the UoW pattern works.

Access the [API documentation](https://documenter.getpostman.com/view/10505101/TzXzDj8hhttps://documenter.getpostman.com/view/16405037/2s93Jxq1UB) to see how to use the endpoints.

## Getting Started

### Prerequisites

To run this example, you'll need the following:

- Docker

## Installation

1. Clone the repository:

```bash
git clone https://github.com/christian-gama/uow.git
```

2. Run the application:

```bash
make run
```

If you don't have `make` installed, you can run the following commands:

```bash
docker compose up api
```

## Usage

The application exposes the following endpoints:

**`GET /users`** Returns all the users in the database.

**`GET /users/:id`** Returns a specific user.

**`POST /users/transfer`** Transfers money from one user to another.

**`POST /users`** Creates a new user or updates an existing one (if the ID is provided).

**`DELETE /users/:id`** Deletes a user.
