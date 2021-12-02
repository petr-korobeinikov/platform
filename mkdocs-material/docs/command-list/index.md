# Список команд

## start

Запуск локального окружения разработки.

```shell
platformctl start
```

## stop

Остановка локального окружения разработки.

```shell
platformctl stop
```

## lib sync

Синхронизация зависимых библиотек.

Скрывает под собой необходимые переменные окружения для похода за зависимостями.

```shell
platformctl lib sync
```

## lib get

Установка новой зависимой библиотеки.
Если версия не указана, устанавливается самая свежая.

Скрывает под собой необходимые переменные окружения для похода за зависимостями.

```shell
platformctl lib get library_name [version]
```

## service create

Создание нового сервиса.

```shell
platformctl service create group_name/service_name
```

## service start

Запуск сервиса в локальном окружении.

```shell
platformctl service start
```

## service debug

Запуск зависимостей сервиса в локальном окружении. Сам сервис предполагается
запускать на хост-машине с помощью отладчика.

```shell
platformctl service debug
```

## service stop

Остановка сервиса в локальном окружении.

Команда останавливает и удаляет все зависимости текущего сервиса из локального окружения.

```shell
platformctl service stop
```
