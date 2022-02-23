# platform-service / aux / quick-start-guide

## Build

```shell
platformctl task template
docker build --tag platform/quick-start-guide .
```

## Run

```shell
docker run --rm --name platform-quick-start-guide -p 9000:9000 platform/quick-start-guide
```
