SERVER_OUT := "bin/server"
CLIENT_OUT := "bin/client"
PKG := "rpc_demo1"
SERVER_PKG_BUILD := "${PKG}/server"
CLIENT_PKG_BUILD := "${PKG}/client"

.PHONY: server client

all: server client

dep: ## Get the dependencies
	@go get -v -d ./...

server: dep
	@go build -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

client: dep
	@go build -v -o $(CLIENT_OUT) $(CLIENT_PKG_BUILD)

clean:
	@rm $(SERVER_OUT) $(CLIENT_OUT)