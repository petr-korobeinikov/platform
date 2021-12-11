# platform

![platform](platform-hack/asset/platform-icon-256.png)

## Документация

Запуск контейнера с документацией для локального просмотра:

```shell
docker run \
    --pull always \
    --rm \
    -it \
    -d \
    -p 8000:8000 \
    -v ${PWD}/platform-doc:/docs \
    squidfunk/mkdocs-material
```

Для оформления документации используется
соглашение [Semantic Line Breaks](https://sembr.org/).
