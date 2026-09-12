# Lumna

Lumna - is a self hosted, all in one, project management system that can be
hosted in your own vps with minimal setup and configuration.

> **Broken state!**
> 
> I am reviewing the data model and backend architecture. The current approach 
> is a dead end: I have minimal working functionality, yet there are already 
> dozens of services and repositories that can—and do—import each other directly. 
> A more flexible, modular architecture is needed to allow for smooth development. 
> The previous version has been saved to the `archive` branch.

### build

```bash
make build
```

`resources/` (migrations and the built frontend) is always embedded into the
binary, so build the frontend into `resources/assets` before `make build` if you
want the UI served from it.

### development

```bash
go mod tidy
cp ./config/config.test.yaml ./config/config.yaml

# modify path to database

cd frontend
yarn build

go run ./app migrate apply --config=config/config.yaml
go run ./app serve  --config=config/config.yaml

# go to http://localhost:8000/setup and setup your workspace
```
