# Version comes from git tags (vX.Y.Z), see bin/release.sh.
# Untagged commits look like v0.42.0-3-gabc1234, local changes add -dirty.
# Override for builds without .git: make build VERSION=v1.0.0 COMMIT=abc1234
VERSION    ?= $(shell git describe --tags --match 'v[0-9]*' --always --dirty 2>/dev/null || echo dev)
COMMIT     ?= $(shell git rev-parse --short=8 HEAD 2>/dev/null || echo unknown)
BUILD_DATE ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

PKG     := gitlab.com/shaninalex/lumna
LDFLAGS := -X $(PKG).Version=$(VERSION) -X $(PKG).Commit=$(COMMIT) -X $(PKG).BuildDate=$(BUILD_DATE)

.PHONY: build build-backend version release migrate run run_frontend init

# resources/ is embedded into the binary: build the frontend into
# resources/assets first, or the result serves no UI.
build:
	yarn --cwd=./frontend build
	$(MAKE) build-backend

build-backend:
	go build -ldflags "$(LDFLAGS)" -o bin/lumna ./app

version:
	@echo $(VERSION)

# make release BUMP=patch|minor|major (default: minor) or BUMP=1.2.3
release:
	./bin/release.sh $(or $(BUMP),minor)

migrate:
	go run app/main.go --config=./config/config.yaml migrate apply

run:
	go run app/main.go\
 		--config=./config/config.yaml\
 		 serve\
 		 --email-queue\
 		 --refresh-tokens-cleaner

run_frontend:
	yarn --cwd=./frontend start

init: migrate
	go run app/main.go --config=./config/config.yaml workspace create --title=Lumna --owner-email=admin@admin.com
	go run app/main.go --config=./config/config.yaml identity create --email=admin@admin.com --full-name="Alex Shanin" --password=test --active=true
	go run app/main.go --config=./config/config.yaml workspace add_member --workspace-id=1 --identity-id=1
	go run app/main.go --config=./config/config.yaml project create --title=Lumna --workspace-id=1 --owner-id=1
	go run app/main.go --config=./config/config.yaml scope create --name=Lumna --description="This is test board description"  --project-id=1
	go run app/main.go --config=./config/config.yaml stage create --name="Todo" --scope-id=1 --position=0.0
	go run app/main.go --config=./config/config.yaml stage create --name="In Progress" --scope-id=1 --position=1.0
	go run app/main.go --config=./config/config.yaml stage create --name="Done" --scope-id=1 --position=2.0
