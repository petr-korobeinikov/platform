# Установка `platformctl`

## Сборка из исходного кода

Получите рабочую копию репозитория `platform`.

```shell
git clone https://github.com/pkorobeinikov/platform ~/Workspace/github.com/pkorobeinikov/platform
```

Перейдите в директорию с рабочей копией.

```shell
cd $_
```

Выполните сборку `platformctl`.

```shell
task platformctl:build
```

Переместите собранный исполняемый файл в `$PATH`, например:

```shell
cp platformctl/bin/platformctl ~/Bin
```
