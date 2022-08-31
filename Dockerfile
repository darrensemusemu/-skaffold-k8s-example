# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18-buster AS build

WORKDIR /usr/src/app

ARG BINARY_NAME

# Download named modules
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app  ./cmd/${BINARY_NAME}/main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /usr/local/bin/app  /usr/local/bin/app

USER nonroot:nonroot

ENTRYPOINT [ "app" ] 
