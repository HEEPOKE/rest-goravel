---
runme:
  id: 01HSTGECX42CFVKYT96R9H4Z13
  version: v3
---

# rest-goravel

## config Environment

```bash
cp .env.example .env
```

- then update environments in .env

## Run Project With Docker

```bash
docker create network Heepoke
```

```bash
docker compose up -d && docker compose logs api --follow
```

## Clear

```bash
go mod tidy
```

## Generate Swagger

```bash
swag init -g cmd/main.go --output=pkg/docs
```

## Local Swagger Doc Api

<http://localhost:6476/apis/docs/index.html>

## Monitor

<http://localhost:6476/apis/monitor>