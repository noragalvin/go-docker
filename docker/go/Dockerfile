FROM golang:1.11.0-alpine

RUN apk add bash ca-certificates git gcc g++ libc-dev

ENV CGO_ENABLED=0

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -build="go build" -include="*.*" -command="./go-docker-example"

# CMD ["CompileDaemon", "-build='go build'", "-command='./go-docker-example'"]

# CMD ["go", "run", "main.go"]
