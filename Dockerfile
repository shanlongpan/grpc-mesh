FROM golang:1.17.6 as builder
WORKDIR /root/grpc-mesh
COPY ./ .
RUN CGO_ENABLED=0 GOOS=linux go build -o grpc-mesh main.go

FROM alpine:3.15
WORKDIR /bin/
COPY --from=builder /root/grpc-mesh/grpc-mesh .
COPY --from=builder /root/grpc-mesh/xds_bootstrap.json .
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Shanghai
ENV GRPC_XDS_BOOTSTRAP=./xds_bootstrap.json
ENTRYPOINT [ "/bin/grpc-mesh" ]

#FROM scratch
#WORKDIR /bin/
#COPY grpc-mesh .
#ENTRYPOINT [ "/bin/grpc-mesh" ]
