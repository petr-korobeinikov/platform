# Работа с плагинами

Плагины используются для более сложной автоматизации платформенных задач.

Их следует разрабатывать и применять только в том случае, если задачу невозможно
автоматизировать и решить через `platformctl task`.

## plugin

Вывод списка установленных плагинов.

```shell
platformctl plugin
```

Пример вывода:

```shell
platformctl plugin
Execute plugins

Usage:
  platformctl plugin [command]

Available Commands:
  shell-showcase See shell-showcase --help for details
...
```

## plugin name

Вызов плагина.

!!! note

    Команда учитывает переменные окружения, заданные в манифесте сервиса
    `platform.yaml`.

```shell
platformctl plugin plugin-name
platformctl plugin plugin-name [plugin-args]

platformctl plugin --service-env local plugin-name [-- plugin-args]
platformctl plugin --service-env staging plugin-name [-- plugin-args]
```

Примеры вызова:

```shell
platformctl plugin shell-showcase -- switch --foo foo --bar bar

platformctl plugin --service-env local shell-showcase -- switch --foo foo --bar bar
platformctl plugin --service-env staging shell-showcase -- switch --foo foo --bar bar
```
