# Go-RSS

Go-RSS is a Golang learning project inspired by the [Boot.dev](https://www.youtube.com/@BootDotDev) YouTube channel. This project implements an RSS server that hosts basic server routes and fetches RSS feeds regularly using Goroutines.

## Features

- Hosts a basic RSS server with RESTful routes.
- Regularly fetches and updates RSS feeds concurrently using Goroutines.
- Utilizes SQL code generation and schema migrations for robust database management.

## Technologies Used

- **[sqlc](https://github.com/kyleconroy/sqlc)**: For generating type-safe SQL queries.
- **[Goose](https://github.com/pressly/goose)**: For managing database schema migrations.
- **[Chi](https://github.com/go-chi/chi)**: For lightweight and fast HTTP routing.

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```
PORT=8080
PG_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
```

- `PORT`: The port on which the server will run.
- `PG_URL`: PostgreSQL connection string.

## Build & Run Instructions

To build and run the Go-RSS project, execute the following command:

```bash
go build -o go-rss cmd/server/main.go && ./go-rss
```

## License

This project is for learning purposes and is open-source under the MIT License.

---

Happy coding! 🚀

