# httpserver

A command line utility to create simple http servers.

## Usage
To start a http server on port 1331
`go run main.go -p 1331`

When the server receives a request, request details will be logged.
The server will also respond with the date and time the request was received.