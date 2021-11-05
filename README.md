# platform

## Документация

Запуск контейнера с документацией для локального просмотра:

```shell
docker run \
    --rm \
    -it \
    -d \
    -p 8000:8000 \
    -v ${PWD}/mkdocs-material:/docs \
    squidfunk/mkdocs-material
```

Для оформления документации используется
соглашение [Semantic Line Breaks](https://sembr.org/).
