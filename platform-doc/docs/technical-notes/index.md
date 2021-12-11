# Технические особенности

## ADR

Технические особенности и замечания описаны в `ADR`.

См. каталог `adr` в корне репозитория.

## Локальная разработка

Для упрощения запуска и отладки конфигурации запуска экспортированы в
каталог `platform-hack/intellij/run_configuration`.

```shell
tree platform-hack/intellij/run_configuration
platform-hack/intellij/run_configuration
└── platformctl
    ├── platformctl.run.xml
    ├── platformctl_lib_get.run.xml
    ├── platformctl_lib_get_foobar.run.xml
    ├── platformctl_lib_sync.run.xml
    ├── platformctl_service_create_group_name.run.xml
    ├── platformctl_service_debug.run.xml
    ├── platformctl_service_start.run.xml
    ├── platformctl_start.run.xml
    └── platformctl_stop.run.xml
```

Экспортированную конфигурацию можнооткрыть с помощью диалогового
окна `Run/Debug`:

![Run Configuration](/assets/exported_run_configuration.png)

## Работа с Taskfile

В качестве таск-раннера выбран `Taskfile`[^taskfile].

Пример запуска команды из `Taskfile.yml`:

```shell
task platformctl:build
```

## Запуск minikube через hyperkit

Вариант запуска через `hyperkit` доступен только для `macOS`.

```shell
minikube start \
    --profile=local \
    --driver=hyperkit \
    --memory=16g \
    --cpus=8 \
    --disk-size=50g \
    --nfs-share="${HOME}" \
    --nfs-shares-root="/"
```

[^taskfile]: [https://taskfile.dev/](https://taskfile.dev/)
