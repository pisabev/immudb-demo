FROM golang:1.20.5

WORKDIR /immudb-demo

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./app ./cmd/main.go

EXPOSE 8080

# Run
CMD ["/immudb-demo/app"]