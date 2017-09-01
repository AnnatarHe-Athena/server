deps:
	@go get github.com/revel/cmd/revel
	@go get github.com/graphql-go/graphql
	@go get -u github.com/go-redis/redis
	@go get github.com/lib/pq
	@go get -u golang.org/x/crypto/scrypt

echo:
	@echo $(PATH)
	@echo $(GOPATH)
	@echo $(revel)

build:
	@cd /tmp
	@GOOS=linux GOARCH=amd64 revel package github.com/douban-girls/server prod
	@echo "please to this project directory to scp this file to your server"