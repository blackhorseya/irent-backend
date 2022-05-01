FROM golang:alpine AS builder

ARG APP_NAME

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY pb ./pb
COPY api ./api
RUN go build -o app ./cmd/${APP_NAME}

FROM alpine:3 AS final

LABEL maintainer.name="blackhorseya"
LABEL maintainer.email="blackhorseya@gmail.com"

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /src/app ./

ENTRYPOINT ./app
