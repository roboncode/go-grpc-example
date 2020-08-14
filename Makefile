default:
	go build -o bin/example cmd/example/main.go

docker:
	docker build -f build/Dockerfile -t example:v1 .

compose:
	docker-compose -f deployment/docker-compose.yml up

generate:
	USE_LOCAL=true go generate

checkhealth:
	$HOME/go/bin/grpc-health-probe -addr=localhost:8080
