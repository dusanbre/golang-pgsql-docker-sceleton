FROM golang:alpine

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

RUN mkdir /app

COPY .. /app

WORKDIR /app

RUN go get github.com/githubnemo/CompileDaemon && go install github.com/githubnemo/CompileDaemon

COPY . .

# Flag -C will change the directory to the one specified, that is why in output flag we need to go back to the root directory
ENTRYPOINT CompileDaemon -build="go build -o bin/main ./cmd/goPgsqlDocker/" -command="./bin/main"