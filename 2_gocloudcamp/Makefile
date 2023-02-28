testcli:
	@go run cmd/testcli/main.go

server:
	@go run cmd/server/main.go

up:
	cd deploy/postgres && \
	docker compose up -d && \
	sleep 10 && \
	docker-compose exec postgres psql -U user -c 'CREATE DATABASE playlist'

down:
	cd deploy/postgres && docker compose down

test:
	go test ./...

gen: clean
	protoc --proto_path=api --go_out=pkg --go-grpc_out=pkg api/*.proto

clean:
	-@rm pkg/plservice/*.go