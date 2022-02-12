create table service.person
(
    id         uuid primary key,
    created_at timestamp not null default now(),
    name       text      not null,
    surname    text      not null
);
