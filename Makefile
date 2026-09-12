# resources/ is embedded into the binary: build the frontend into
# resources/assets first, or the result serves no UI.
build:
	go build -o bin/lumna ./app

migrate:
	go run app/main.go --config=./config/config.yaml migrate apply

run:
	go run app/main.go --config=./config/config.yaml serve

run_frontend:
	yarn --cwd=./frontend start

init: migrate
	go run app/main.go --config=./config/config.yaml workspace create --title=Lumna --owner-email=admin@admin.com
	go run app/main.go --config=./config/config.yaml identities create --email=admin@admin.com --full-name="Alex Shanin" --password="test" --active=true --workspace-id=1
#	go run app/main.go --config=./config/config.yaml projects create --title="Lumna" --workspace_id=1 --owner_id=1
#	go run app/main.go --config=./config/config.yaml boards create --title="Lumna" --project_id=1
#	go run app/main.go --config=./config/config.yaml column create --title="Todo" --board_id=1
#	go run app/main.go --config=./config/config.yaml column create --title="In Progress" --board_id=1
#	go run app/main.go --config=./config/config.yaml column create --title="Done" --board_id=1
