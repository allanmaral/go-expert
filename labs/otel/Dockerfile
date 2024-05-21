# Builder
FROM golang:1.22.2 as build

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o microservice ./cmd/microservice


# Runner
FROM scratch

WORKDIR /app

COPY --from=build /app/microservice .

ENTRYPOINT ["./microservice"]
