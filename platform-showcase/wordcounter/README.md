# wordcounter

> wordcounter is a platform based service

## Running

### Using `platformctl`

For debugging:

```shell
platformctl service debug
```

### Directly

```shell
docker build \
  --file .platform/docker/Dockerfile \
  --tag wordcounter:latest \
  .

docker compose \
  --file .platform/docker-compose/docker-compose.yaml \
  --env-file .platform/env/.env \
  up -d
```

## Stopping

### Using `platformctl`

```shell
platformctl service stop
```

### Directly

```shell
docker compose \
  --file .platform/docker-compose/docker-compose.yaml \
  --env-file .platform/env/.env \
  down --remove-orphans
```
