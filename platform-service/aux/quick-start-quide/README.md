# platform-service / aux / quick-start-guide

## Build

```shell
platformctl task template
docker build --tag platform-service/aux/quick-start-guide .
```

## Run

```shell
docker run --rm --name platform-aux-quick-start-guide -p 9000:9000 platform-service/aux/quick-start-guide
```
