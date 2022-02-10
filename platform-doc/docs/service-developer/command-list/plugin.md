# Работа с плагинами

## plugin

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
