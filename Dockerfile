FROM golang:1.21 as build

RUN apt-get update && \
    apt-get install -y protobuf-compiler && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
    go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest && \
    find /go -exec chmod go+rw {} \;

ADD . /app
WORKDIR /app
RUN ./build.sh

FROM gcr.io/distroless/base-debian12
COPY --from=build /app/bin/* /app/

WORKDIR /app

CMD ["/app/helium-route-updater"]
