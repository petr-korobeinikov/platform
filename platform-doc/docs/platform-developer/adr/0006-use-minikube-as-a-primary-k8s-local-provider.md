# 6. Use minikube as a primary k8s local provider

Date: 2022-02-18

## Status

Accepted

Supercedes [2. Use Docker Desktop instead of desktop VMs](0002-use-docker-desktop-instead-of-desktop-vms.md)

## Context

Выбор `Docker Desktop` на старте оказался удачным решением.

Это позволило снизить порог входа и провести адаптацию инструмента среди
разработчиков.

Однако ограничения `Docker Desktop`, налагаемые на `k8s` не позволяют полноценно
использовать его в качестве "легкой" альтернативы `minikube`.

В частности, проблемы с поддержкой `ingress` не дают возможности одновременно
доступаться до нескольких одновременно развёрнутых сервисов.

## Decision

Перейти на `minikube`, оставив поддержку `Docker Desktop` и
поддержав `Rancher Desktop` в качестве альтернативы.

## Consequences

Переход на `minikube` открывает прямой путь к поддержке нескольких одновременно
развёрнутых сервисов, единообразию подходов к развёртыванию и эксплуатации
сервисов.
