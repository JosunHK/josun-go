.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./static/input.css -o ./static/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	./tailwindcss -i ./static/input.css -o ./static/style.min.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch

.PHONY: dev
dev:
	go build -o ./tmp/$(APP_NAME) ./src/$(APP_NAME)/main.go && air

.PHONY: build
build:
	make tailwind-build
	make templ-generate
	go build -ldflags "-X main.Environment=production" -o ./bin/$(APP_NAME) ./src/$(APP_NAME)/main.go

.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: test
test:
	  go test -race -v -timeout 30s ./...