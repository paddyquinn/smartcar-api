# smartcar-api

## Set Up
Install go via instructions: https://golang.org/doc/install

```
go get -u github.com/kardianos/govendor
govendor sync
```

## Run API Web Server
`go run main.go`

## Endpoints

```
curl localhost:8080/vehicles/:id -X GET
curl localhost:8080/vehicles/:id/doors -X GET
curl localhost:8080/vehicles/:id/fuel -X GET
curl localhost:8080/vehicles/:id/battery -X GET
curl localhost:8080/vehicles/:id/engine -X POST -H 'Content-Type: application/json' -d '{"action":"START|STOP"}'
```

## Test Code
`go test ./...`