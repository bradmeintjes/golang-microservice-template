# webservice-template

A basic template for creating a microservice in golang

## Requirements

Golang: 1.16+
Docker (optional)
Postgres (optional)

## Development

If you have docker installed then you can run scripts/setup_db to run a local postgres container, otherwise you will need to update the config to point to your instance of postgres


### Debugging

Using vscode you can just tap f5 when you have /cmd/main.go open and it will kick off the debugger.

### Running

```./scripts/dev```

### Docker

```./scripts/docker```