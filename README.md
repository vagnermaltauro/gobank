# GoBank

GoBank is a Go project that provides a simple banking API.

## Prerequisites

- Go version 1.x
- Docker and Docker Compose

## Installation

Clone the repository:

```bash
git clone https://github.com/<your-username>/gobank.git
```

Install the dependencies:

```bash
make deps
```

## Running

To run the project, use the command:

```bash
make run
```

## Testing

To run the tests, use the command:

```bash
make test
```

## API Endpoints

### POST `/login`

- **Description**: Authenticates a user and returns a JWT token.
- **Parameters**: JSON body with `Number` and `Password`.
- **Response**: JSON with `Token` and `Number`.

### GET `/account`

- **Description**: Retrieves all accounts.
- **Parameters**: None.
- **Response**: JSON array of accounts.

### POST `/account`

- **Description**: Creates a new account.
- **Parameters**: JSON body with `FirstName`, `LastName`, and `Password`.
- **Response**: JSON of the created account.

### GET `/account/{id}`

- **Description**: Retrieves an account by its ID.
- **Parameters**: `id` in the URL path.
- **Response**: JSON of the requested account.

### DELETE `/account/{id}`

- **Description**: Deletes an account by its ID.
- **Parameters**: `id` in the URL path.
- **Response**: JSON with `deleted` field containing the ID of the deleted account.

### POST `/transfer`

- **Description**: Initiates a transfer between accounts.
- **Parameters**: JSON body with details of the transfer.
- **Response**: JSON of the transfer request.

## Contributing

Contributions are very welcome! If you'd like to contribute, these are the steps:

1. Fork the repository.
2. Create a new branch in your forked repository.
3. Make your changes in the new branch.
4. Submit a pull request from the new branch in your forked repository to the main branch in the original repository.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

