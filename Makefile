default:
	go build -o bin/example cmd/example/main.go

docker:
	docker build -f build/Dockerfile -t example:v1 .

compose:
	docker-compose -f deployment/docker-compose.yml up

generate:
	USE_LOCAL=true go generate

health:
	RUN GRPC_HEALTH_PROBE_VERSION=v0.2.0 && \
        wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
        chmod +x /bin/grpc_health_probe