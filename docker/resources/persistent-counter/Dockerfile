# syntax=docker/dockerfile:1

FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY persistent-counter.go ./

RUN CGO_ENABLED=0 go build -o /counter

FROM scratch

COPY --from=build /counter /go/bin/counter

EXPOSE 8080

ENTRYPOINT [ "/go/bin/counter" ]
