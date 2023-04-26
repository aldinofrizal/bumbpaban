FROM golang:1.19-alpine AS build

WORKDIR /usr/app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

EXPOSE 8000

RUN go build -o /app


FROM alpine:3.16

WORKDIR /

COPY --from=build /app /app

EXPOSE 8000

ENTRYPOINT [ "/app" ]

