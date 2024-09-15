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
	make tailwind-build
	go install github.com/a-h/templ/cmd/templ@latest
	make sqlc-generate
	go build -ldflags "-X main.Environment=production" -o ./bin ./cmd/main.go

.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: test
test:
	  go test -race -v -timeout 30s ./...
