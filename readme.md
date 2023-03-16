# UOW Example

This repository contains an example implementation of the Unit of Work pattern in Go using GORM and Gin. Please keep in mind that this is just a PoC and it's not following the best practices of architecture or Go, as it was made quickly for demo purposes.

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

## Usage

The application exposes two endpoints:

GET /users - Returns all the users in the database.
POST /users - Creates a new user.
