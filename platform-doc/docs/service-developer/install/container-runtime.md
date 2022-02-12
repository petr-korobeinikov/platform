# Установка среды исполнения контейнеров

## Docker Desktop

Преимущества:

- Самый простой и быстрый способ.
- Использует для развёртывания `docker compose` — подходит для простых сервисов.

Недостатки:

- Не позволяет разворачивать одновременно несколько сервисов.
- Для доступности команды `docker compose` на `Linux` требуется дополнительная
  настройка [^docker-compose-setup].

### Установка Docker Desktop на macOS

Для установки `Docker Desktop` на `macOS` выполните команду:

```shell
brew install --cask docker
```

## minikube

В данный момент не доступен. Решение в проработке.

[^docker-compose-setup]: [Настройка команды `docker compose`](https://docs.docker.com/compose/cli-command/#install-on-linux)
