.PHONY: compile
compile: ## Compile the proto file.
	protoc -I src/pkg/proto/emiten/ src/pkg/proto/emiten/emiten.proto --go_out=plugins=grpc:src/pkg/proto/emiten/

.PHONY: compile-auth
compile-auth: ## Compile the proto file auth service.
	protoc -I src/pkg/proto/auth/ src/pkg/proto/auth/auth.proto --go_out=plugins=grpc:src/pkg/proto/auth/
 
.PHONY: server
server: ## Build and run server.
	go build -race -ldflags "-s -w" -o bin/server server/main.go
	bin/server
 
.PHONY: client
client: ## Build and run client.
	go build -race -ldflags "-s -w" -o bin/client src/main.go
	bin/client