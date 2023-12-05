# g3offrey's honeypot

This is a web server that log and save (in SQLite) all visits to critical web pages.

## Prerequisites

- Go 1.16 or higher

## Getting Started

1. Clone the repository:

   ```shell
   git clone https://github.com/g3offrey/honeypot.git
   ```

2. Change to the project directory:

   ```shell
   cd honeypot
   ```

3. Build the project:

   ```shell
   go build ./cmd/server
   ```

4. Run the web server:

   ```shell
   ./server
   ```

## Configuration

By default, the web server listens on port 3000. You can change the port by setting a `PORT` environment variable.
The default database path is `honeypot.db`, you can change that by setting the `DB_PATH` environment variable.

## Contributing

Contributions are welcome! If you find any issues or have suggestions, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
