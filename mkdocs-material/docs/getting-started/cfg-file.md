# Конфигурационный файл

`platformctl` по умолчанию использует конфигурационный
файл `~/.platformctl.yaml`.

Ниже перечислены конфигурационные параметры.

## platform.minikube

Параметр `platform.minikube` содержит параметры для запуска `minikube`.

```yaml
platform:
  minikube:
    memory: "4g"
    cpus: "4"
    disk-size: "50g"
```

## go_env_vars

Параметр `go_env_vars` содержит список переменных окружения, необходимых для
установки и обновления пакетов `go`.

```yaml
go_env_vars:
  - "GOSUMDB=off"
  - "GONOPROXY=none"
```
