# Postgres

## Ведите разработку приложения в выделенной схеме

Создайте отдельную схему для приложения:

```postgresql
create schema app;
```

Это позволит избежать конфликтов с функциями и типами данных из устанавливаемых
расширений.

## Устанавливайте расширения в отдельные схемы

```postgresql
create extension "uuid-ossp" with schema uuid;

create extension intarray with schema intarray;
```

## Заведите отдельную схему `maintenance` для служебных данных

```postgresql
create schema maintenance;
```

В ней удобно хранить

* данные о применённых миграциях;
* таблицы, необходимые для временного хранения данных при их переносе;

## Используйте единственное число

```postgresql
create table person (
);
create table item (
);
create table request (
);
```

Заводя таблицы в `postgres`, Вы, фактически, создаёете тип данных.

```postgresql
create table person (
    name    text,
    surname text
);

select ($$(John,Doe)$$::person).name;
```

Вы же не именуете типы данных во множественном числе?

```go
type Requests struct { ... }

type People struct {
Name, Surname string
}
```

## Используйте `text` вместо `varchar`

1. `text` короче, чем `varchar`.
2. `text` принято использовать в `postgresql`-сообществе.
3. `text` и `varchar` используют один и тот же механизм хранения и тип
   данный `varlena`.
4. Нет особого смысла навешивать ограничение длины строки на базе в
   виде `varchar(16)`.

    * Во-первых, проверка на стороне приложения проще и быстрее.
    * Во-вторых, обновление типа для больших таблиц при расширении диапазона
      потребует эксклюзивной блокировки, что может вызвать длительный простой
      приложения.

## Пишите sql в нижнем регистре

А стандарт же? Как же стандарт?

Выразительные `sql`-запросы отлично пишутся и отлично читаются даже при
форматировании в нижнем регистре.

В `postgresql`-сообществе "капсить" считается моветоном.

## Полезные ссылки

* Отличный материал, полностью совпадающий с мнением
  автора: [https://levelup.gitconnected.com/how-to-design-a-clean-database-2c7158114e2f](https://levelup.gitconnected.com/how-to-design-a-clean-database-2c7158114e2f)
