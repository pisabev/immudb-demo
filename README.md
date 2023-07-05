# immudb-demo

Write logs to immudb using REST API

### Install
Running the applicatoin using `docker-compose`
```bash
git clone git@github.com:pisabev/immudb-demo.git
cd immudb-demo
docker-compose up
```

### Exposed addresses and ports:

##### API: http://localhost:8080/api/v1 - *uses basic HTTP Authentication - [RFC 7235](https://datatracker.ietf.org/doc/html/rfc7235) (Username: immudb, Password: immudb)*

##### Immudb client: http://localhost:8081
##### Immudb-demo client: http://localhost:8080/

## Client service

### curl.sh
Provides curl examples of the API

### client/main.go
Simple client application written in Go.

```
$ go run client/main.go -h
Usage of /tmp/go-build2740329148/b001/exe/main:
  -address string
        API Address/port (default "immudb:immudb@localhost:8080")
  -count
        Show logs count
  -insert
        Inserts a random log line
  -query-all
        Fetch all logs
  -query-last int
        Fetch last X logs

```