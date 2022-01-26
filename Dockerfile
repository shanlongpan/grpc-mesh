#FROM golang:1.17.6 as builder
#WORKDIR /root/grpc-mesh
#COPY ./ .
#RUN CGO_ENABLED=0 GOOS=linux go build -o grpc-mesh main.go
#
#FROM scratch
#WORKDIR /bin/
#COPY --from=builder /root/grpc-mesh/grpc-mesh .
#ENTRYPOINT [ "/bin/grpc-mesh" ]

FROM scratch
WORKDIR /bin/
COPY grpc-mesh .
ENTRYPOINT [ "/bin/grpc-mesh" ]