# Образы библиотеки задач

Задачи, запускаемые командой `platformctl task`, выполняются в изолированном
окружении контейнера.

`Dockerfile`-ы контейнеров находятся в каталоге `platform-image/task`.

Задачи для сборки собственных локальных образов определены в `Taskfile`:

- `platform-image:task:directory-lister:build`
- `platform-image:task:golangci-lint:build`
- `platform-image:task:golang-migrate:build`

Собрать все образы можно с помощью задачи `platform-image:task:all:build`:

```shell
task platform-image:task:all:build
```
