# wordcounter

> wordcounter is a docker-compose based service

## Running

### Using `platformctl`

For debugging:

```shell
platformctl service debug
```

### Directly

```shell
docker-compose --env-file .env up -d
```

## Stopping

### Using `platformctl`

```shell
platformctl service stop
```

### Directly

```shell
docker-compose down --remove-orphans
```
