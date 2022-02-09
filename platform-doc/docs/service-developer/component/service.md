# Компоненты сервиса

Компоненты, необходимые для работы сервиса перечисляются в разделе `component` в
манифесте сервиса `platform.yaml`.

<!-- @formatter:off -->
```yaml
# ...

component:
  - type: postgres
    name: master
    enabled: true

  - type: minio
    name: minio
    enabled: true

# ...
```
<!-- @formatter:on -->

Далее в этом разделе перечислены компоненты, поддерживаемые в платформе в
настоящий момент.

!!! note

    Компонент одного и того же типа можно подклчать несколько раз с разными
    именами.

    Напремер, подключенный несколько раз компонент `postgres` может играть роль
    настоящих шардов в локальном окружении или на стейджинге.

## postgres

СУБД Postgresql:

<!-- @formatter:off -->
```yaml
component:
  - type: postgres
    name: master
    enabled: true
```
<!-- @formatter:on -->

## vault

Hashicorp Vault:

<!-- @formatter:off -->
```yaml
component:
  - type: vault
    name: vault
    enabled: true
```
<!-- @formatter:on -->

## minio

MiniO Object Storage:

<!-- @formatter:off -->
```yaml
component:
  - type: minio
    name: minio
    enabled: true
```
<!-- @formatter:on -->
