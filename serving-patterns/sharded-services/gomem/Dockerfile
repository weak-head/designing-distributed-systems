# Build our service
FROM golang:1.15-alpine AS build
WORKDIR /go/src/gomem
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/gomem .
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.2 && \
    wget -qO/go/bin/grpc_health_probe \
    https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /go/bin/grpc_health_probe


# Copy the binaries of our service to the new lightweigh container
FROM scratch
COPY --from=build /go/bin/gomem /bin/gomem
COPY --from=build /go/bin/grpc_health_probe /bin/grpc_health_probe
ENTRYPOINT ["/bin/gomem"]