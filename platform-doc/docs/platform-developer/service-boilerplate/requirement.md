# Требования к шаблону

1. Сервис имеет единую точку входа. Точка входа остаётся неизменяемой частью
   сервиса.
2. Сервис имеет один главный пакет для размещения кода.
3. Сервис предоставляет `http`-интерфейс на одном порту — по умолчанию для
   платформы `:9000`.
4. Сервис предоставляет эндпоинт `/metrics`, публикующий метрики в формате,
   совместимом с `Prometheus`.
5. Сервис настроен на отправку трейсов.
6. В сервисе сконфигурировано структурированное логгирование.
7. Логгер сервиса поддерживает отправку ошибок в `Sentry`.
8. Сервис не имеет собственного `Dockerfile`. Решение о допустимости
   собственного `Dockerfile` принимается Архитектурным Комитетом.
9. Шаблон сервиса содержит в себе настроенный каркас для написания документации.