FROM bitnami/golang:1.21.5 as package

WORKDIR /go/src
ENV TZ=Asia/Jakarta
ENV GOPROXY=https://proxy.golang.org

COPY . .
RUN go mod tidy && go build

FROM bitnami/golang:1.21.5 as build

WORKDIR /go/src
ENV TZ=Asia/Jakarta

COPY --from=package /go/src/golang-grpc /go/src/golang-grpc
CMD ["./golang-grpc"]
