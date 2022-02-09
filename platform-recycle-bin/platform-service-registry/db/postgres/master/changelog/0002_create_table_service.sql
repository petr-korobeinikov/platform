create table platform_service_registry.service (
    name       text not null,
    namespace  text not null,
    created_at timestamp,

    primary key (name, namespace)
);
