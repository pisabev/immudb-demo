FROM node:18-alpine AS client

ADD web ./web

RUN npm install --global pnpm

WORKDIR ./web
RUN pnpm run build

FROM golang:1.20.5

WORKDIR /immudb-demo

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY --from=client web/dist ./public

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./app ./cmd/main.go

EXPOSE 8080

# Run
CMD ["/immudb-demo/app"]