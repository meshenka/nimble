# Nimble 5e Random Character Generator

This project is a full-stack application designed to generate
random fantasy characters.

It features a Go backend that powers a RESTful API for creating and
managing character data, complemented by a React-based frontend for
a seamless user experience. The application is structured with separate
directories for the frontend and backend code, and it includes
comprehensive documentation and testing to ensure reliability
and ease of use.

The backend is built with Go and leverages the stdlib to provide a robust
and efficient API. It allows for the random generation of characters
with various attributes, including ancestry, background, and class. The
API is fully documented using the OpenAPI (Swagger) specification,
which can be found in the `docs/swagger.yaml` file. The project also
includes a command-line interface (CLI) for generating characters
directly from the terminal.

```bash
make cli
```

The frontend is a single-page application (SPA) built with React and
TypeScript, using Webpack for bundling and development. It communicates
with the backend API to fetch and display character information in a
simple interface.

The project is set up for easy development and deployment, with a
`Makefile` that provides commands for building, testing, and running
the application. It is also configured for deployment using Docker
and Fly.io, as indicated by the `Dockerfile` and `fly.toml` files.

```bash
# see all actions
make help

# start local http server
make api
```
