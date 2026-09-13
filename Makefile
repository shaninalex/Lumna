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
	go run app/main.go --config=./config/config.yaml identity create --email=admin@admin.com --full-name="Alex Shanin" --password=test --active=true
	go run app/main.go --config=./config/config.yaml workspace add_member --workspace-id=1 --identity-id=1
	go run app/main.go --config=./config/config.yaml project create --title=Lumna --workspace-id=1 --owner-id=1
	go run app/main.go --config=./config/config.yaml scope create --name=Lumna --description="This is test board description"  --project-id=1
	go run app/main.go --config=./config/config.yaml stage create --name="Todo" --scope-id=1 --position=0.0
	go run app/main.go --config=./config/config.yaml stage create --name="In Progress" --scope-id=1 --position=1.0
	go run app/main.go --config=./config/config.yaml stage create --name="Done" --scope-id=1 --position=2.0
