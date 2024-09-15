.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./web/static/input.css -o ./web/static/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	./tailwindcss -i ./web/static/input.css -o ./web/static/style.min.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch

.PHONY: sqlc-generate
sqlc-watch:
	sqlc generate 

.PHONY: dev
dev:
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air

.PHONY: build
build:
	go env -w GOPATH=$${HOME}/go
	export PATH=$${PATH}:`go env GOPATH`/bin
	make tailwind-build
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	`go env GOPATH`/bin/templ generate
	`go env GOPATH`/bin/sqlc generate
	go build -ldflags "-X main.Environment=production" -o ./bin/ ./cmd/main.go

.PHONY: start
start:
	./bin/main

.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: test
test:
	  go test -race -v -timeout 30s ./...
