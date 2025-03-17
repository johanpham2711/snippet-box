.PHONY: db-start db-stop run run-dev install-air get-assets tls
db-start:
	docker compose up -d
db-stop:
	docker compose down
install-air:
	go install github.com/air-verse/air@latest
	export PATH=$(go env GOPATH)/bin:$PATH
	source ~/.zshrc
	export PATH="$HOME/go/bin:$PATH"
	source ~/.zshrc
get-assets:
	@mkdir -p ui/static
	curl https://www.alexedwards.net/static/sb-v2.tar.gz | tar -xvz -C ./ui/static/
tls:
	@mkdir -p tls
	@cd tls && go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
run:
	go run ./cmd/web
run-dev:
	air
test:
	go test -v ./...
all: 
	make db-start
	make install-air
	make tls
	make run-dev