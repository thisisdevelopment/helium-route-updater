FROM golang:1.18

RUN apt-get update && \
    apt-get install -y protobuf-compiler && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
    go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest && \
    go install github.com/spf13/cobra-cli@latest && \
    go install github.com/thisisdevelopment/CompileDaemon@forked && \
    find /go -exec chmod go+rw {} \;

ENV GOCACHE=/app/.cache

WORKDIR /app

ENTRYPOINT CompileDaemon -graceful-kill -exclude-dir=.cache -exclude-dir=data -exclude-dir=bin --build="go build -o /app/bin/ $BUILD" --command="$COMMAND"
