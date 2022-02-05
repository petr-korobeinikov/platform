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

Установка новой зависимой библиотеки. Если версия не указана, устанавливается
самая свежая.

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

## service log

Просмотр логов сервиса, запущенного в локальном окружении.

```shell
platformctl service log
```

## service component

Список компонентов сервиса и платформы с адресами доступа.

```shell
platformctl service component
```

## service env

Вывод переменных окружения сервиса для просмотра и экспорта.

```shell
platformctl service env
platformctl service env --service-env local
platformctl service env --service-env staging
platformctl service env --service-env production

eval $(platformctl service env)
eval $(platformctl service env --service-env local)
eval $(platformctl service env --service-env staging)
eval $(platformctl service env --service-env production)
```
